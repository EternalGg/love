// Package sort provides ten basic sorting algorithm.
package sort

import (
	"fmt"
	"math/rand"
)

//OrderlySort
/* Author zsj
 * Desc 输入两个正序数组 返回一个由两个数组合并的正序数组
 * Return
 */
func orderlySort(num1 []int, num2 []int) []int {
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

func exponent(a, n uint64) uint64 {
	result := uint64(1)
	for i := n; i > 0; i >>= 1 {
		if i&1 != 0 {
			result *= a
		}
		a *= a
	}
	return result
}

//
//func merge(leftScript int, rightScript int, nums []int, leftNums []int, rightNums []int) []int {
//	// normal situation
//	var leftLocal, rightLocal, tempo = 0, 0, 0
//	shouldcopy := make([]int, len(leftNums)+len(rightNums))
//	leftHead := len(leftNums)
//	rightHead := len(rightNums)
//
//	copy(nums, shouldcopy)
//	return nums
//}
//
//func mergeSecond() {
//
//}
//
////MergeSort 又名归并排序
////https://en.wikipedia.org/wiki/Merge_sort#:~:text=In%20computer%20science%2C%20merge%20sort,in%20the%20input%20and%20output.
//func MergeSort(nums []int) []int {
//	// 1
//	numsLen := len(nums)
//	if numsLen < 2 {
//		return nums
//	}
//	// 2
//	for i := 1; exponent(2, uint64(i)) < uint64(numsLen); i++ {
//		index := exponent(2, uint64(i)) //2,4,8,16...
//		nums = func(nums []int, index uint64) []int {
//
//			turns := uint64(numsLen) / index // how many group in list
//			remain := uint64(numsLen) % index
//			for j := 0; uint64(j) < turns; j++ {
//				// 0:n-1 situation
//				leftScript := int(index) * j
//				rightScript := int(index) * (j + 1)
//				midScript := int(index)*(j+1) - int(index/2)
//
//				if remain <= index/2 && uint64(j)+1 == turns && remain != 0 { //remain smaller than before,then skip
//					break
//				} else if remain != 0 && uint64(j)+1 == turns { //remain more than last index and in last turn
//					rightScript = numsLen
//				}
//
//				leftNums := nums[leftScript:midScript]
//				rightNums := nums[midScript:rightScript]
//				nums = merge(leftScript, rightScript, nums, leftNums, rightNums)
//			}
//			return nums
//		}(nums, index)
//
//	}
//	return nums
//}

func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	m := (len(a)) / 2
	f := mergeSort(a[:m])
	s := mergeSort(a[m:])
	return merge(f, s)
}

func merge(f []int, s []int) []int {
	var i, j int
	fmt.Println(f, s)
	size := len(f) + len(s)
	a := make([]int, size, size)
	for z := 0; z < size; z++ {
		lenF := len(f)
		lenS := len(s)
		if i > lenF-1 && j <= lenS-1 {
			a[z] = s[j]
			j++
		} else if j > lenS-1 && i <= lenF-1 {
			a[z] = f[i]
			i++
		} else if f[i] < s[j] {
			a[z] = f[i]
			i++
		} else {
			a[z] = s[j]
			j++
		}
	}

	return a
}

// HeapSort 堆積排序（英語：Heapsort）是指利用堆積這種數據結構所設計的一種排序演算法。堆積是一個近似完全二叉樹的結構，
// 並同時滿足堆積的性質：即子節點的鍵值或索引總是小於（或者大於）它的父節點。
// https://zh.m.wikipedia.org/zh-hk/%E5%A0%86%E6%8E%92%E5%BA%8F
// https://www.bilibili.com/video/BV1AF411G7cA?spm_id_from=333.337.search-card.all.click
func heapSort(a []int) {

}

//

func init() {
	a := rand.Perm(16)
	fmt.Println(a)
	fmt.Println(mergeSort(a))
}
