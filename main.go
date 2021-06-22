package main

import (
	"fmt"

	"github.com/TuxedoFish/golang-learning/src/simple"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := simple.SimpleMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 3, 5, 7},
	}

	fmt.Println(sm)
}
