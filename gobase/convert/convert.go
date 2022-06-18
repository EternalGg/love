package convert

import (
	"strconv"
	"strings"
)

func reverseInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//TODO
/* Author zsj
 * Desc 把int转换为int数组
 * Return
 */
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

//TODO
/* Author zsj
 * Desc 反转int数组
 * Return
 */
func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

//TODO
/* Author zsj
 * Desc 反转英文字符串
 * Return
 */
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//TODO
/* Author zsj
 * Desc string转换为int数组
 * Return
 */
func AizuArray(A string) []int {
	strs := strings.Split(A, "")
	ary := make([]int, len(strs))
	for i := range ary {
		ary[i], _ = strconv.Atoi(strs[i])
	}
	return ary
}
