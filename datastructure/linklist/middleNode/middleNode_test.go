package middleNode

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	a := rand.Perm(5)
	fmt.Println(a)
	var z = new(ListNode)
	z.ListToLinked(a)
	zz := MiddleNode(z)
	fmt.Println(zz)
}
