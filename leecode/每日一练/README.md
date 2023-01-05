# 每日一练

[LeeCode 2351. 第一个出现两次的字母](https://leetcode.cn/problems/first-letter-to-appear-twice/)<br>
- [第一个出现两次的字母 repeatedCharacter.go](https://github.com/heshaofeng1991/tools/blob/master/leecode/%E6%AF%8F%E6%97%A5%E4%B8%80%E7%BB%83/repeatedCharacter.go)<br>
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
[LeeCode 3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)<br>
- [无重复字符的最长子串 lengthOfLongestSubstring.go](https://github.com/heshaofeng1991/tools/blob/master/leecode/%E6%AF%8F%E6%97%A5%E4%B8%80%E7%BB%83/lengthOfLongestSubstring.go)<br>
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
[LeeCode 1801. 积压订单中的订单总数](https://leetcode.cn/problems/number-of-orders-in-the-backlog/)<br>
- [积压订单中的订单总数 getNumberOfBacklogOrders.go](https://github.com/heshaofeng1991/tools/blob/master/leecode/%E6%AF%8F%E6%97%A5%E4%B8%80%E7%BB%83/getNumberOfBacklogOrders.go)<br>
```go
type pair struct{ price, left int }
type buy []pair

func (h buy) Len() int            { return len(h) }
func (h buy) Less(i, j int) bool  { return h[i].price > h[j].price }
func (h buy) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *buy) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *buy) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type sell []pair

func (h sell) Len() int            { return len(h) }
func (h sell) Less(i, j int) bool  { return h[i].price < h[j].price }
func (h sell) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *sell) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *sell) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

const MOD = 1e9 + 7

func getNumberOfBacklogOrders(orders [][]int) (ans int) {
	buyOrders, sellOrders := buy{}, sell{}

	for _, ord := range orders {
		price, amount := ord[0], ord[1]

		if ord[2] == 0 {
			for amount > 0 && len(sellOrders) > 0 && sellOrders[0].price <= price {
				if sellOrders[0].left > amount {
					sellOrders[0].left -= amount
					amount = 0

					break
				}

				amount -= heap.Pop(&sellOrders).(pair).left
			}

			if amount > 0 {
				heap.Push(&buyOrders, pair{price, amount})
			}
		} else {
			for amount > 0 && len(buyOrders) > 0 && buyOrders[0].price >= price {
				if buyOrders[0].left > amount {
					buyOrders[0].left -= amount
					amount = 0

					break
				}

				amount -= heap.Pop(&buyOrders).(pair).left
			}

			if amount > 0 {
				heap.Push(&sellOrders, pair{price, amount})
			}
		}
	}

	for _, p := range buyOrders {
		ans += p.left
	}

	for _, p := range sellOrders {
		ans += p.left
	}

	return ans % MOD
}
```

[LeeCode 2042. 检查句子中的数字是否递增](https://leetcode.cn/problems/check-if-numbers-are-ascending-in-a-sentence/)<br>
- [检查句子中的数字是否递增 areNumbersAscending.go](https://github.com/heshaofeng1991/tools/blob/master/leecode/%E6%AF%8F%E6%97%A5%E4%B8%80%E7%BB%83/areNumbersAscending.go)<br>
```go
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
```
[1802. 有界数组中指定下标处的最大值](https://leetcode.cn/problems/maximum-value-at-a-given-index-in-a-bounded-array/)<br>
- [有界数组中指定下标处的最大值 maxValue.go](https://github.com/heshaofeng1991/tools/blob/master/leecode/%E6%AF%8F%E6%97%A5%E4%B8%80%E7%BB%83/maxValue.go)<br>
```go
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
```
[1803. 统计异或值在范围内的数对有多少](https://leetcode.cn/problems/count-pairs-with-xor-in-a-range/) <br>
 - [countPairs](https://github.com/heshaofeng1991/tools/blob/master/leecode/%E6%AF%8F%E6%97%A5%E4%B8%80%E7%BB%83/countPairs.go)<br>
```go
// 超出时间限制
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
```