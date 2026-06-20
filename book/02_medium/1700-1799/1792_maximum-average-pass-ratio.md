# 1792 — Maximum Average Pass Ratio

## Deskripsi

**Soal:** [1792. Maximum Average Pass Ratio](https://leetcode.com/problems/maximum-average-pass-ratio/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O((n+k) log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Heap (priority queue)

**Fungsi Solusi:** `func maxAverageRatio(classes [][]int, extraStudents int) float64`

## Solusi Go

```go
package main

// LeetCode #1792: Maximum Average Pass Ratio
// https://leetcode.com/problems/maximum-average-pass-ratio/
// Difficulty: Medium
// Time: O((n+k) log n), Space: O(n)

import (
	"container/heap"
	"fmt"
)

type Class struct {
	pass, total int
	gain        float64
}

type MaxHeap []Class

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].gain > h[j].gain }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Class)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := &MaxHeap{}
	heap.Init(h)

	for _, c := range classes {
		pass, total := c[0], c[1]
		gain := float64(pass+1)/float64(total+1) - float64(pass)/float64(total)
		heap.Push(h, Class{pass, total, gain})
	}

	for i := 0; i < extraStudents; i++ {
		c := heap.Pop(h).(Class)
		c.pass++
		c.total++
		c.gain = float64(c.pass+1)/float64(c.total+1) - float64(c.pass)/float64(c.total)
		heap.Push(h, c)
	}

	sum := 0.0
	for h.Len() > 0 {
		c := heap.Pop(h).(Class)
		sum += float64(c.pass) / float64(c.total)
	}
	return sum / float64(len(classes))
}

func main() {
	fmt.Printf("%.5f\n", maxAverageRatio([][]int{{1, 2}, {3, 5}, {2, 2}}, 2)) // Expected: 0.78333
	fmt.Printf("%.5f\n", maxAverageRatio([][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}}, 4)) // Expected: 0.53485
	fmt.Printf("%.5f\n", maxAverageRatio([][]int{{1, 2}}, 1)) // Expected: 0.66667
}
```
