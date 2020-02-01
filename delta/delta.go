package delta

import (
	"fmt"
)

type State string

func NewState(idx int) State {
	return State(fmt.Sprintf("q%d", idx))
}

type Delta struct {
	StartState State
	Token      string
	EndState   State
}

type Trigger struct {
	State State
	Token string
}
