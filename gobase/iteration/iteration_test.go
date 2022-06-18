package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	iterationFirst(0)
}

func TestFibonacci(t *testing.T) {
	fmt.Println(Fibonacci(10))
}

func TestFactorial(t *testing.T) {
	fmt.Println(Factorial(10, 10))
}
