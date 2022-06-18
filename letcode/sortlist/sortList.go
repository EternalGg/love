package sortlist

import (
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addListNode(Int []int) ListNode {
	Node := make([]ListNode, len(Int))
	for key, value := range Int {
		Node[key].Val = value
		if key == len(Int)-1 {
			goto label
		}
		Node[key].Next = &Node[key+1]
	}
label:
	return Node[0]
}

func sortList(head *ListNode) *ListNode {
	list := []int{0, 1, 2}
	listnode := addListNode(list)
	return &listnode
}

func init() {
	list := new(ListNode)
	log.Println(sortList(list).Next)
}
