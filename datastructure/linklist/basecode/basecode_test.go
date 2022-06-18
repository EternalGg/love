package basecode

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestListToLinkedNode(t *testing.T) {
	var aaa = rand.Perm(10)
	var zzz = new(IntLinkedNode)
	zzz.ListToLinked(aaa)
	fmt.Println(zzz)
}
