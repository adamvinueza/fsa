package fsa

import (
	"fmt"
	"strings"
)

// Represents a state in a finite-state automaton.
type State string

// NewState just builds a state string from an index.
func NewState(idx int) State {
	return State(fmt.Sprintf("q%d", idx))
}

type StateSet struct {
	states map[State]bool
}

func (ss *StateSet) String() string {
	s := []string{}
	for k := range ss.states {
		s = append(s, fmt.Sprintf("(%s)", k))
	}
	return strings.Join(s, ",")
}

func (ss *StateSet) Add(s State) {
	ss.states[s] = true
}

func (ss *StateSet) AddSlice(states []State) {
	for _, s := range states {
		ss.Add(s)
	}
}

func (ss *StateSet) Remove(s State) {
	delete(ss.states, s)
}

func (ss *StateSet) Contains(s State) bool {
	return ss.states[s]
}

func (ss *StateSet) Subset(tt *StateSet) bool {
	for k, _ := range ss.states {
		if !tt.Contains(k) {
			return false
		}
	}
	return true
}

func (ss *StateSet) Superset(tt *StateSet) bool {
	return tt.Subset(ss)
}

func (ss *StateSet) Union(tt *StateSet) {
	for k, _ := range tt.states {
		ss.states[k] = true
	}
}

func NewEmptyStateSet() *StateSet {
	ss := StateSet{}
	ss.states = make(map[State]bool)
	return &ss
}

func NewStateSet(states []State) *StateSet {
	ss := NewEmptyStateSet()
	for _, s := range states {
		ss.Add(s)
	}
	return ss
}
