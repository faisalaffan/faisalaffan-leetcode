package main

// LeetCode #2542: Maximum Subsequence Score
// https://leetcode.com/problems/maximum-subsequence-score/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
	"sort"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	pairs := make([][2]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = [2]int{nums2[i], nums1[i]}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})

	h := &MinHeap{}
	heap.Init(h)
	var sum int64
	var ans int64

	for _, p := range pairs {
		val1, val2 := p[1], p[0]
		sum += int64(val1)
		heap.Push(h, val1)
		if h.Len() > k {
			sum -= int64(heap.Pop(h).(int))
		}
		if h.Len() == k {
			score := sum * int64(val2)
			if score > ans {
				ans = score
			}
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxScore([]int{1, 3, 3, 2}, []int{2, 1, 3, 4}, 3))
	// Expected: 12

	// Test case 2
	fmt.Println("Test 2:", maxScore([]int{4, 2, 3, 1, 1}, []int{7, 5, 10, 9, 6}, 1))
	// Expected: 30

	// Test case 3
	fmt.Println("Test 3:", maxScore([]int{2, 1, 14, 12}, []int{11, 7, 13, 6}, 3))
	// Expected: 168
}
