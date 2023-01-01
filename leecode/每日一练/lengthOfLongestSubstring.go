package main

import (
	"fmt"
)

/*
	3. 无重复字符的最长子串

	给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

	示例 1:

	输入: s = "abcabcbb"
	输出: 3
	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
	示例 2:

	输入: s = "bbbbb"
	输出: 1
	解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
	示例 3:

	输入: s = "pwwkew"
	输出: 3
	解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
	     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

	提示：

	0 <= s.length <= 5 * 104
	s 由英文字母、数字、符号和空格组成

	来源：力扣（LeetCode）
	转载自LeeCode 链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters
*/

func lengthOfLongestSubstring(s string) int {
	res := 0
	left := -1

	for right := 0; right < len(s); right++ {
		if left == -1 {
			left = right
		} else {
			for j := right - 1; j >= left; j-- {
				if s[j] == s[right] {
					left = j + 1

					break
				}
			}
		}

		if right-left+1 > res {
			res = right - left + 1
		}
	}

	return res
}

func lengthOfLongestSubstring2(s string) int {
	start, tmp, res := 0, 0, 0

	for i := 0; i < len(s); i++ {
		tmp = i - start + 1

		for j := start; j < i; j++ {
			if s[i] == s[j] {
				tmp = i - start
				start = j + 1
				break
			}
		}

		if tmp > res {
			res = tmp
		}
	}

	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcd"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
