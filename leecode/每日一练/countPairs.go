package main

import (
	"fmt"
)

// 超时时间限制
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

func countPairs2(nums []int, low, high int) (ans int) {
	arr := make(map[int]int, 0)

	for _, val := range nums {
		arr[val]++
	}

	for high++; high > 0; high >>= 1 {
		result := make(map[int]int, 0)

		for x, c := range arr {
			if high&1 > 0 {
				ans += c * arr[x^(high-1)]
			}

			if low&1 > 0 {
				ans -= c * arr[x^(low-1)]
			}

			result[x>>1] += c
		}

		arr = result
		low >>= 1
	}

	return ans / 2
}

func main() {
	fmt.Println(countPairs([]int{1, 4, 2, 7}, 2, 6))
	fmt.Println(countPairs([]int{9, 8, 4, 2, 1}, 5, 14))

	fmt.Println(countPairs2([]int{1, 4, 2, 7}, 2, 6))
	fmt.Println(countPairs2([]int{9, 8, 4, 2, 1}, 5, 14))
}
