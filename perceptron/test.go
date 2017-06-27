package main

// This is a test to make sure the percetronis working, so that I can train the perceptron.

import (
	"fmt"
)

type point struct {
	x, y float64
}

func (input point) toSlice() []float64 {
	output := make([]float64, 2)

	output[0] = input.x
	output[1] = input.y

	return output
}

func main() {
	inputs := [2]point{{x: 2, y: 7}, {x: 4, y: -2}}

	fmt.Printf("%d\n", Guess(inputs[0].toSlice()))
	fmt.Printf("%d\n", Guess(inputs[1].toSlice()))
}
