package distance

import (
	"math"
)

// Calculate the cosine distance of two vectors
func CosineDistance(v1, v2 []float64) float64 {
	if len(v1) != len(v2) {
		panic("vectors must be the same length")
	}
	var sum float64
	for i := range v1 {
		sum += v1[i] * v2[i]
	}
	return sum / (Norm(v1) * Norm(v2))
}

// Normalize the vector
func Norm(v []float64) float64 {
	var sum float64
	for _, x := range v {
		sum += x * x
	}
	return math.Sqrt(sum)
}

func CosineTest() float64 {
	// Create two vectors
	v1 := []float64{1, 2, 3}
	v2 := []float64{4, 5, 6}

	// Calculate the cosine distance
	cosineDistance := CosineDistance(v1, v2)
	return cosineDistance
}
