package fsa

import "fmt"

// Epsilon represents the empty string, used in epsilon transitions.
const Epsilon = "EPSILON"

// Delta represents a mapping of a Token and a State to a State.
type Delta struct {
	Key  StateTokenPair
	Next State
}

// NDelta represents a mapping of a Tokn and a State to a set of States.
type NDelta struct {
	Key  StateTokenPair
	Next *StateSet
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

// NTransition represents a change from a beginning State to a set of ending
// States via a Token.
type NTransition struct {
	Start State
	Token string
	End   *StateSet
}

// creates a non-deterministic delta from a transition.
func (n NTransition) delta() NDelta {
	return NDelta{
		Key: StateTokenPair{
			State: n.Start,
			Token: n.Token,
		},
		Next: n.End,
	}
}

func (n *NDelta) String() string {
	return fmt.Sprintf("Key: %s, Next: %s", n.Key, n.Next)
}

// NDeltas represents a non-deterministic delta function.
type NDeltas map[StateTokenPair]*StateSet

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

// StateTokenPair represents the argument to a delta function.
type StateTokenPair struct {
	State State
	Token string
}

func (s StateTokenPair) String() string {
	return fmt.Sprintf("(%s, %s) => ", s.State, s.Token)
}
