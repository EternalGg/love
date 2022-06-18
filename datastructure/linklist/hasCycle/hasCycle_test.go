package hasCycle

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {
	a := []int{5, 4, 3, 2, 1}

	var z = new(ListNode)
	z.ListToLinked(a)
	z.Next.Next.Next.Next.Next = z.Next.Next
	fmt.Println(HasCycle(z))
}

func TestHashHasCycle(t *testing.T) {
	a := []int{5, 4, 3, 2, 1}

	var z = new(ListNode)
	z.ListToLinked(a)
	z.Next.Next.Next.Next.Next = z.Next.Next
	fmt.Println(HashHasCycle(z))
}
