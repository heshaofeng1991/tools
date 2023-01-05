# 字节

[LeeCode 1. 两数之和](https://leetcode.cn/problems/two-sum/)<br>
- [两数之和 twoSum.go](https://github.com/heshaofeng1991/tools/blob/master/leecode/%E5%AD%97%E8%8A%82/twoSums.go)<br>
```go
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
```
[LeeCode 4. 寻找两个正序数组的中位数](https://leetcode.cn/problems/median-of-two-sorted-arrays/)<br>
- [寻找两个正序数组的中位数 findMedianSortedArrays.go](https://github.com/heshaofeng1991/tools/blob/master/leecode/%E5%AD%97%E8%8A%82/findMedianSortedArrays.go)<br>
```go
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
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
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
```