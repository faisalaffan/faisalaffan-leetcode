# 0767 — Reorganize String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func reorganizeString(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O(n log k) where k is alphabet size  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Alokasi slice integer
	freq := make([]int, 26)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	h := &CharHeap{}
	heap.Init(h)
	for i := 0; i < 26; i++ {
		if freq[i] > 0 {
  // Masukkan elemen ke priority queue
			heap.Push(h, CharCount{byte(i + 'a'), freq[i]})
		}
	}

	result := make([]byte, 0, len(s))

	for h.Len() >= 2 {
  // Ambil elemen terkecil/terbesar dari heap
		c1 := heap.Pop(h).(CharCount)
  // Ambil elemen terkecil/terbesar dari heap
		c2 := heap.Pop(h).(CharCount)

		result = append(result, c1.char, c2.char)

		c1.cnt--
		c2.cnt--
		if c1.cnt > 0 {
  // Masukkan elemen ke priority queue
			heap.Push(h, c1)
		}
		if c2.cnt > 0 {
  // Masukkan elemen ke priority queue
			heap.Push(h, c2)
		}
	}

	if h.Len() == 1 {
  // Ambil elemen terkecil/terbesar dari heap
		c := heap.Pop(h).(CharCount)
		if c.cnt > 1 {
			return ""
		}
		result = append(result, c.char)
	}

	return string(result)
}
```
