# 0692 — Top K Frequent Words

## Deskripsi

**Soal:** [0692. Top K Frequent Words](https://leetcode.com/problems/top-k-frequent-words/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log k)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Heap (priority queue)

## Solusi Go

```go
package main

// LeetCode #692: Top K Frequent Words
// https://leetcode.com/problems/top-k-frequent-words/
// Difficulty: Medium
// Time: O(n log k)
// Space: O(n)

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
	fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))
}

type Item struct {
	word string
	freq int
}

type MinHeap []Item

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].freq == h[j].freq {
		return h[i].word > h[j].word
	}
	return h[i].freq < h[j].freq
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for word, f := range freq {
		heap.Push(h, Item{word, f})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(Item).word
	}

	return result
}
```
