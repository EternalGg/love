package iteration

import "fmt"

func iterationFirst(enter int) {
	if enter == 10 {
		return
	}
	enter++
	fmt.Println(enter)
	iterationFirst(enter)
}

func Fibonacci(n int) int {
	a, b := n%2, 1

	for i := 0; i < n/2; i++ {
		a += b
		b += a
	}

	return a
}

func Factorial(enter int, value int) int {
	if enter == 1 {
		return value
	}

	value = value * (enter - 1)
	enter--
	return Factorial(enter, value)
}
