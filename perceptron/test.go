package main

// This is a test to make sure the percetronis working, so that I can train the perceptron.

import (
	"fmt"
)

func main() {
	inputs := [2]Point{{x: 2, y: 7}, {x: 4, y: -2}}

	fmt.Printf("%d\n", Train(inputs[0].ToSlice()))
	fmt.Printf("%d\n", Train(inputs[1].ToSlice()))

}
