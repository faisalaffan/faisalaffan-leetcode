# 3066 — Minimum Operations To Exceed Threshold Value Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minOperations3066(nums []int, k int) (ans int)
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3066: Minimum Operations to Exceed Threshold Value II
// https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-ii/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	fmt.Println(minOperations3066([]int{2, 11, 10, 1, 3}, 10))
	fmt.Println(minOperations3066([]int{1, 1, 2, 4, 9}, 20))
}

func minOperations3066(nums []int, k int) (ans int) {
	h := &minHeap{}
	heap.Init(h)
	for _, x := range nums {
		if x < k {
  // Masukkan elemen ke priority queue
			heap.Push(h, x)
		}
	}
	for h.Len() >= 2 {
  // Ambil elemen terkecil/terbesar dari heap
		x := heap.Pop(h).(int)
  // Ambil elemen terkecil/terbesar dari heap
		y := heap.Pop(h).(int)
		val := x*2 + y
		if val < k {
  // Masukkan elemen ke priority queue
			heap.Push(h, val)
		}
		ans++
	}
	if h.Len() > 0 {
		ans++
	}
	return
}
```
