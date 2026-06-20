# 1675 — Minimize Deviation In Array

## Deskripsi

**Soal:** [1675. Minimize Deviation In Array](https://leetcode.com/problems/minimize-deviation-in-array/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Heap (priority queue)

**Fungsi Solusi:** `func minimumDeviation(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #1675: Minimize Deviation in Array
// https://leetcode.com/problems/minimize-deviation-in-array/
// Difficulty: Hard
// Strategy: Max-heap. Multiply all odds by 2 (maximize), then reduce max even by /2.

import (
	"container/heap"
	"fmt"
)

// MaxHeap for ints
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minimumDeviation(nums []int) int {
	h := &MaxHeap{}
	heap.Init(h)
	minVal := int(1 << 60)

	// Step 1: multiply all odd numbers by 2 (they can only increase),
	// track the minimum value
	for _, num := range nums {
		if num%2 == 1 {
			num *= 2
		}
		if num < minVal {
			minVal = num
		}
		heap.Push(h, num)
	}

	ans := int(1 << 60)
	for {
		maxVal := heap.Pop(h).(int)
		ans = min(ans, maxVal-minVal)
		if maxVal%2 == 1 {
			// max is odd -> can't be reduced further
			break
		}
		maxVal /= 2
		if maxVal < minVal {
			minVal = maxVal
		}
		heap.Push(h, maxVal)
	}
	return ans
}

func main() {
	// Example 1: [1,2,3,4] -> 1
	fmt.Printf("minimumDeviation([1,2,3,4]) = %d (expected 1)\n", minimumDeviation([]int{1, 2, 3, 4}))

	// Example 2: [4,1,5,20,3] -> 3
	fmt.Printf("minimumDeviation([4,1,5,20,3]) = %d (expected 3)\n", minimumDeviation([]int{4, 1, 5, 20, 3}))

	// Additional test: [3,5] -> 1
	fmt.Printf("minimumDeviation([3,5]) = %d (expected 1)\n", minimumDeviation([]int{3, 5}))

	// Single element
	fmt.Printf("minimumDeviation([10]) = %d (expected 0)\n", minimumDeviation([]int{10}))

	// All odds (correct answer is 3 not 2)
	fmt.Printf("minimumDeviation([1,3,5]) = %d (expected 3)\n", minimumDeviation([]int{1, 3, 5}))
}
```
