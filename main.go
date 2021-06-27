package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/TuxedoFish/golang-learning/src/simple"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	// Create a pointer to a sample message
	sm := doSimple()

	// Read and write demo
	readAndWriteDemo(sm)

	// Read and write JSON demo
	JSONDemo(sm)
}

func JSONDemo(sm proto.Message) {
	// Convert to JSON
	smAsString := toJSON(sm)
	fmt.Println(protojson.Format(sm))

	// Convert from JSON
	sm2 := &simple.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println(sm2)
}

// Marshals from JSON using protojson
func toJSON(pb proto.Message) []byte {
	out, err := protojson.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't write JSON to bytes")
		return make([]byte, 0)
	}

	fmt.Println("Converted to json bytes!")
	return out
}

// Converts from json using protojson
func fromJSON(in []byte, pb proto.Message) {
	err := protojson.Unmarshal(in, pb)

	if err != nil {
		log.Fatalln("Can't write JSON to bytes")
	}

	fmt.Println("Read from json bytes!")
}

// Writes a protobuffer to file and reads it
func readAndWriteDemo(sm proto.Message) {
	// Write to a binary file
	writeToFile("simple.bin", sm)

	// Read into a data structure
	sm2 := &simple.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println("Read from file: ", sm2)
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

// Reading from a binary file
func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Can't serialize to bytes")
		return err
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Failed to read bytes")
		return err
	}

	fmt.Println("Data has been read!")
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
