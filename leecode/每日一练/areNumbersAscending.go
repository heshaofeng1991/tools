package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
	2042. 检查句子中的数字是否递增

    句子是由若干 token 组成的一个列表，token 间用 单个 空格分隔，句子没有前导或尾随空格。每个 token 要么是一个由数字 0-9 组成的不含前导零的 正整数 ，要么是一个由小写英文字母组成的 单词 。

	示例，"a puppy has 2 eyes 4 legs" 是一个由 7 个 token 组成的句子："2" 和 "4" 是数字，其他像 "puppy" 这样的 tokens 属于单词。
	给你一个表示句子的字符串 s ，你需要检查 s 中的 全部 数字是否从左到右严格递增（即，除了最后一个数字，s 中的 每个 数字都严格小于它 右侧 的数字）。

	如果满足题目要求，返回 true ，否则，返回 false 。

	来源：力扣（LeetCode）
	转载链接：https://leetcode.cn/problems/check-if-numbers-are-ascending-in-a-sentence
*/

func areNumbersAscending(s string) bool {
	str := make([]string, 0)
	res := make([]int, 0)

	str = strings.Split(s, " ")

	for i := 0; i <= len(str)-1; i++ {
		if str[i] == "0" {
			val, _ := strconv.Atoi(str[i])

			res = append(res, val)

			continue
		}

		val, err := strconv.Atoi(str[i])
		if err != nil {
			continue
		}

		res = append(res, val)
	}

	if len(res) == 0 {
		return false
	}

	if containsDuplicate(res) {
		return false
	}

	return sort.IntsAreSorted(res)
}

func containsDuplicate(nums []int) bool {
	mp := make(map[int]bool)

	for _, v := range nums {
		if _, ok := mp[v]; ok {
			return true
		}

		mp[v] = false
	}

	return false
}

func main() {
	fmt.Println(areNumbersAscending("1 box has 3 blue 4 red 6 green and 12 yellow marbles"))
	fmt.Println(areNumbersAscending("hello world 5 x 5"))
	fmt.Println(areNumbersAscending("sunset is at 7 51 pm overnight lows will be in the low 50 and 60"))
	fmt.Println(areNumbersAscending("4 5 11 26"))
	fmt.Println(areNumbersAscending("4 5 5 11 26"))
	fmt.Println(areNumbersAscending("36 claim 37 38 39 39 41 hire final 42 43 twist shift young 44 miss 45 46 sad 47 48 dig 49 50 green 51 train 52 broad 53"))
}
