# 3711 — Maximum Transactions Without Negative Balance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumTransactionsWithoutNegativeBalance(transactions []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3711: Maximum Transactions Without Negative Balance
// https://leetcode.com/problems/maximum-transactions-without-negative-balance/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type minHeapInt []int

func (h minHeapInt) Len() int           { return len(h) }
func (h minHeapInt) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeapInt) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeapInt) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeapInt) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maximumTransactionsWithoutNegativeBalance(transactions []int) int {
	h := &minHeapInt{}
	heap.Init(h)
	s := 0
	ans := len(transactions)

	for _, x := range transactions {
		s += x
  // Masukkan elemen ke priority queue
		heap.Push(h, x)
		for s < 0 {
  // Ambil elemen terkecil/terbesar dari heap
			y := heap.Pop(h).(int)
			s -= y
			ans--
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumTransactionsWithoutNegativeBalance([]int{2, -5, 3, -1, -2}))
	fmt.Println(maximumTransactionsWithoutNegativeBalance([]int{-1, -2, -3}))
	fmt.Println(maximumTransactionsWithoutNegativeBalance([]int{3, -2, 3, -2, 1, -1}))
}
```
