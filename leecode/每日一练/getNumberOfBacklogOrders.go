package main

import (
	"container/heap"
	"fmt"
)

/*
	1801. 积压订单中的订单总数

	给你一个二维整数数组 orders ，其中每个 orders[i] = [pricei, amounti, orderTypei] 表示有 amounti 笔类型为 orderTypei 、价格为 pricei 的订单。

	订单类型 orderTypei 可以分为两种：

	0 表示这是一批采购订单 buy
	1 表示这是一批销售订单 sell
	注意，orders[i] 表示一批共计 amounti 笔的独立订单，这些订单的价格和类型相同。对于所有有效的 i ，由 orders[i] 表示的所有订单提交时间均早于 orders[i+1] 表示的所有订单。

	存在由未执行订单组成的 积压订单 。积压订单最初是空的。提交订单时，会发生以下情况：

	如果该订单是一笔采购订单 buy ，则可以查看积压订单中价格 最低 的销售订单 sell 。如果该销售订单 sell 的价格 低于或等于 当前采购订单 buy 的价格，则匹配并执行这两笔订单，并将销售订单 sell 从积压订单中删除。否则，采购订单 buy 将会添加到积压订单中。
	反之亦然，如果该订单是一笔销售订单 sell ，则可以查看积压订单中价格 最高 的采购订单 buy 。如果该采购订单 buy 的价格 高于或等于 当前销售订单 sell 的价格，则匹配并执行这两笔订单，并将采购订单 buy 从积压订单中删除。否则，销售订单 sell 将会添加到积压订单中。
	输入所有订单后，返回积压订单中的 订单总数 。由于数字可能很大，所以需要返回对 109 + 7 取余的结果。

	来源：力扣（LeetCode）
	链接：https://leetcode.cn/problems/number-of-orders-in-the-backlog
*/

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

func main() {
	req := make([][]int, 0)

	req = append(req, []int{10, 5, 0})
	req = append(req, []int{15, 2, 1})
	req = append(req, []int{25, 1, 1})
	req = append(req, []int{30, 4, 0})

	fmt.Println(getNumberOfBacklogOrders(req))
}
