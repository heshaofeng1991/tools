package main

import (
	"fmt"
	"sort"
)

/*
	1802. 有界数组中指定下标处的最大值

	给你三个正整数 n、index 和 maxSum 。你需要构造一个同时满足下述所有条件的数组 nums（下标 从 0 开始 计数）：

	nums.length == n
	nums[i] 是 正整数 ，其中 0 <= i < n
	abs(nums[i] - nums[i+1]) <= 1 ，其中 0 <= i < n-1
	nums 中所有元素之和不超过 maxSum
	nums[index] 的值被 最大化
	返回你所构造的数组中的 nums[index] 。

	注意：abs(x) 等于 x 的前提是 x >= 0 ；否则，abs(x) 等于 -x 。

	来源：力扣（LeetCode）
	转载链接：https://leetcode.cn/problems/maximum-value-at-a-given-index-in-a-bounded-array
*/

func maxValue(n int, index int, maxSum int) int {
	sum := func(x, cnt int) int {
		if x >= cnt {
			return (x*2 - cnt + 1) * cnt / 2
		}

		return (x+1)*x/2 + cnt - x
	}

	// sort.Search 使用二分查找查找并返回最小的索引 i
	return sort.Search(maxSum, func(x int) bool {
		x++

		return sum(x-1, index)+sum(x, n-index) > maxSum
	})
}

func main() {
	fmt.Println(maxValue(4, 2, 6))
	fmt.Println(maxValue(6, 1, 10))
}
