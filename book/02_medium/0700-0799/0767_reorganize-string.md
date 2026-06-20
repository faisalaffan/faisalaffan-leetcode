# 0767 — Reorganize String

## Deskripsi

**Soal:** [0767. Reorganize String](https://leetcode.com/problems/reorganize-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log k) where k is alphabet size  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Heap (priority queue)

## Solusi Go

```go
package main

// LeetCode #767: Reorganize String
// https://leetcode.com/problems/reorganize-string/
// Difficulty: Medium
// Time: O(n log k) where k is alphabet size
// Space: O(n)

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(reorganizeString("aab"))
	fmt.Println(reorganizeString("aaab"))
}

type CharCount struct {
	char byte
	cnt  int
}

type CharHeap []CharCount

func (h CharHeap) Len() int            { return len(h) }
func (h CharHeap) Less(i, j int) bool  { return h[i].cnt > h[j].cnt }
func (h CharHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *CharHeap) Push(x interface{}) { *h = append(*h, x.(CharCount)) }
func (h *CharHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func reorganizeString(s string) string {
  // Membuat slice untuk menyimpan hasil
	freq := make([]int, 26)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	h := &CharHeap{}
	heap.Init(h)
	for i := 0; i < 26; i++ {
		if freq[i] > 0 {
			heap.Push(h, CharCount{byte(i + 'a'), freq[i]})
		}
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]byte, 0, len(s))

	for h.Len() >= 2 {
		c1 := heap.Pop(h).(CharCount)
		c2 := heap.Pop(h).(CharCount)

		result = append(result, c1.char, c2.char)

		c1.cnt--
		c2.cnt--
		if c1.cnt > 0 {
			heap.Push(h, c1)
		}
		if c2.cnt > 0 {
			heap.Push(h, c2)
		}
	}

	if h.Len() == 1 {
		c := heap.Pop(h).(CharCount)
		if c.cnt > 1 {
			return ""
		}
		result = append(result, c.char)
	}

	return string(result)
}
```
