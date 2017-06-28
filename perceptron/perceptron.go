package main

// This is a Go implementation of the perceptron as created by Daniel Shiffman of The Coding Train in this video:
// https://youtu.be/ntKn5TPHHAk

import "github.com/google/gofuzz"

var weights []float64

// Starts up the Perceptron, initializing the two input weights, and checking if we are in debug mode
func setup() {
	f := fuzz.New().NilChance(0)
	// Generate the weights
	weights = generateWeights(*f)
}

// Guess returns +1 or -1 based on the sum of its inputs times those inputs' weights.
func Guess(inputs []float64) int {
	var sum float64

	setup()

	for i := range inputs {
		sum += inputs[i] * weights[i]
	}

	if sum >= 0 {
		return 1
	}
	return -1
}

// Simply fills out the weights of the Perceptron, giving them a random value between -1 and 1
func generateWeights(f fuzz.Fuzzer) []float64 {
	output := make([]float64, 2)

	// Generate random weights
	for i := range output {
		f.Fuzz(&output[i])
	}
	return output
}
