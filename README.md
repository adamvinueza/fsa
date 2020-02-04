# FSA Implementations in Go

This repository contains implementations of deterministic and non-deterministic
finite-state automata in Go. It's just an exercise.

The formal definition of a deterministic finite-state automaton (DFA) is as
follows:
```
A DFA is a tuple (Q, A, t, q0, F), where
  1. Q is a finite set of states,
  2. A is a finite set of symbols (the alphabet),
  3. t is the transition function (a mapping of state-symbol pairs to states),
  4. q0 is a member of Q denoting the start state, and
  5. F is a proper subset of Q denoting the set of accepting states.
```
The formal definition of a non-deterministic finite-state automaton (NFA) is as
follows:
```
An NFA is a tuple (Q, A, t, q0, F), where
  1. Q is a finite set of states,
  2. A is a finite set of symbols (the alphabet),
  3. t is the transition function (a mapping of state-symbol pairs to sets of
     states),
  4. q0 is a member of Q denoting the start state, and
  5. F is a proper subset of Q denoting the set of accepting states.
```
This FSA was implemented via TDD. See [dfa_test.go](/dfa_test.go) and
[nfa_test.go](nfa_test.go) for the tests.

A cute wrinkle is that there's a quick-and-dirty way to implement a set of typed
items in Go, using a map from the item to a boolean value: adding an item to the
set just involves setting the map whose key is that item to `true`. (It's very
possible that certain types of things may not be the most suitable as keys in
maps, but we'll ignore that nicety, as we're not designing a library here.) When
we do that it's trivial to define the obvious operations on sets, which makes
the DFA and NFA implementations a lot simpler. It also allows us to easily
define the notion of an automaton accepting a language (as opposed to a string).

([Fatih Arslan](https://arslan.io) has a [far more complete
version](https://github.com/fatih/set) of this way of implementing sets.
Unfortunately, I didn't know about his version at the time, but already decided
that I liked the idea of having strongly typed sets, and his version has empty
interfaces as set elements.)

### Possible extensions

1. Provide a function that transforms an NFA into a DFA.
2. Create functionality to implement the operations on FSAs under which they are
   closed: union, concatenation, and Kleene star.
