package fsa

import (
	"fmt"
)

// Represents a state in a finite-state automaton.
type State string

// NewState just builds a state string from an index.
func NewState(idx int) State {
	return State(fmt.Sprintf("q%d", idx))
}

// StatesContains returns true if the specified slice of States contains the
// specified State.
func StatesContains(ss []State, s State) bool {
	for _, e := range ss {
		if e == s {
			return true
		}
	}
	return false
}

// Delta represents a mapping of a Token and a State to a State.
type Delta struct {
	Key  StateTokenPair
	Next State
}

func (d Delta) String() string {
	return fmt.Sprintf(`Key: "%s", Next: %s`, d.Key, d.Next)
}

// Represents a set of Deltas, i.e. a function.
type Deltas map[StateTokenPair]State

// Transition represents a change from a beginning State to an ending State via
// a Token.
type Transition struct {
	Start State
	Token string
	End   State
}

// creates a Delta from a Transition.
func (t Transition) delta() Delta {
	return Delta{
		Key: StateTokenPair{
			State: t.Start,
			Token: t.Token,
		},
		Next: t.End,
	}
}

// Represents the argument to a Deltas function.
type StateTokenPair struct {
	State State
	Token string
}
