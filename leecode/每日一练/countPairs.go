package main

import (
	"fmt"
)

// 超时，时间限制
func countPairs(nums []int, low int, high int) int {
	length := len(nums)
	i, j, x, result := 0, 0, 0, 0

	for i = 0; i < length; i++ {
		for j = i + 1; j < length; j++ {
			x = nums[i] ^ nums[j]

			if x >= low && x <= high {
				result++
			}
		}
	}

	return result
}

func main() {
	fmt.Println(countPairs([]int{1, 4, 2, 7}, 2, 6))
	fmt.Println(countPairs([]int{9, 8, 4, 2, 1}, 5, 14))
}
