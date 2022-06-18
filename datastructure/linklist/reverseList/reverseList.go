package reverseList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList interface {
	ListToLinked()
}

func (n *ListNode) ListToLinked(a []int) {
	Linked := make([]ListNode, len(a))

	// handle data
	for key, value := range a {
		Linked[key].Val = value
		if key == len(a)-1 {
			goto label
		}
		Linked[key].Next = &Linked[key+1]
	}
label:
	n.Val = Linked[0].Val
	n.Next = Linked[0].Next

}

// ReverseList iterate edition
func ReverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		second := head.Next
		second.Next = head
		head.Next = nil
		return second
	}
	head = ListIterate(head, head.Next, 0)
	fmt.Println(head)
	return head
}

func ListIterate(up *ListNode, now *ListNode, flag int) *ListNode {

	if now == nil && flag != 1 {

		return up
	}
	next := now.Next
	now.Next = up
	if flag == 0 {
		up.Next = nil
	}
	flag++

	return ListIterate(now, next, flag)
}
