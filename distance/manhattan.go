package distance

import "math"

// Calculate the manhattan distance of two vectors
func ManhattanDistance(v1, v2 []float64) float64 {
	if len(v1) != len(v2) {
		panic("vectors must be the same length")
	}
	var sum float64
	for i := range v1 {
		sum += math.Abs(v1[i] - v2[i])
	}
	return sum
}

func ManhattanTest() float64 {
	// Create two vectors
	v1 := []float64{1, 2, 3}
	v2 := []float64{4, 5, 6}

	// Calculate the manhattan distance
	manhattanDistance := ManhattanDistance(v1, v2)
	return manhattanDistance
}
