package main

// LeetCode #3377: Digit Operations to Make Two Integers Equal
// https://leetcode.com/problems/digit-operations-to-make-two-integers-equal/
// Difficulty: Medium
// Time: O(N log N) Space: O(N)

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(minOperations(10, 12)) // 6 (not 10, paths: 10+11+12=33, 10+12=22 if allowed)
}

type Item struct {
	cost int
	num  int
}

type ItemHeap []Item

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ItemHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

const MAX = 10000

var isPrime []bool

func init() {
	isPrime = make([]bool, MAX+1)
	for i := 2; i <= MAX; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= MAX; i++ {
		if isPrime[i] {
			for j := i * i; j <= MAX; j += i {
				isPrime[j] = false
			}
		}
	}
}

func minOperations(n int, m int) int {
	if n == m {
		return n
	}
	if isPrime[n] || isPrime[m] {
		return -1
	}

	dist := make([]int, MAX+1)
	for i := range dist {
		dist[i] = 1 << 60
	}
	dist[n] = n

	h := &ItemHeap{}
	heap.Init(h)
	heap.Push(h, Item{n, n})

	for h.Len() > 0 {
		item := heap.Pop(h).(Item)
		if item.cost > dist[item.num] {
			continue
		}
		if item.num == m {
			return item.cost
		}

		// Generate neighbors by changing each digit
		digits := getDigits(item.num)
		for pos := 0; pos < len(digits); pos++ {
			orig := digits[pos]

			// Increment digit
			if orig < 9 {
				digits[pos] = orig + 1
				next := fromDigits(digits)
				if !isPrime[next] {
					nc := item.cost + next
					if nc < dist[next] {
						dist[next] = nc
						heap.Push(h, Item{nc, next})
					}
				}
			}

			// Decrement digit
			if orig > 0 && !(pos == 0 && orig == 1) {
				// Can't decrement to leading zero
				if pos > 0 || orig > 1 {
					digits[pos] = orig - 1
					next := fromDigits(digits)
					if !isPrime[next] {
						nc := item.cost + next
						if nc < dist[next] {
							dist[next] = nc
							heap.Push(h, Item{nc, next})
						}
					}
				}
			}

			digits[pos] = orig
		}
	}
	return -1
}

func getDigits(num int) []int {
	var d []int
	for num > 0 {
		d = append([]int{num % 10}, d...)
		num /= 10
	}
	return d
}

func fromDigits(d []int) int {
	r := 0
	for _, v := range d {
		r = r*10 + v
	}
	return r
}
