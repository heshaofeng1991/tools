package main

func minOperations(nums []int, x int) int {
	length, sum := len(nums), 0

	for _, num := range nums {
		sum += num
	}

	// 所有元素之和都小于x, 直接返回-1
	if sum < x {
		return -1
	}

	// 左右求和的初始位置
	left, right := -1, 0

	// 左右和的初始值
	leftSum, rightSum := 0, sum
	result := length + 1

	// 遍历数组
	for ; left < length; left++ {
		// 左边和的累加
		if left != -1 {
			leftSum += nums[left]
		}

		// 右边和的累减
		for right < length && leftSum+rightSum > x {
			rightSum -= nums[right]

			right++
		}

		if leftSum+rightSum == x {
			result = min(result, (left+1)+(length-right))
		}
	}

	if result > length {
		return -1
	}

	return result
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}
