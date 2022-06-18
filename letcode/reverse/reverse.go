package reverse

import "love/letcode/addtwonumbers"

//func ModAndMulti(x int) int {
//	var result int
//	var multi = 10
//	var mod int
//	for true {
//		if x == 0 {
//			return result
//		}
//		mod = x % 10
//		x /= 10
//		multi = multi * 10
//		result = result + multi*mod
//	}
//	return result
//}

//TODO
/* Author zsj
 * Desc 通过链表的形式 实现反转数字
 * Return
 */
func ListNodeMultiMod(x int) int {
	//TODO
	/* 1 把链表一个一个存进去
	 * 2 链表中每个值*相对应的个数 再for相加
	 * Return
	 */
	//round1

	Node := make([]addtwonumbers.ListNode, 1)

	var mod int
	var flag = 0
	for true {
		if x == 0 {
			goto round2
		}
		mod = x % 10
		x /= 10

		Node[flag].Val = mod
		Node[flag].Next = &Node[flag+1]
		flag++
	}

	//round2
round2:
	return 0
}

func LetReverse(x int) int {

	is := ListNodeMultiMod(x)

	return is
	//switch true {
	//case x == 0:
	//	return 0
	//case x < 0:
	//	//return ModAndMulti(x) * -1
	//	return 0
	//default:
	//	var is := ListNodeMultiMod(x)
	//
	//	//return ModAndMulti(x)
	//	return is
	//}
}

//func init() {
//	LetReverse(1234056789)
//}
