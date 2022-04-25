package letcode

//TODO
/* Author zsj
 * Desc给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的
中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。



示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2

示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
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

//TODO
/* Author zsj
 * Desc v1暴力破解版本
 * Return
 */
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//TODO
	/* 1 中位数收集
	 * 2


		 * 3
		 * Return
	*/
	var mid float64
	num := make([]int, len(nums1)+len(nums2))
	switch true {
	//如果nums2第一个比nums1最后一个大
	case len(nums1) == 0:
		num = nums2
	case len(nums2) == 0:
		num = nums1
	case nums2[0] > nums1[len(nums1)-1]:
		num = append(nums1, nums2...)
		//如果nums1第一个比nums2最后一个大
	case nums1[0] > nums2[len(nums2)-1]:
		num = append(nums2, nums1...)
		//正规排序
	default:
		num = OrderlySort(nums1, nums2)
	}

	//如果是偶数 中间的两个数相加/2  如果是奇数 中间的就是中间数

	//偶数
	if len(num)%2 == 0 {
		var i1, i2 int
		i1 = len(num) / 2
		i2 = i1 - 1
		mid = float64(num[i1]+num[i2]) / float64(2)
	} else {
		mid = float64(num[len(num)/2])
	}

	return mid
}

//TODO
/* Author zsj
 * Desc v2 暴力破解指针版
 * Return
 */
func FindMedianSortedArraysSecondVersion(nums1 []int, nums2 []int) float64 {
	var mid float64

	return mid
}

//TODO
/* Author zsj
 * Desc 中位数删除法
 * Return
 */
func FindMedianSortedArraysThreadVersion() {

}

//func init() {
//	//num1 := []int{0, 1, 2, 3, 4}
//	//num2 := []int{5, 6, 9, 10}
//	//num3 := []int{7, 8}
//	num4 := []int{}
//	num5 := []int{1}
//	//fmt.Println(letcode.FindMedianSortedArrays(num1, num2))
//	//fmt.Println(letcode.FindMedianSortedArrays(num2, num3))
//	fmt.Println(FindMedianSortedArrays(num4, num5))
//}
