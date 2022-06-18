package middleNode

type ListNode struct {
	Val  int
	Next *ListNode
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

// MiddleNode return the middle LinkedNode
func MiddleNode(head *ListNode) *ListNode {
	mid := FindMiddle(head, head)
	return mid
}

func FindMiddle(slow *ListNode, fast *ListNode) *ListNode {
	if fast.Next == nil {
		return slow
	}
	slow = slow.Next
	fast = fast.Next.Next
	mid := FindMiddle(slow, fast)
	return mid
}
