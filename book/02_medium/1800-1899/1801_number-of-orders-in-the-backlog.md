# 1801 — Number Of Orders In The Backlog

## Deskripsi

**Soal:** [1801. Number Of Orders In The Backlog](https://leetcode.com/problems/number-of-orders-in-the-backlog/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Heap (priority queue)

**Fungsi Solusi:** `func getNumberOfBacklogOrders(orders [][]int) int`

## Solusi Go

```go
package main

// LeetCode #1801: Number of Orders in the Backlog
// https://leetcode.com/problems/number-of-orders-in-the-backlog/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"container/heap"
	"fmt"
)

type Order struct {
	price  int
	amount int
}

type MaxHeap []Order

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].price > h[j].price }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Order)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MinHeap []Order

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].price < h[j].price }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Order)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

const mod = 1_000_000_007

func getNumberOfBacklogOrders(orders [][]int) int {
	buys := &MaxHeap{}
	sells := &MinHeap{}
	heap.Init(buys)
	heap.Init(sells)

	for _, o := range orders {
		price, amount, orderType := o[0], o[1], o[2]
		if orderType == 0 { // buy
			for amount > 0 && sells.Len() > 0 && (*sells)[0].price <= price {
				top := &(*sells)[0]
				if top.amount > amount {
					top.amount -= amount
					amount = 0
				} else {
					amount -= top.amount
					heap.Pop(sells)
				}
			}
			if amount > 0 {
				heap.Push(buys, Order{price, amount})
			}
		} else { // sell
			for amount > 0 && buys.Len() > 0 && (*buys)[0].price >= price {
				top := &(*buys)[0]
				if top.amount > amount {
					top.amount -= amount
					amount = 0
				} else {
					amount -= top.amount
					heap.Pop(buys)
				}
			}
			if amount > 0 {
				heap.Push(sells, Order{price, amount})
			}
		}
	}

	total := 0
	for buys.Len() > 0 {
		total = (total + heap.Pop(buys).(Order).amount) % mod
	}
	for sells.Len() > 0 {
		total = (total + heap.Pop(sells).(Order).amount) % mod
	}
	return total
}

func main() {
	fmt.Println(getNumberOfBacklogOrders([][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}})) // Expected: 6
	fmt.Println(getNumberOfBacklogOrders([][]int{{7, 1000000000, 1}, {15, 3, 0}, {5, 999999995, 0}, {5, 1, 1}})) // Expected: 999999984
}
```
