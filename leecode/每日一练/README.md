# 每日一练

[LeeCode 2351. 第一个出现两次的字母 ](https://leetcode.cn/problems/first-letter-to-appear-twice/)<br>
```go
// 题解思路1：总共就26个字母，利用ASCII码表，统计每次字母字符出现的次数。

func repeatedCharacter1(s string) byte {
	var result [26]int

	for _, val := range s {
		result[val-'a']++

		if result[val-'a'] == 2 {
			return byte(val)
		}
	}

	return 0
}


// 题解思路2：golang字符串每个字符有个rune，利用golang map特性，统计字母字符出现的频次。

func repeatedCharacter(s string) byte {
	result := make(map[rune]bool, 0)

	for _, val := range s {
		if result[val] {
			return byte(val)
		}

		result[val] = true
	}

	return 0
}
```
[LeeCode 3. 无重复字符的最长子串 ](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)<br>
```go
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
```