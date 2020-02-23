package fsa_test

import (
	"testing"

	"github.com/leanovate/gopter/gen"

	. "github.com/adamvinueza/fsa"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/prop"
)

func generateLanguage(t *testing.T, f interface{}) *Language {
	langGen := gen.AnyString().SuchThat(f)
	var sentences []string
	var i int
	for i < 100 {
		s, ok := langGen.Sample()
		if !ok {
			continue
		}
		i++
		sentences = append(sentences, s.(string))
	}
	lang, err := NewLanguage(sentences)
	if err != nil {
		t.Fatal(err.Error())
	}
	return lang
}

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
	properties := gopter.NewProperties(nil)
	properties.Property("Nothing is accepted", prop.ForAll(
		func(s string) bool {
			a := f.Accepts(s)
			return a == false
		},
		gen.AnyString(),
	))
	properties.TestingRun(t)
}

func TestAcceptsOnlyEmptyString(t *testing.T) {
	q0 := NewState(1)
	q1 := NewState(2)
	states := []State{q0, q1}
	alphabet := []string{}
	f, err := NewDFA(
		states,   // allowable states
		alphabet, // alphabet
		q0,       // initial state
		[]Transition{
			// Any symbol takes the FSA from its sole final state to its sole
			// non-final state; this is the only transition. Saves us the
			// trouble of using regular expressions, or creating a humongous
			// number of transitions from distinct symbols, or modifying
			// transitions to take slices of symbols.
			Transition{
				q0,
				AnySymbol,
				q1,
			},
		}, // allowable transitions
		[]State{q0}, // final states
	)
	if err != nil {
		t.Fatalf("error creating Automaton: %s", err.Error())
	}
	properties := gopter.NewProperties(nil)
	properties.Property("Only the empty string is accepted", prop.ForAll(
		func(s string) bool {
			a := f.Accepts(s)
			empty := len(s) == 0
			return a == empty
		},
		gen.AnyString(),
	))
	properties.TestingRun(t)
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
	properties := gopter.NewProperties(nil)
	properties.Property("The empty string or any string parsing to an even binary is accepted", prop.ForAll(
		func(s string) bool {
			a := f.Accepts(s)
			if len(s) == 0 {
				return a
			}
			// even binary strings must end in "0"
			even := string(s[len(s)-1]) == "0"
			return a == even
		},
		gen.RegexMatch("^[01]*$"),
	))
	properties.Property("Any string that is not binary is not accepted", prop.ForAll(
		func(s string) bool {
			a := f.Accepts(s)
			return a == false
		},
		gen.RegexMatch("[^01]"),
	))
	properties.TestingRun(t)
}

func TestAcceptsEvenNumberOfSymbols(t *testing.T) {
	q0 := NewState(0)
	q1 := NewState(1)
	q2 := NewState(2)
	states := []State{q0, q1, q2}
	a := AnySymbol
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
	properties := gopter.NewProperties(nil)

	properties.Property("Any string with an even number of symbols is accepted", prop.ForAll(
		func(s string) bool {
			a := f.Accepts(s)
			even := len(s)%2 == 0
			return a == even
		},
		gen.AnyString(),
	))
	properties.TestingRun(t)
}

func TestAcceptsEvenAsLanguage(t *testing.T) {
	q0 := NewState(0)
	q1 := NewState(1)
	q2 := NewState(2)
	states := []State{q0, q1, q2}
	a := AnySymbol
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
		input    *Language
		accepted bool
	}{
		{
			generateLanguage(t, func(s string) bool {
				return len(s)%2 == 0
			}),
			true,
		},
		{
			generateLanguage(t, func(s string) bool {
				return len(s)%2 == 1
			}),
			false,
		},
	}
	for i, tt := range tests {
		accepted := f.AcceptsLanguage(tt.input)
		if tt.accepted != accepted {
			t.Fatalf("In test %d, expected accepted value of %t, found %t", i, tt.accepted, accepted)
		}
	}
}

func TestAcceptsOdd(t *testing.T) {
	q0 := NewState(0)
	q1 := NewState(1)
	q2 := NewState(2)
	q3 := NewState(3)
	states := []State{q0, q1, q2, q3}
	a := AnySymbol
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
				End:   q3,
			},
			Transition{
				Start: q3,
				Token: a,
				End:   q0,
			},
		},
		[]State{
			q1,
			q3,
		},
	)
	if err != nil {
		t.Fatalf("error creating Automaton: %s", err.Error())
	}
	properties := gopter.NewProperties(nil)

	properties.Property("Any string with an odd number of symbols is accepted", prop.ForAll(
		func(s string) bool {
			a := f.Accepts(s)
			odd := len(s)%2 == 1
			return a == odd
		},
		gen.AnyString(),
	))
	properties.TestingRun(t)
}
