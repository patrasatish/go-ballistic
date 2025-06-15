package main

import "fmt"

// enum -> type fixed number of values
// distinct name
// not a distinct language feature
// simple to implement using existing language idiom

// enum type ServerState has underlying int type
type ServerState int

// possible values of enum are defined as constants
const (

	// iota generates successive constant values, here 0, 1, 2 and so on
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// implemented the fmt.Stringer interface
// values of ServerState can be printed out or converted to strings.
// alt - stringer tool in conjuction with go:generate
func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// checking some predicates
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("Unknown State: %s", s))
	}
}
