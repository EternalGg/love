package getIntersectionNode

import (
	"fmt"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	a := []int{5, 4, 3, 2, 1}
	b := []int{3, 2, 1}
	var z = new(ListNode)
	var x = new(ListNode)
	z.ListToLinked(a)
	x.ListToLinked(b)
	fmt.Println(GetIntersectionNode(z, x))
}
