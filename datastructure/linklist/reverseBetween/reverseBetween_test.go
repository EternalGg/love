package reverseBetween

import (
	"fmt"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	a := []int{10, 9, 2, 3, 4, 5, 6, 7, 8, 1}
	fmt.Println(a)
	var z = new(ListNode)
	z.ListToLinked(a)
	zz := ReverseBetween(z, 3, 9)
	fmt.Println(zz)
}
