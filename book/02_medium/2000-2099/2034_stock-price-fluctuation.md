# 2034 — Stock Price Fluctuation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() StockPrice
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O(log n) per operation  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2034: Stock Price Fluctuation
// https://leetcode.com/problems/stock-price-fluctuation/
// Difficulty: Medium
// Time: O(log n) per operation | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type StockPrice struct {
	prices   map[int]int
	latestTs int
	latestP  int
	minHeap  *MinHeap
	maxHeap  *MaxHeap
}

type PriceEntry struct {
	price int
	ts    int
}

type MinHeap []PriceEntry
func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].price < h[j].price }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(PriceEntry)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MaxHeap []PriceEntry
func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].price > h[j].price }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(PriceEntry)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func Constructor() StockPrice {
	return StockPrice{
		prices:  make(map[int]int),
		minHeap: &MinHeap{},
		maxHeap: &MaxHeap{},
	}
}

func (sp *StockPrice) Update(timestamp int, price int) {
	sp.prices[timestamp] = price
	if timestamp >= sp.latestTs {
		sp.latestTs = timestamp
		sp.latestP = price
	}
  // Masukkan elemen ke priority queue
	heap.Push(sp.minHeap, PriceEntry{price, timestamp})
  // Masukkan elemen ke priority queue
	heap.Push(sp.maxHeap, PriceEntry{price, timestamp})
}

func (sp *StockPrice) Current() int {
	return sp.latestP
}

func (sp *StockPrice) Maximum() int {
	for sp.maxHeap.Len() > 0 {
		top := (*sp.maxHeap)[0]
		if sp.prices[top.ts] == top.price {
			return top.price
		}
  // Ambil elemen terkecil/terbesar dari heap
		heap.Pop(sp.maxHeap)
	}
	return 0
}

func (sp *StockPrice) Minimum() int {
	for sp.minHeap.Len() > 0 {
		top := (*sp.minHeap)[0]
		if sp.prices[top.ts] == top.price {
			return top.price
		}
  // Ambil elemen terkecil/terbesar dari heap
		heap.Pop(sp.minHeap)
	}
	return 0
}

func main() {
	sp := Constructor()
	sp.Update(1, 10)
	sp.Update(2, 5)
	fmt.Println("Test 1 Current:", sp.Current())  // 5
	fmt.Println("Test 1 Maximum:", sp.Maximum())  // 10
	fmt.Println("Test 1 Minimum:", sp.Minimum())  // 5

	sp2 := Constructor()
	sp2.Update(1, 1)
	sp2.Update(2, 2)
	sp2.Update(3, 3)
	sp2.Update(1, 4)
	fmt.Println("Test 2 Maximum:", sp2.Maximum()) // 4
	fmt.Println("Test 2 Minimum:", sp2.Minimum()) // 2
}
```
