package fsa_test

import (
	. "github.com/adamvinueza/fsa"
	"testing"
)

func TestAcceptsNothing(t *testing.T) {
	q0 := NewState(0)
	states := []State{q0}
	alphabet := []string{}
	f, err := NewDFA(
		states,         // allowable states
		alphabet,       // alphabet
		q0,             // initial state
		[]Transition{}, // allowable transitions
		[]State{},      // final states
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
	q1 := NewState(1)
	q2 := NewState(2)
	states := []State{q1, q2}
	alphabet := []string{}
	f, err := NewDFA(
		states,   // allowable states
		alphabet, // alphabet
		q1,       // initial state
		[]Transition{
			// Any symbol takes the FSA from its sole final state to its sole
			// non-final state; this is the only transition. Saves us the
			// trouble of using regular expressions, or creating a humongous
			// number of transitions from distinct symbols, or modifying
			// transitions to take slices of symbols.
			Transition{
				q1,
				AnySymbol,
				q2,
			},
		}, // allowable transitions
		[]State{q1}, // final states
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
	q1 := NewState(1)
	q2 := NewState(2)
	states := []State{q1, q2}
	zero := "0"
	one := "1"
	alphabet := []string{
		zero,
		one,
	}
	f, err := NewDFA(
		states,
		alphabet,
		q1,
		[]Transition{
			Transition{
				Start: q1,
				Token: zero,
				End:   q1,
			},
			Transition{
				Start: q1,
				Token: one,
				End:   q2,
			},
			Transition{
				Start: q2,
				Token: one,
				End:   q2,
			},
			Transition{
				Start: q2,
				Token: zero,
				End:   q1,
			},
		},
		[]State{
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
	q0 := NewState(0)
	q1 := NewState(1)
	q2 := NewState(2)
	states := []State{q0, q1, q2}
	a := "a"
	alphabet := []string{a}
	f, err := NewDFA(
		states,
		alphabet,
		q0,
		[]Transition{
			Transition{
				Start: q0,
				Token: a,
				End:   q1,
			},
			Transition{
				Start: q1,
				Token: a,
				End:   q2,
			},
			Transition{
				Start: q2,
				Token: a,
				End:   q1,
			},
		},
		[]State{
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

func TestAcceptsEvenAsLanguage(t *testing.T) {
	q0 := NewState(0)
	q1 := NewState(1)
	q2 := NewState(2)
	states := []State{q0, q1, q2}
	a := "a"
	alphabet := []string{a}
	f, err := NewDFA(
		states,
		alphabet,
		q0,
		[]Transition{
			Transition{
				Start: q0,
				Token: a,
				End:   q1,
			},
			Transition{
				Start: q1,
				Token: a,
				End:   q2,
			},
			Transition{
				Start: q2,
				Token: a,
				End:   q1,
			},
		},
		[]State{
			q0,
			q2,
		},
	)
	if err != nil {
		t.Fatalf("error creating Automaton: %s", err.Error())
	}
	acceptableLang, err := NewLanguage([]string{"aa", "aaaa", "aaaaaa", "aaaaaaaa", "aaaaaaaaaaaaaaaa"})
	if err != nil {
		t.Fatalf("error creating language: %s", err.Error())
	}
	unacceptableLang, err := NewLanguage([]string{"aa", "aaaa", "aaaaaa", "aaaaaaaaa", "aaaaaaaaaaaaaaaa"})
	if err != nil {
		t.Fatalf("error creating language: %s", err.Error())
	}
	tests := []struct {
		input    *Language
		accepted bool
	}{
		{
			acceptableLang,
			true,
		},
		{
			unacceptableLang,
			false,
		},
	}
	for i, tt := range tests {
		f.Reset()
		accepted := f.AcceptsLanguage(tt.input)
		if tt.accepted != accepted {
			t.Fatalf("In test %d, expected accepted value of %t, found %t", i, tt.accepted, accepted)
		}
	}
}
