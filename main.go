package main

import (
	"fmt"
	r "test/random"
	//d "test/distance"
)

func main() {
	r.Test()
	var a int = 1
	fmt.Println("a: ", a)
	var b *int = &a
	fmt.Println("b: ", *b)
	a = 45
	fmt.Println("b: ", *b)
	*b = 31
	fmt.Println("a: ", a)
}
