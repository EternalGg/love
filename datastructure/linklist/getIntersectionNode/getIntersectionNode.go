// Package getIntersectionNode 相交链表 https://leetcode.cn/problems/intersection-of-two-linked-lists/
package getIntersectionNode

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

// GetIntersectionNode 双指针解法 a,b的长度为 a = A , b = B
// 假设c为相交的点 A = c+ (A-c) , B = c +(B-c)
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	first := headA
	second := headB

	for first.Next != nil && second.Next != nil {
		first = first.Next
		second = second.Next
	}

	if first.Next == nil {
		first = headB
		for first.Next != nil && second.Next != nil {
			first = first.Next
			second = second.Next
		}
		second = headA
		for first.Val != second.Val {
			first = first.Next
			second = second.Next
		}
	} else {
		second = headA
		for first.Next != nil && second.Next != nil {
			first = first.Next
			second = second.Next
		}
		first = headB
		for first.Val != second.Val {
			first = first.Next
			second = second.Next
		}
	}

	return first
}

//func findLast(first, second *ListNode) {
//
//}
