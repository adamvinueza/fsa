package fsa_test

import (
	"testing"

	"github.com/adamvinueza/fsa"
)

func TestAcceptsNothing(t *testing.T) {
	q0 := fsa.NewState(0)
	states := []fsa.State{q0}
	alphabet := []string{}
	f, err := fsa.NewAutomaton(
		states,             // allowable states
		alphabet,           // alphabet
		q0,                 // initial state
		[]fsa.Transition{}, // allowable transitions
		[]fsa.State{},      // final states
	)
	if err != nil {
		t.Fatalf("error creating Automaton: %s", err.Error())
	}
	tests := []struct {
		input    string
		accepted bool
	}{
		{
			"",
			false,
		},
		{
			"a",
			false,
		},
	}
	for _, tt := range tests {
		accepted := f.Accepts(tt.input)
		if tt.accepted != accepted {
			t.Fatalf("Expected accepted value of %t, found %t", tt.accepted, accepted)
		}
	}
}

func TestAcceptsOnlyEmptyString(t *testing.T) {
	q1 := fsa.NewState(1)
	q2 := fsa.NewState(2)
	states := []fsa.State{q1, q2}
	alphabet := []string{}
	f, err := fsa.NewAutomaton(
		states,   // allowable states
		alphabet, // alphabet
		q1,       // initial state
		[]fsa.Transition{
			fsa.Transition{
				q1,
				fsa.AnySymbol,
				q2,
			},
		}, // allowable transitions
		[]fsa.State{q1}, // final states
	)
	if err != nil {
		t.Fatalf("error creating Automaton: %s", err.Error())
	}
	tests := []struct {
		input    string
		accepted bool
	}{
		{
			"",
			true,
		},
		{
			"a",
			false,
		},
	}
	for _, tt := range tests {
		accepted := f.Accepts(tt.input)
		if tt.accepted != accepted {
			t.Fatalf("Expected accepted value of %t, found %t", tt.accepted, accepted)
		}
	}
}

func TestAcceptsEmptyStringOrEvenBinary(t *testing.T) {
	q1 := fsa.NewState(1)
	q2 := fsa.NewState(2)
	states := []fsa.State{q1, q2}
	zero := "0"
	one := "1"
	alphabet := []string{
		zero,
		one,
	}
	f, err := fsa.NewAutomaton(
		states,
		alphabet,
		q1,
		[]fsa.Transition{
			fsa.Transition{
				Start: q1,
				Token: zero,
				End:   q1,
			},
			fsa.Transition{
				Start: q1,
				Token: one,
				End:   q2,
			},
			fsa.Transition{
				Start: q2,
				Token: one,
				End:   q2,
			},
			fsa.Transition{
				Start: q2,
				Token: zero,
				End:   q1,
			},
		},
		[]fsa.State{
			q1,
		},
	)
	if err != nil {
		t.Fatalf("error creating Automaton: %s", err.Error())
	}
	tests := []struct {
		input    string
		accepted bool
	}{
		{
			"",
			true,
		},
		{
			"a",
			false,
		},
		{
			"0",
			true,
		},
		{
			"00000000000000000000000",
			true,
		},
		{
			"0000000000000000000000100000000000000000000000",
			true,
		},
		{
			"01",
			false,
		},
		{
			"1110",
			true,
		},
		{
			"1111111111111111111111111111",
			false,
		},
		{
			"000011101110",
			true,
		},
		{
			"0000x11101110",
			false,
		},
	}
	for i, tt := range tests {
		f.Reset()
		accepted := f.Accepts(tt.input)
		if tt.accepted != accepted {
			t.Fatalf("In test %d, expected accepted value of %t, found %t", i, tt.accepted, accepted)
		}
	}
}

func TestAcceptsEvenNumberOfSymbols(t *testing.T) {
	q0 := fsa.NewState(0)
	q1 := fsa.NewState(1)
	q2 := fsa.NewState(2)
	states := []fsa.State{q0, q1, q2}
	a := "a"
	alphabet := []string{a}
	f, err := fsa.NewAutomaton(
		states,
		alphabet,
		q0,
		[]fsa.Transition{
			fsa.Transition{
				Start: q0,
				Token: a,
				End:   q1,
			},
			fsa.Transition{
				Start: q1,
				Token: a,
				End:   q2,
			},
			fsa.Transition{
				Start: q2,
				Token: a,
				End:   q1,
			},
		},
		[]fsa.State{
			q0,
			q2,
		},
	)
	if err != nil {
		t.Fatalf("error creating Automaton: %s", err.Error())
	}
	tests := []struct {
		input    string
		accepted bool
	}{
		{
			"",
			true,
		},
		{
			"a",
			false,
		},
		{
			"aa",
			true,
		},
		{
			"aaa",
			false,
		},
		{
			"aaaa",
			true,
		},
	}
	for i, tt := range tests {
		f.Reset()
		accepted := f.Accepts(tt.input)
		if tt.accepted != accepted {
			t.Fatalf("In test %d, expected accepted value of %t, found %t", i, tt.accepted, accepted)
		}
	}
}
