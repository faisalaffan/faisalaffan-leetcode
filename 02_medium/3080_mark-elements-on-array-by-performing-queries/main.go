package main

// LeetCode #3080: Mark Elements on Array by Performing Queries
// https://leetcode.com/problems/mark-elements-on-array-by-performing-queries/
// Difficulty: Medium
// Time: O((n + q) log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type pair3080 struct {
	val int
	idx int
}
type minHeap3080 []pair3080

func (h minHeap3080) Len() int            { return len(h) }
func (h minHeap3080) Less(i, j int) bool  { return h[i].val < h[j].val || (h[i].val == h[j].val && h[i].idx < h[j].idx) }
func (h minHeap3080) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap3080) Push(x any)         { *h = append(*h, x.(pair3080)) }
func (h *minHeap3080) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	fmt.Println(unmarkedSumArray([]int{1, 2, 2, 1, 2, 3, 1}, [][]int{{1, 2}, {3, 3}, {4, 2}}))
	fmt.Println(unmarkedSumArray([]int{1, 4, 2, 3}, [][]int{{0, 1}}))
}

func unmarkedSumArray(nums []int, queries [][]int) []int64 {
	n := len(nums)
	sum := int64(0)
	h := &minHeap3080{}
	heap.Init(h)
	for i, x := range nums {
		sum += int64(x)
		heap.Push(h, pair3080{x, i})
	}
	marked := make([]bool, n)
	ans := make([]int64, len(queries))
	for qi, q := range queries {
		idx, k := q[0], q[1]
		if !marked[idx] {
			marked[idx] = true
			sum -= int64(nums[idx])
		}
		for k > 0 && h.Len() > 0 {
			p := heap.Pop(h).(pair3080)
			if !marked[p.idx] {
				marked[p.idx] = true
				sum -= int64(p.val)
				k--
			}
		}
		ans[qi] = sum
	}
	return ans
}
