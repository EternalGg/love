package benchmark

import (
	"math/rand"
	"testing"
)

func BenchmarkRangess(b *testing.B) {
	a := rand.Perm(10000000)
	Rangess(a)
}

func BenchmarkForrr(b *testing.B) {
	a := rand.Perm(10000000)
	Forrr(a)
}
