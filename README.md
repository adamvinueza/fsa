# FSA Implementation in Go

This is an implementation of finite-state automatons (FSAs) in Go. The formal
definition of an FSA is as follows:
```
An FSA is a tuple (Q, A, t, q0, F), where
  1. Q is a finite set of states,
  2. A is a finite set of symbols (the alphabet),
  3. t is the transition function (a mapping of state-symbol pairs to states),
  4. q0 is a member of Q denoting the start state, and
  5. F is a proper subset of Q denoting the set of accepting states.
```
To implement a FSA, we will need:
  - an array of states
  - an array of symbols
  - a map from pairs of states and symbols to states
  - an initial state
  - an array of accepting states

I want to implement this via TDD. What will the tests look like? Well, I want
the FSA to recognize regular languages, so some of the tests should verify that
particular regular languages are recognized by particular FSAs. This suggests I
should also have a way of producing regular languages.
