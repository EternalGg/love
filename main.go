package main

import (
	"fmt"
	"love/gobase"
)

func main() {

	//letcode.TwoSum()
	//letcode.AddTwoNumbers()
	list1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	l11 := gobase.AddListNode(list1)

	list2 := []int{0}
	l22 := gobase.AddListNode(list2)
	Node := gobase.ListNodeMultiPlusSecondVersion(l11, l22)
	fmt.Println(gobase.ListNodeToSrting(&Node))
}
