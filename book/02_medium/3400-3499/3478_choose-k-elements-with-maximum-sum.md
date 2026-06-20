# 3478 — Choose K Elements With Maximum Sum

## Deskripsi

**Soal:** [3478. Choose K Elements With Maximum Sum](https://leetcode.com/problems/choose-k-elements-with-maximum-sum/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Heap (priority queue)

## Solusi Go

```go
package main

// LeetCode #3478: Choose K Elements With Maximum Sum
// https://leetcode.com/problems/choose-k-elements-with-maximum-sum/
// Difficulty: Medium
// Complexity: O(n log n) time, O(n) space

import (
	"container/heap"
	"fmt"
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

func main() {
	// Test case 1
	fmt.Println("Test 1:", ChooseKElementsWithMaximumSum([]int{1, 2, 3, 4}, 2))
	// Test case 2
	fmt.Println("Test 2:", ChooseKElementsWithMaximumSum([]int{5, 3, 1, 2}, 3))
	// Test case 3
	fmt.Println("Test 3:", ChooseKElementsWithMaximumSum([]int{1, 1, 1}, 2))
}

func ChooseKElementsWithMaximumSum(nums []int, k int) int {
	if k >= len(nums) {
		sum := 0
		for _, v := range nums {
			sum += v
		}
		return sum
	}
	h := &MinHeap{}
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	sum := 0
	for h.Len() > 0 {
		sum += heap.Pop(h).(int)
	}
	return sum
}
```
