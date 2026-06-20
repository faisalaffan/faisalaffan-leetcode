# 2233 — Maximum Product After K Increments

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumProduct(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O(n + k log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2233: Maximum Product After K Increments
// https://leetcode.com/problems/maximum-product-after-k-increments/
// Difficulty: Medium
// Time: O(n + k log n) | Space: O(n)

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

func maximumProduct(nums []int, k int) int {
	const mod = 1_000_000_007
	h := &minHeap{}
	heap.Init(h)
	for _, v := range nums {
  // Masukkan elemen ke priority queue
		heap.Push(h, v)
	}
	for i := 0; i < k; i++ {
  // Ambil elemen terkecil/terbesar dari heap
		v := heap.Pop(h).(int)
  // Masukkan elemen ke priority queue
		heap.Push(h, v+1)
	}
	prod := 1
	for h.Len() > 0 {
  // Ambil elemen terkecil/terbesar dari heap
		prod = (prod * heap.Pop(h).(int)) % mod
	}
	return prod
}

func main() {
	// Test case 1
	fmt.Println(maximumProduct([]int{0, 4}, 5))
	// Expected: 20

	// Test case 2
	fmt.Println(maximumProduct([]int{6, 3, 3, 2}, 2))
	// Expected: 216
}
```
