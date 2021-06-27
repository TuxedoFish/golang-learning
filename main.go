package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/TuxedoFish/golang-learning/src/simple"
	"google.golang.org/protobuf/proto"
)

func main() {
	sm := doSimple()

	writeToFile("simple.bin", sm)
	// readFromFile()
}

// Writing a protobuffer to file
func writeToFile(fname string, pb proto.Message) error {
	// Create output for protobuffer
	out, err := proto.Marshal(pb)
	// Check if any error
	if err != nil {
		log.Fatalln("Can't serialize to bytes")
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
	}

	fmt.Println("Data has been written!")
	return nil
}

// Returns a pointer to a simple message
func doSimple() *simple.SimpleMessage {
	sm := simple.SimpleMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 3, 5, 7},
	}
	fmt.Println(sm)

	sm.Name = "I have been renamed"
	fmt.Println(sm)

	fmt.Println(sm.GetId())

	return &sm
}
