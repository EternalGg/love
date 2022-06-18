// Package hasCycle 是否是环形链表 https://leetcode.cn/problems/linked-list-cycle/
package hasCycle

import "fmt"

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

// HasCycle 快慢指针版本
func HasCycle(head *ListNode) bool {
	slow := head.Next
	fast := head.Next.Next

	for true {
		if fast.Next == nil {
			return false
		}
		if fast == slow {
			fmt.Println(fast)
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// HashHasCycle 哈希表版本
func HashHasCycle(head *ListNode) bool {
	var hash = map[*ListNode]int{}

	for true {
		if head.Next == nil {
			return false
		}
		if hash[head] != 0 {
			return true
		}
		hash[head] = head.Val

		head = head.Next
	}
	return true
}
