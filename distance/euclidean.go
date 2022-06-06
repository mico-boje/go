package distance

import (
	"math"
)

// Calculate the euclidean distance of two vectors
func EuclideanDistance(v1, v2 []float64) float64 {
	if len(v1) != len(v2) {
		panic("vectors must be the same length")
	}
	var sum float64
	for i := range v1 {
		sum += math.Pow(v1[i]-v2[i], 2)
	}
	return math.Sqrt(sum)
}

func EuclideanTest() float64 {
	// Create two vectors
	v1 := []float64{1, 2, 3}
	v2 := []float64{4, 5, 6}

	// Calculate the euclidean distance
	euclideanDistance := EuclideanDistance(v1, v2)
	return euclideanDistance
}
