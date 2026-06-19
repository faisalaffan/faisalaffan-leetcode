package main

// LeetCode #358: Rearrange String k Distance Apart
// https://leetcode.com/problems/rearrange-string-k-distance-apart/
// Difficulty: Hard [Paid]
//
// Use a max-heap (by frequency) to always pick the most frequent available
// character. A cooldown queue enforces that a character cannot be reused
// within k distance. If the heap empties before the string is built, it's
// impossible and we return "".

import (
	"container/heap"
	"fmt"
)

func main() {
	// Example 1: "aabbcc", k=3 -> "abcabc" (or any valid)
	fmt.Println(rearrangeString("aabbcc", 3))
	// Example 2: "aaabc", k=3 -> "" (impossible)
	fmt.Println(rearrangeString("aaabc", 3))
	// Example 3: "aaadbbcc", k=2 -> "abacabcd" (or valid)
	fmt.Println(rearrangeString("aaadbbcc", 2))
	// Edge: k=0
	fmt.Println(rearrangeString("aabb", 0))
	// Edge: single char
	fmt.Println(rearrangeString("a", 1))
}

type charFreq struct {
	ch    byte
	count int
}

type maxHeap []charFreq

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i].count > h[j].count }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(charFreq)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type cooldownItem struct {
	ch      byte
	count   int
	readyAt int // position when this char becomes available again
}

func rearrangeString(s string, k int) string {
	if k <= 1 {
		return s
	}

	// Count frequencies
	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	// Build max-heap
	h := &maxHeap{}
	heap.Init(h)
	for ch, cnt := range freq {
		heap.Push(h, charFreq{ch, cnt})
	}

	result := make([]byte, 0, len(s))
	q := make([]cooldownItem, 0)

	for len(result) < len(s) {
		// Replenish cooldown items whose readyAt <= current position
		if len(q) > 0 && q[0].readyAt <= len(result) {
			item := q[0]
			q = q[1:]
			heap.Push(h, charFreq{item.ch, item.count})
		}

		if h.Len() == 0 {
			return "" // impossible
		}

		cf := heap.Pop(h).(charFreq)
		result = append(result, cf.ch)
		cf.count--
		if cf.count > 0 {
			q = append(q, cooldownItem{cf.ch, cf.count, len(result) - 1 + k})
		}
	}

	return string(result)
}
