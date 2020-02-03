package fsa_test

import (
	. "github.com/adamvinueza/fsa"
	"testing"
)

func TestNFA(t *testing.T) {
	zero := "0"
	one := "1"
	alphabet := []string{zero, one}
	q1 := NewState(1)
	q2 := NewState(2)
	q3 := NewState(3)
	q4 := NewState(4)
	states := []State{q1, q2, q3, q4}
	finals := []State{q4}
	start := q1
	transitions := []NTransition{
		NTransition{
			Start: q1,
			Token: zero,
			End:   NewStateSet([]State{q1}),
		},
		NTransition{
			Start: q1,
			Token: one,
			End:   NewStateSet([]State{q1, q2}),
		},
		NTransition{
			Start: q2,
			Token: zero,
			End:   NewStateSet([]State{q3}),
		},
		NTransition{
			Start: q2,
			Token: Epsilon,
			End:   NewStateSet([]State{q3}),
		},
		NTransition{
			Start: q3,
			Token: one,
			End:   NewStateSet([]State{q4}),
		},
		NTransition{
			Start: q4,
			Token: zero,
			End:   NewStateSet([]State{q4}),
		},
		NTransition{
			Start: q4,
			Token: one,
			End:   NewStateSet([]State{q4}),
		},
	}
	n, err := NewNFA(
		states,
		alphabet,
		start,
		finals,
		transitions)
	if err != nil {
		t.Fatalf("error creating NFA: %s", err.Error())
	}
	// create a simple test that is bound to fail
	tests := []struct {
		input   string
		accepts bool
	}{
		{"101", true},
		{"000101", true},
		{"00101", true},
		{"0011", true},
		{"0000", false},
		{"100100100", false},
		{"1001010100100", true},
		{"111111111111", true},
	}
	if n == nil {
		t.Fatalf("NFSA is nil")
	}
	for _, test := range tests {
		accepts := n.Accepts(test.input)
		if test.accepts != accepts {
			t.Fatalf("Expected NFA.Accepts to return %t, found %t", test.accepts, accepts)
		}
		n.Reset()
	}
}
