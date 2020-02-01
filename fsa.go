package fsa

import (
	"fmt"
	"github.com/adamvinueza/fsa/delta"
)

type Automaton struct {
	Alphabet   []string
	StartState delta.State
	Deltas     map[delta.Trigger]delta.State
	Finals     []delta.State
	current    delta.State
}

func NewAutomaton(start delta.State, abc []string, deltas []delta.Delta, accepts []delta.State) *Automaton {
	a := Automaton{
		Alphabet:   abc,
		StartState: start,
		Finals:     accepts,
		Deltas:     make(map[delta.Trigger]delta.State),
	}
	for _, d := range deltas {
		t := delta.Trigger{
			State: d.StartState,
			Token: d.Token,
		}
		a.Deltas[t] = d.EndState
	}
	a.current = start
	return &a
}

func (a *Automaton) Reset() {
	a.current = a.StartState
}

func (a *Automaton) Accepts(s string) bool {
	head, tail := split(s)
	if len(head) == 0 {
		return a.inFinalState()
	}
	// check for a valid transition from that symbol
	if err := a.transition(head); err == nil {
		return a.Accepts(tail)
	}
	return false
}

func split(s string) (head string, tail string) {
	if len(s) < 2 {
		return s, ""
	}
	return string(s[0]), s[1:]
}

func (a *Automaton) inFinalState() bool {
	for _, e := range a.Finals {
		if e == a.current {
			return true
		}
	}
	return false
}

func (a *Automaton) transition(s string) error {
	err := fmt.Errorf("could not transition from symbol %s", s)
	if a.recognizes(s) {
		if next, ok := a.Deltas[delta.Trigger{
			State: a.current,
			Token: s,
		}]; ok {
			a.current = next
			return nil
		}
	}
	return err
}

func (a *Automaton) recognizes(t string) bool {
	for _, s := range a.Alphabet {
		if s == t {
			return true
		}
	}
	return false
}
