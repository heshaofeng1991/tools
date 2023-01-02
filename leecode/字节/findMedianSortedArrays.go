package main

import (
	"fmt"
	"sort"
)

/*

	4. 寻找两个正序数组的中位数

	给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

	算法的时间复杂度应该为 O(log (m+n)) 。

	示例 1：

	输入：nums1 = [1,3], nums2 = [2]
	输出：2.00000
	解释：合并数组 = [1,2,3] ，中位数 2
	示例 2：

	输入：nums1 = [1,2], nums2 = [3,4]
	输出：2.50000
	解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

	提示：

	nums1.length == m
	nums2.length == n
	0 <= m <= 1000
	0 <= n <= 1000
	1 <= m + n <= 2000
	-106 <= nums1[i], nums2[i] <= 106

	转载自：来源：力扣（LeetCode）
*/

/*
题解1：
 1. 数组合并
 2. 对合并数组进行排序处理
 3. 针对排序后的数组进行奇偶取中位数处理
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := make([]int, 0)

	// 1. 数组合并
	nums = append(nums, nums1...)
	nums = append(nums, nums2...)

	// 2. 对合并数组进行排序处理
	sort.Ints(nums)

	// 3. 针对排序后的数组进行奇偶取中位数处理
	length := len(nums)

	switch length % 2 {
	case 0:
		return float64(nums[length/2]+nums[length/2-1]) / 2.0
	default:
		return float64(nums[length/2])
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

/*
题解2：

	二分查找处理
*/
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)

	if length%2 == 1 {
		mid := length / 2

		return float64(getKthElement(nums1, nums2, mid+1))
	} else {
		left, right := length/2-1, length/2

		return float64(getKthElement(nums1, nums2, left+1)+getKthElement(nums1, nums2, right+1)) / 2.0
	}

	return 0.00
}

func getKthElement(nums1, nums2 []int, k int) int {
	left, right := 0, 0

	for {
		if left == len(nums1) {
			return nums2[right+k-1]
		}

		if right == len(nums2) {
			return nums1[left+k-1]
		}

		if k == 1 {
			return min(nums1[left], nums2[right])
		}

		mid := k / 2
		index1 := min(left+mid, len(nums1)) - 1
		index2 := min(right+mid, len(nums2)) - 1
		val1, val2 := nums1[index1], nums2[index2]

		if val1 <= val2 {
			k -= (index1 - left + 1)
			left = index1 + 1
		} else {
			k -= (index2 - right + 1)
			right = index2 + 1
		}
	}

	return 0
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))

	fmt.Println(findMedianSortedArrays2([]int{1, 2}, []int{3}))
	fmt.Println(findMedianSortedArrays2([]int{1, 2}, []int{3, 4}))
}
