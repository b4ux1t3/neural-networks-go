package main

// Points represent a point in 2D space, and are computed by perceptrons.
// They have an x and y coordinate and a label that tells whether it is
// above or below an arbitrary line. (y = x, in this case).

type Point struct {
	x, y  float64
	label int
}

// ToSlice turns the x and y coordinates into a slice that the perceptron can understand
func (input Point) ToSlice() []float64 {
	output := make([]float64, 3)

	output[0] = input.x
	output[1] = input.y
	calculateLabel(&input)

	return output
}

// This will calculate whether the point is above or below the line.
func calculateLabel(point *Point) {
	if point.x < point.y {
		point.label = -1
	} else {
		point.label = 1
	}
}
