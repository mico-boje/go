package fft

import "math"

// Compute the one dimensional discrete Fourier Transform of the input vector.
func DiscreteFourierTransform(v []float64) []float64 {
	// Initialize the output vector
	output := make([]float64, len(v))

	// Compute the discrete Fourier Transform
	for k := 0; k < len(v); k++ {
		for n := 0; n < len(v); n++ {
			output[k] += v[n] * math.Cos(2*math.Pi*float64(k)*float64(n)/float64(len(v)))
		}
	}

	return output
}

func DiscreteFourierTransformTest() []float64 {
	// Create a vector
	v := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Compute the discrete Fourier Transform
	discreteFourierTransform := DiscreteFourierTransform(v)

	return discreteFourierTransform
}
