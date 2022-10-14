package main

import "fmt"

func main() {
	fib := createFibGenerator()

	for i := 0; i < 30; i++ {
		fmt.Println(fib())
	}
}

func createFibGenerator() func() int {
	a, b := 1, 1

	var t int

	return func() int {
		t = a
		a = b
		b += t

		return b
	}
}
