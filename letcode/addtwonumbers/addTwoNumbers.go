package addtwonumbers

import (
	"fmt"
	"love/gobase/listnode"
	"strconv"
	"strings"
)

func AddTwoNumbers() {
	list1 := []int{9, 9, 9, 9, 9, 9, 9}
	l11 := listnode.AddListNode(list1)

	list2 := []int{9, 9, 9, 9}
	l22 := listnode.AddListNode(list2)
	fmt.Println(FirstaddTwoNumbers(&l11, &l22))
}

func reverseInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func splitToDigits(n int) []int {
	var ret []int
	if n == 0 {
		return []int{0}
	} else {

		for n != 0 {
			ret = append(ret, n%10)
			n /= 10
		}

		reverseInt(ret)

		return ret
	}
}

const MaxDigits = 100
const BASE = 10

func MakePositiveInt(s string) (n [MaxDigits]int) {
	// make n zero
	for i := 0; i < MaxDigits; i++ {
		n[i] = 0
	}

	for index, digit := range s {
		i := len(s) - index - 1
		switch digit {
		case '0':
			n[i] = 0
		case '1':
			n[i] = 1
		case '2':
			n[i] = 2
		case '3':
			n[i] = 3
		case '4':
			n[i] = 4
		case '5':
			n[i] = 5
		case '6':
			n[i] = 6
		case '7':
			n[i] = 7
		case '8':
			n[i] = 8
		case '9':
			n[i] = 9
		default:
			panic("invalid digit in number string")
		}
	}

	return
}

func AddPositiveInt(a, b [MaxDigits]int) (c [MaxDigits]int) {
	var carry, sum = 0, 0

	// make c zero
	for i := 0; i < MaxDigits; i++ {
		c[i] = 0
	}

	for i := 0; i < MaxDigits; i++ {
		sum = a[i] + b[i] + carry

		if sum >= BASE {
			carry = 1
			sum -= BASE
		} else {
			carry = 0
		}

		c[i] = sum
	}

	if carry != 0 {
		panic("overflow in addition")
	}

	return
}

func PositiveIntToString(a [MaxDigits]int) (result string) {
	isLeadingZero := true
	for i := MaxDigits - 1; i >= 0; i-- {
		if isLeadingZero && a[i] == 0 {
			continue
		} else {
			isLeadingZero = false
			switch a[i] {
			case 0:
				result += "0"
			case 1:
				result += "1"
			case 2:
				result += "2"
			case 3:
				result += "3"
			case 4:
				result += "4"
			case 5:
				result += "5"
			case 6:
				result += "6"
			case 7:
				result += "7"
			case 8:
				result += "8"
			case 9:
				result += "9"
			default:
				panic("invalid digit in int array")
			}
		}
	}
	return
}

func multi(str1, str2 string) (result string) {

	if len(str1) == 0 && len(str2) == 0 {
		result = "0"
		return
	}
	var index1 = len(str1) - 1
	var index2 = len(str2) - 1
	var left int

	for index1 >= 0 && index2 >= 0 {

		c1 := str1[index1] - '0'
		c2 := str2[index2] - '0'
		sum := int(c1) + int(c2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		//fmt.Printf("c3__%c,result:%s\n",c3,result)
		result = fmt.Sprintf("%c%s", c3, result)
		//fmt.Println("result:",result,"\n")
		index1--
		index2--
	}

	for index1 >= 0 {
		c1 := str1[index1] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
	}

	for index2 >= 0 {

		c1 := str2[index2] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index2--
	}
	if left == 1 {
		fmt.Sprintf("1%s", result)
	}
	return
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func AizuArray(A string) []int {
	strs := strings.Split(A, "")
	ary := make([]int, len(strs))
	for i := range ary {
		ary[i], _ = strconv.Atoi(strs[i])
	}
	return ary
}

func FirstaddTwoNumbers(l1 *listnode.ListNode, l2 *listnode.ListNode) *listnode.ListNode {
	//TODO
	/* 1 LIST转STR
	 * 2 STR转INT相加
	 * 3 INT 转 STR 转 LISTNODE
	 * Return
	 */
	//1
	var Str1 string
	Str1 = listnode.ListNodeToSrting(l1)

	var Str2 string
	Str2 = listnode.ListNodeToSrting(l2)

	//2
	//int1, _ := strconv.Atoi(Str1)
	//int2, _ := strconv.Atoi(Str2)

	//3
	fmt.Println(Str1, Str2)
	//fmt.Println(AddPositiveInt(MakePositiveInt(Str1), MakePositiveInt(Str2)))

	str := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(AddPositiveInt(MakePositiveInt(Str1), MakePositiveInt(Str2)))), ","), "[]")

	//fmt.Println(str)
	num := Reverse(str)
	fmt.Println(num)
	reverse := AizuArray(num)
	fmt.Println(reverse)
	rListNode := listnode.AddListNode(reverse)
	return &rListNode
}

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

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func SecondaddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//TODO
	/* 1 两个链表 每一位都加 如果加起来>10 下一位+1 递归操作 创建一个新链表
	 * 2
	 * 3
	 * Return
	 */
	Node := ListNodeMultiPlusSecondVersion(*l1, *l2)
	return &Node
}

//func init() {
//	list1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
//	l11 := gobase.AddListNode(list1)
//
//	list2 := []int{0}
//	l22 := gobase.AddListNode(list2)
//	Node := gobase.ListNodeMultiPlusSecondVersion(l11, l22)
//	fmt.Println(gobase.ListNodeToSrting(&Node))
//}
