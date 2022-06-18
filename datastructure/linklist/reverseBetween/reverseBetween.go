package reverseBetween

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

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	flag := 0
	now := new(ListNode)
	firstPointer := head
	leftPointer := new(ListNode)
	rightPointer := new(ListNode)
	for true {
		flag++
		switch flag {
		case left - 1:
			leftPointer = head
		case left:
			rightPointer = head
		case right:
			leftPointer.Next = head
		case right + 1:
			rightPointer.Next = head
			return firstPointer
		}
		if flag >= left && flag <= right {
			next := head.Next
			head.Next = now
			now = head
			head = next
		} else {
			head = head.Next
		}
	}
	return head
}
