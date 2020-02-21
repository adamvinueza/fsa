package fsa

import "fmt"

type NFA struct {
	Automaton
	Deltas   NDeltas
	current  *StateSet
}

func NewNFA(
	states []State,
	alphabet []string,
	start State,
	finals []State,
	transitions []NTransition) (*NFA, error) {
	n := NFA{
		Automaton: {
			States:   NewStateSet(states),
			Alphabet: alphabet,
			Start:    start,
			Finals:   NewStateSet(finals),
		},
	}
	n.Reset()
	err := n.addDeltas(transitions)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

func (n *NFA) Reset() {
	n.current = NewEmptyStateSet()
	n.current.Add(n.Start)
}

func (n *NFA) Accepts(s string) bool {
	head, tail := behead(s)
	if len(head) == 0 {
		return n.Finals.Subset(n.current)
	}
	// check for a valid transition
	if err := n.transition(head); err == nil {
		return n.Accepts(tail)
	}
	// check for AnySymbol (NOT THE SAME AS AN EPSILON TRANSITION!)
	if err := n.transition(AnySymbol); err == nil {
		return n.Accepts(tail)
	}
	return false
}

func (n *NFA) addDeltas(tt []NTransition) error {
	n.Deltas = NDeltas(make(map[StateTokenPair]*StateSet))
	for _, t := range tt {
		if !n.States.Contains(t.Start) {
			return fmt.Errorf(`start state "%s" not found in States`, t.Start)
		}
		for s := range t.End.states {
			if !n.States.Contains(s) {
				return fmt.Errorf(`end state "%s" not found in States`, s)
			}
		}
		if !stringsContains(n.Alphabet, t.Token) && t.Token != Epsilon {
			return fmt.Errorf(`token "%s" not found in alphabet`, t.Token)
		}
		d := t.delta()
		_, ok := n.Deltas[d.Key]
		if ok {
			return fmt.Errorf(`duplicate delta: "%v"`, d)
		}
		n.Deltas[d.Key] = d.Next
	}
	return nil
}

func (n *NFA) nextStates(tok string) *StateSet {
	ss := NewEmptyStateSet()
	for c := range n.current.states {
		stp := StateTokenPair{
			State: c,
			Token: tok,
		}
		if next, ok := n.Deltas[stp]; ok {
			ss.Union(next)
		}
	}
	// add epsilon states from states found
	for s := range ss.states {
		eps := StateTokenPair{
			State: s,
			Token: Epsilon,
		}
		if next, ok := n.Deltas[eps]; ok {
			ss.Union(next)
		}
	}
	return ss
}

func (n *NFA) transition(s string) error {
	err := fmt.Errorf("could not transition from symbol %s", s)
	ss := n.nextStates(s)
	if len(ss.states) > 0 {
		n.current = ss
		return nil
	}
	return err
}
