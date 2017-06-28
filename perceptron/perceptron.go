package main

// This is a Go implementation of the perceptron as created by Daniel Shiffman of The Coding Train in this video:
// https://youtu.be/ntKn5TPHHAk

import "github.com/google/gofuzz"

var weights []float64
var learningRate float64

// Starts up the Perceptron, initializing the two input weights
func setup() {
	f := fuzz.New().NilChance(0)
	// Generate the weights
	weights = generateWeights(*f)
	learningRate = 0.1
}

// Guess returns +1 or -1 based on the sum of its inputs times those inputs' weights.
func Guess(inputs []float64) int {
	var sum float64

	setup()

	for i := 0; i < 2; i++ {
		sum += inputs[i] * weights[i]
	}

	if sum >= 0 {
		return 1
	}
	return -1
}

// Train is for training perceptrons using gradient descent.
// This will have 4 steps:
// 1. Proviude perceptron with inputs for which there is a known answer.
// 2. Ask the perceptron to guess teh answer
// 3. Compute the error.
// 4. Adjust all weights according to the error
// TODO: Make this work
func Train(inputs []float64, target int) {
	guess := Guess(inputs)
	err := float64(target - guess)
	for i := 0; i < 2; i++ {
		weights[i] += err * inputs[i] * learningRate
	}
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
