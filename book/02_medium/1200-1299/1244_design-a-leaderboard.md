# 1244 — Design A Leaderboard

## Deskripsi

**Soal:** [1244. Design A Leaderboard](https://leetcode.com/problems/design-a-leaderboard/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** addScore O(1), top O(K log n), reset O(1)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Heap (priority queue)

**Fungsi Solusi:** `func Constructor() Leaderboard`

## Solusi Go

```go
package main

import (
	"container/heap"
	"fmt"
)

// LeetCode #1244: Design A Leaderboard
// https://leetcode.com/problems/design-a-leaderboard/
// Difficulty: Medium [Paid]

// Leaderboard supports addScore, top(K), and reset.

// Time: addScore O(1), top O(K log n), reset O(1)
// Space: O(n)

type Leaderboard struct {
	scores map[int]int
}

func Constructor() Leaderboard {
	return Leaderboard{scores: make(map[int]int)}
}

func (lb *Leaderboard) AddScore(playerId int, score int) {
	lb.scores[playerId] += score
}

func (lb *Leaderboard) Top(K int) int {
	// Use min-heap of size K
	h := &minHeap{}
	heap.Init(h)
	for _, s := range lb.scores {
		heap.Push(h, s)
		if h.Len() > K {
			heap.Pop(h)
		}
	}
	sum := 0
	for h.Len() > 0 {
		sum += heap.Pop(h).(int)
	}
	return sum
}

func (lb *Leaderboard) Reset(playerId int) {
	delete(lb.scores, playerId)
}

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	lb := Constructor()
	lb.AddScore(1, 73)
	lb.AddScore(2, 56)
	lb.AddScore(3, 39)
	lb.AddScore(4, 51)
	lb.AddScore(5, 4)
	fmt.Printf("top(1) = %d (expected: 73)\n", lb.Top(1))
	lb.Reset(1)
	lb.Reset(2)
	lb.AddScore(2, 51)
	fmt.Printf("top(3) = %d\n", lb.Top(3))
}
```
