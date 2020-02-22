package fsa

import (
	"fmt"
)

const AnySymbol = "ANY_SYMBOL"

// AutomatonBase represents the core of a finite-state automaton.
type AutomatonBase struct {
	// States represents this automaton's possible States.
	States *StateSet
	// Alphabet represets this automaton's possible symbols.
	Alphabet []string
	// Start represents this automaton's initial State.
	Start State
	// Finals represent this automaton's accepting States.
	Finals *StateSet
}

type DFA struct {
	AutomatonBase
	// Deltas is this automaton's mapping function from States and symbols to
	// States.
	Deltas Deltas
	// Represents this automaton's current State.
	current State
}

// NewDFA creates a deterministic automaton from the specified configuration.
func NewDFA(
	states []State,
	alphabet []string,
	start State,
	transitions []Transition,
	finals []State) (*DFA, error) {

	dfa := DFA{
		AutomatonBase: AutomatonBase{
			States:   NewStateSet(states),
			Alphabet: alphabet,
			Start:    start,
			Finals:   NewStateSet(finals),
		},
		Deltas: Deltas(make(map[StateTokenPair]State)),
	}
	// validate final states
	for _, s := range finals {
		if !dfa.States.Contains(s) {
			return nil, fmt.Errorf(`final state "%s" not found in States`, s)
		}
	}
	for _, t := range transitions {
		if err := dfa.addDelta(t); err != nil {
			return nil, err
		}
	}
	dfa.Reset()
	return &dfa, nil
}

// Reset sets this Automaton's current state to its initial state.
func (dfa *DFA) Reset() {
	dfa.current = dfa.Start
}

// Accepts returns true if this Automaton's Deltas function takes the sequence
// of symbols in the specified string from the initial state to an accepting
// state.
func (dfa *DFA) Accepts(s string) bool {
	head, tail := behead(s)
	if len(head) == 0 {
		return dfa.Finals.Contains(dfa.current)
	}
	// check for a valid transition from that symbol
	if err := dfa.transition(head); err == nil {
		return dfa.Accepts(tail)
	}
	// if that particular symbol does not work, see if AnySymbol works
	if err := dfa.transition(AnySymbol); err == nil {
		return dfa.Accepts(tail)
	}
	return false
}

// AcceptsLanguage returns true if this automaton accepts every symbol in the
// specified Language.
func (dfa *DFA) AcceptsLanguage(l *Language) bool {
	for s := range l.expressions.expressions {
		if !dfa.Accepts(s) {
			return false
		}
		dfa.Reset()
	}
	return true
}

// addDelta validates the specified Transition and adds a corresponding Delta if
// it is valid.
func (dfa *DFA) addDelta(t Transition) error {
	if !dfa.States.Contains(t.Start) {
		return fmt.Errorf(`start state "%s" not found in States`, t.Start)
	}
	if !dfa.States.Contains(t.End) {
		return fmt.Errorf(`end state "%s" not found in States`, t.End)
	}
	if !stringsContains(dfa.Alphabet, t.Token) && t.Token != AnySymbol {
		return fmt.Errorf(`token "%s" not found in alphabet`, t.Token)
	}
	d := t.delta()
	_, ok := dfa.Deltas[d.Key]
	if ok {
		return fmt.Errorf(`duplicate delta: "%s"`, d)
	}
	dfa.Deltas[d.Key] = d.Next
	return nil
}

// behead breaks a string into a head and tail, where the head is the string
// consisting of the first character and the tail is the remaining part of the
// string (empty if the string is at most one character long).
func behead(s string) (head string, tail string) {
	if len(s) < 2 {
		return s, ""
	}
	return string(s[0]), s[1:]
}

func stringsContains(ss []string, s string) bool {
	for _, e := range ss {
		if e == s {
			return true
		}
	}
	return false
}

// This function tries to transition from the current state and the specified
// token string to a next valid state in this Automaton. Returns an error if it
// cannot find a valid state.
func (dfa *DFA) transition(s string) error {
	err := fmt.Errorf("could not transition from symbol %s", s)
	stp := StateTokenPair{
		State: dfa.current,
		Token: s,
	}
	next, ok := dfa.Deltas[stp]
	if ok {
		dfa.current = next
		return nil
	}
	return err
}
