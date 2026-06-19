package main

// LeetCode #2593: Find Score of an Array After Marking All Elements
// https://leetcode.com/problems/find-score-of-an-array-after-marking-all-elements/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type Item struct {
	val int
	idx int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].val < h[j].val || (h[i].val == h[j].val && h[i].idx < h[j].idx) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findScore(nums []int) int64 {
	n := len(nums)
	marked := make([]bool, n)
	h := &MinHeap{}
	heap.Init(h)

	for i, v := range nums {
		heap.Push(h, Item{v, i})
	}

	var score int64
	for h.Len() > 0 {
		item := heap.Pop(h).(Item)
		val, idx := item.val, item.idx
		if marked[idx] {
			continue
		}
		score += int64(val)
		marked[idx] = true
		if idx-1 >= 0 {
			marked[idx-1] = true
		}
		if idx+1 < n {
			marked[idx+1] = true
		}
	}
	return score
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findScore([]int{2, 1, 3, 4, 5, 2}))
	// Expected: 7

	// Test case 2
	fmt.Println("Test 2:", findScore([]int{2, 3, 5, 1, 3, 2}))
	// Expected: 5

	// Test case 3
	fmt.Println("Test 3:", findScore([]int{5, 4, 3, 2, 1}))
	// Expected: 9
}
