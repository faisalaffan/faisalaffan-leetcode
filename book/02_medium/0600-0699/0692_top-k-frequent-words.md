# 0692 — Top K Frequent Words

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func topKFrequent(words []string, k int) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Heap

**Waktu:** O(n log k)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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
  // HashMap: O(1) lookup
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for word, f := range freq {
  // Push ke priority queue
		heap.Push(h, Item{word, f})
		if h.Len() > k {
  // Pop dari priority queue
			heap.Pop(h)
		}
	}

	result := make([]string, k)
	for i := k - 1; i >= 0; i-- {
  // Pop dari priority queue
		result[i] = heap.Pop(h).(Item).word
	}

	return result
}
```
