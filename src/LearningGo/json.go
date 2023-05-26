package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
  ]`)

	type Animal struct {
		Name  string
		Order string
	}

	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("Parsing error: ", err)
	}

	fmt.Printf("%+v", animals)

	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}

	group := ColorGroup{
		ID:   1,
		Name: "Reds",
		Colors: []string{
			"Crimson", "Red", "Ruby", "Maroon",
		},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("Encoding error: ", err)
	}

	os.Stdout.Write(b)

	// Printing it out directly, all you can see is an array of numbers.
	fmt.Println(b)
}
