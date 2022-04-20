package gobase

import "fmt"

//int形listnode
type ListNode struct {
	Val  int
	Next *ListNode
}

//TODO
/* Author zsj
 * Desc 把int数组改为反方向的listnode
 * Return
 */
func AddListNode(Int []int) ListNode {
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

//TODO
/* Author zsj
 * Desc listnode变成string
 * Return
 */
func ListNodeToSrting(node *ListNode) string {
	var Str1 string
	p := node
	for true {
		str := fmt.Sprintf("%d", p.Val)
		Str1 = str + Str1
		if p.Next != nil {
			p = p.Next
		} else {
			return Str1
		}
	}
	return ""
}

//TODO
/* Author zsj
 * Desc 传入多个链表	 倒着加 like: [6,5,3] [3,2,1] 得到新链表 [9,8,4]
 * Return
 */
func ListNodeMultiPlus(list ...ListNode) ListNode {
	//TODO
	/* 1 第一个FOR 代表永恒迭代 代表可能存在无穷大的listnode
	* 2 第二个FOR 代表每个list 直到全部listnode的下一个为空 否则一直循环
	* 3 十进制加法法则：turn1: 53 turn2：66               turn3:100                   turn4：
	                t1  :3   turn2: 6+5=11%10= >1+6+5 turn3:100%10=0 100/10=10 7 >10
	* 4 t1 now%10 = 余 now/10为 z t2 （余+up）%10 为当前 （余+up）/10 +z 为UP
	* Return
	*/

	var Node ListNode
	rNode := &Node
	//rNode := make([]ListNode, 1)
	var uP int
	var nFlag int
	//1
	for true {
		//2

		flag := true
		var now int
		for i := 0; i < len(list); i++ {
			//判断是否所有下一个是否全部都是nil
			//if list[i].Val == 0 {
			//	break
			//}
			now = list[i].Val + now
			if list[i].Next == nil {
				fmt.Println(i)
				fmt.Println("lmao")
				break
			}
			list[i] = *list[i].Next
			if list[i].Next != nil {
				fmt.Println(list[0].Val)
				fmt.Println("lmao")
				flag = false
			}
		}

		t1 := now % 10
		t11 := now / 10
		t2 := t1 + uP

		iNode := new(ListNode)

		if nFlag == 0 {
			Node.Val = t2 % 10
			Node.Next = iNode

			Node = *iNode
		} else {
			iNode.Val = t2 % 10
			rNode.Next = iNode
			rNode = iNode
		}

		uP = (t2 / 10) + t11
		if flag == true {
			goto final
		}
	}
final:
	return Node
}

//TODO
/* Author zsj
 * Desc 传入多个链表	 倒着加 like: [6,5,3] [3,2,1] 得到新链表 [9,8,4]
 * Return
 */
func ListNodeMultiPlusSecondVersion(list ...ListNode) ListNode {
	//TODO
	/* 1 第一个FOR 代表永恒迭代 代表可能存在无穷大的listnode
		 * 2 第二个FOR 代表每个list 直到全部listnode的下一个为空 否则一直循环
	 	 * 3 十进制加法法则：turn1: 53 turn2：66               turn3:100                   turn4：
		                t1  :3   turn2: 6+5=11%10= >1+6+5 turn3:100%10=0 100/10=10 7 >10
		 * 4 t1 now%10 = 余 now/10为 z t2 （余+up）%10 为当前 （余+up）/10 +z 为UP
		 * Return
	*/

	lIst := make([]int, 0)
	var uP int
	//var nFlag int
	//1
	for true {

		//2
		flag := true
		var now int
		for i := 0; i < len(list); i++ {
			now = list[i].Val + now

			if list[i].Next == nil {
				list[i].Val = 0
				continue
			} else {
				list[i] = *list[i].Next
				flag = false
			}
		}

		t1 := now % 10
		t11 := now / 10
		t2 := t1 + uP

		lISt := t2 % 10
		lIst = append(lIst, lISt)
		uP = (t2 / 10) + t11

		if flag == true {
			if uP != 0 {
				lISt := uP
				lIst = append(lIst, lISt)
			}

			goto final
		}

	}
final:
	return AddListNode(reverseInts(lIst))
}
