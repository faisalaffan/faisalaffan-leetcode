# 0767 — Reorganize String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func reorganizeString(s string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Heap

**Waktu:** O(n log k) where k is alphabet size  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Heap** — sering muncul di interview!

## 💻 Solusi Go

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
  // Alokasi slice
	freq := make([]int, 26)
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	h := &CharHeap{}
	heap.Init(h)
	for i := 0; i < 26; i++ {
		if freq[i] > 0 {
  // Push ke priority queue
			heap.Push(h, CharCount{byte(i + 'a'), freq[i]})
		}
	}

	result := make([]byte, 0, len(s))

	for h.Len() >= 2 {
  // Pop dari priority queue
		c1 := heap.Pop(h).(CharCount)
  // Pop dari priority queue
		c2 := heap.Pop(h).(CharCount)

		result = append(result, c1.char, c2.char)

		c1.cnt--
		c2.cnt--
		if c1.cnt > 0 {
  // Push ke priority queue
			heap.Push(h, c1)
		}
		if c2.cnt > 0 {
  // Push ke priority queue
			heap.Push(h, c2)
		}
	}

	if h.Len() == 1 {
  // Pop dari priority queue
		c := heap.Pop(h).(CharCount)
		if c.cnt > 1 {
			return ""
		}
		result = append(result, c.char)
	}

	return string(result)
}
```
