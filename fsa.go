package fsa

import (
	"fmt"
)

// Automaton represents a finite-State automaton.
type Automaton struct {
	// States represents this automaton's possible States.
	States []State
	// Alphabet represets this automaton's possible symbols.
	Alphabet []string
	// Start represents this automaton's initial State.
	Start State
	// Deltas is this automaton's mapping function from States and symbols to
	// States.
	Deltas Deltas
	// Finals represent this automaton's accepting States.
	Finals []State
	// Represents this automaton's current State.
	current State
}

// NewAutomaton creates an Automaton from the specified configuration.
func NewAutomaton(
	states []State,
	alphabet []string,
	start State,
	transitions []Transition,
	finals []State) (*Automaton, error) {

	a := Automaton{
		States:   states,
		Alphabet: alphabet,
		Start:    start,
		Deltas:   Deltas(make(map[StateTokenPair]State)),
		Finals:   finals,
	}
	// validate final states
	for _, s := range finals {
		if !StatesContains(a.States, s) {
			return nil, fmt.Errorf(`final state "%s" not found in States`, s)
		}
	}
	for _, t := range transitions {
		if err := a.addDelta(t); err != nil {
			return nil, err
		}
	}
	a.current = a.Start
	return &a, nil
}

// Reset sets this Automaton's current state to its initial state.
func (a *Automaton) Reset() {
	a.current = a.Start
}

// Accepts returns true if this Automaton's Deltas function takes the sequence
// of symbols in the specified string from the initial state to an accepting
// state.
func (a *Automaton) Accepts(s string) bool {
	head, tail := behead(s)
	if len(head) == 0 {
		return a.inFinalState()
	}
	// check for a valid transition from that symbol
	if err := a.transition(head); err == nil {
		return a.Accepts(tail)
	}
	return false
}

// addDelta validates the specified Transition and adds a corresponding Delta if
// it is valid.
func (a *Automaton) addDelta(t Transition) error {
	if !StatesContains(a.States, t.Start) {
		return fmt.Errorf(`start state "%s" not found in States`, t.Start)
	}
	if !StatesContains(a.States, t.End) {
		return fmt.Errorf(`end state "%s" not found in States`, t.End)
	}
	if !stringsContains(a.Alphabet, t.Token) {
		return fmt.Errorf(`token "%s" not found in alphabet`, t.Token)
	}
	d := t.delta()
	_, ok := a.Deltas[d.Key]
	if ok {
		return fmt.Errorf(`duplicate delta: "%s"`, d)
	}
	a.Deltas[d.Key] = d.Next
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

func (a *Automaton) inFinalState() bool {
	return StatesContains(a.Finals, a.current)
}

// This function tries to transition from the current state and the specified
// token string to a next valid state in this Automaton. Returns an error if it
// cannot find a valid state.
func (a *Automaton) transition(s string) error {
	err := fmt.Errorf("could not transition from symbol %s", s)
	stp := StateTokenPair{
		State: a.current,
		Token: s,
	}
	next, ok := a.Deltas[stp]
	if ok {
		a.current = next
		return nil
	}
	return err
}
