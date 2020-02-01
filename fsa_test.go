package fsa_test

import (
	"testing"

	"github.com/adamvinueza/fsa"
	"github.com/adamvinueza/fsa/delta"
)

func TestAcceptsNoSymbols(t *testing.T) {
	q0 := delta.NewState(0)
	ss := []string{}
	f := fsa.NewAutomaton(
		q0,              // initial state
		ss,              // alphabet
		[]delta.Delta{}, // allowable deltas
		[]delta.State{}, // final states
	)
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

func TestAcceptsEmptyStringOrEvenBinary(t *testing.T) {
	q1 := delta.NewState(1)
	q2 := delta.NewState(2)
	zero := "0"
	one := "1"
	ss := []string{
		zero,
		one,
	}
	f := fsa.NewAutomaton(
		q1,
		ss,
		[]delta.Delta{
			delta.Delta{
				q1,
				zero,
				q1,
			},
			delta.Delta{
				q1,
				one,
				q2,
			},
			delta.Delta{
				q2,
				one,
				q2,
			},
			delta.Delta{
				q2,
				zero,
				q1,
			},
		},
		[]delta.State{
			q1,
		},
	)
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
	q0 := delta.NewState(0)
	q1 := delta.NewState(1)
	q2 := delta.NewState(2)
	a := "a"
	ss := []string{a}
	f := fsa.NewAutomaton(
		q0,
		ss,
		[]delta.Delta{
			delta.Delta{
				q0,
				a,
				q1,
			},
			delta.Delta{
				q1,
				a,
				q2,
			},
			delta.Delta{
				q2,
				a,
				q1,
			},
		},
		[]delta.State{
			q0,
			q2,
		},
	)
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
