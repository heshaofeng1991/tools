package main

import (
	"fmt"
)

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
