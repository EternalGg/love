package basecode

import (
	"errors"
)

var ErrType = errors.New("TypeFail")

type LinkedList interface {
	ListToLinked()
}

// StringLinkedNode Provide a struct of string Linked List
type StringLinkedNode struct {
	Val  string
	Next *StringLinkedNode
}

// IntLinkedNode Provide a struct of int Linked List
type IntLinkedNode struct {
	Val  int
	Next *IntLinkedNode
}

// FloatLinkedNode Provide a struct of float Linked List
type FloatLinkedNode struct {
	Val  float64
	Next *FloatLinkedNode
}

func (n *StringLinkedNode) ListToLinked(a []string) {
	Linked := make([]StringLinkedNode, len(a))

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

func (n *IntLinkedNode) ListToLinked(a []int) {
	Linked := make([]IntLinkedNode, len(a))

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

func (n *FloatLinkedNode) ListToLinked(a []float64) {
	Linked := make([]FloatLinkedNode, len(a))

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
