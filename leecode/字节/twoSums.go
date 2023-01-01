package main

import (
	"fmt"
)

/*
暴力枚举

复杂度分析

	时间复杂度：O(N^2)，其中 NN 是数组中的元素数量。最坏情况下数组中任意两个数都要被匹配一次。
	空间复杂度：O(1)。
*/
func twoSum(nums []int, target int) []int {
	result := make([]int, 0)

	for i, val := range nums {
		for j := i + 1; j < len(nums); j++ {
			if val+nums[j] == target {
				result = append(result, []int{i, j}...)

				return result
			}
		}
	}

	return result
}

/*
哈希表

复杂度分析

	时间复杂度：O(N)，其中 NN 是数组中的元素数量。对于每一个元素 x，我们可以 O(1) 寻找 target - val。
	空间复杂度：O(N)，其中 NN 是数组中的元素数量。主要为哈希表的开销。
*/
func twoSum2(nums []int, target int) []int {
	result := make([]int, 0)
	hashTable := make(map[int]int, 0)

	for i, val := range nums {
		if p, ok := hashTable[target-val]; ok {
			result = append(result, []int{p, i}...)

			return result
		}

		hashTable[val] = i
	}

	return result
}

func main() {
	fmt.Println("暴力枚举方式...")
	arr := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(arr, target))

	nums := []int{3, 2, 4}
	target = 6
	fmt.Println(twoSum(nums, target))

	nums = []int{3, 3}
	target = 6
	fmt.Println(twoSum(nums, target))

	fmt.Println("----------------")
	fmt.Println("哈希表方式...")
	arr = []int{2, 7, 11, 15}
	target = 13
	fmt.Println(twoSum2(arr, target))
}
