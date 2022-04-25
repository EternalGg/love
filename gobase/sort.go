package gobase

//TODO
/* Author zsj
 * Desc 输入两个正序数组 返回一个由两个数组合并的正序数组
 * Return
 */
func OrderlySort(num1 []int, num2 []int) []int {
	num := make([]int, len(num1)+len(num2))

	var l1, l2 int
	for i := 0; i < len(num); i++ {
		if l1 == len(num1) {
			for j := i; j < len(num); j++ {
				num[j] = num2[l2]
				l2++
			}
			return num
		}
		if l2 == len(num2) {
			for j := i; j < len(num); j++ {
				num[j] = num1[l1]
				l1++
			}
			return num
		}
		if num1[l1] < num2[l2] {
			num[i] = num1[l1]
			l1++
		} else {
			num[i] = num2[l2]
			l2++
		}
	}
	return num
}
