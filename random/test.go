package random

import (
	"fmt"
	d "test/distance"
	fft "test/fft"
)

func Test() {
	fmt.Println("hello world")
	fmt.Println("Cosine: ", d.CosineTest())
	fmt.Println("Euclidean: ", d.EuclideanTest())
	fmt.Println("Manhattan: ", d.ManhattanTest())
	fmt.Println("Discrete Fourier Transform: ", fft.DiscreteFourierTransformTest())
}
