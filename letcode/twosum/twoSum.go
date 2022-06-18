package twosum

import (
	"fmt"
)

func TwoSum() {
	list := []int{3, 3}
	target := 6

	list1 := []int{2, 7, 11, 15}
	target1 := 9

	list2 := []int{3, 2, 4}
	target2 := 6
	fmt.Println(SecondTwoSum(list, target))
	fmt.Println(SecondTwoSum(list1, target1))
	fmt.Println(SecondTwoSum(list2, target2))
}

//TODO
/* Author zsj
 * Desc n*n复杂度 暴力求解
 * Return
 */
func FirstTwoSum(nums []int, target int) []int {
	var list []int
	for key, value := range nums {
		sub := target - value
		for i := key + 1; i < len(nums); i++ {
			if nums[i] == sub {
				return []int{key, i}
			}
		}
	}
	return list
}

//TODO
/* Author zsj
 * Desc map解法
 * Return
 */
func SecondTwoSum(nums []int, target int) []int {
	//TODO
	/* 1 创建大小为num个数大小的map
	 * 2 target-num[i]=sub 如果map中有sub return数组下标
	 * 3 如果map中没有对应下标 则继续
	 * Return 两个对应数组下标
	 */
	Map := make(map[int]int, len(nums)+1)
	Map[nums[0]] = 1

	for i := 1; i < len(nums); i++ {
		sub := target - nums[i]

		if Map[sub] != 0 {
			return []int{i, Map[sub] - 1}
		}

		Map[nums[i]] = i + 1

	}
	var list []int
	return list
}
