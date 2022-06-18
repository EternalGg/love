package reverseList

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestReverseList(t *testing.T) {
	a := rand.Perm(100)
	var z = new(ListNode)
	z.ListToLinked(a)
	zz := ReverseList(z)
	fmt.Println(zz)
}
