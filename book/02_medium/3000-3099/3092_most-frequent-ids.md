# 3092 — Most Frequent Ids

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func mostFrequentIDs(nums []int, freq []int) []int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3092: Most Frequent IDs
// https://leetcode.com/problems/most-frequent-ids/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type Item struct {
	id    int
	count int64
	idx   int
}

type MaxHeap []*Item

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].idx = i
	h[j].idx = j
}
func (h *MaxHeap) Push(x any) {
	n := len(*h)
	item := x.(*Item)
	item.idx = n
	*h = append(*h, item)
}
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.idx = -1
	*h = old[:n-1]
	return item
}

func mostFrequentIDs(nums []int, freq []int) []int64 {
	n := len(nums)
  // Alokasi slice integer
	ans := make([]int64, n)
  // Membuat map (HashMap) — pencarian O(1)
	counts := make(map[int]int64)
	h := &MaxHeap{}
	heap.Init(h)

	for i := 0; i < n; i++ {
		id := nums[i]
		counts[id] += int64(freq[i])
  // Masukkan elemen ke priority queue
		heap.Push(h, &Item{id: id, count: counts[id]})
		for h.Len() > 0 && (*h)[0].count != counts[(*h)[0].id] {
  // Ambil elemen terkecil/terbesar dari heap
			heap.Pop(h)
		}
		if h.Len() > 0 {
			ans[i] = (*h)[0].count
		}
	}
	return ans
}

func main() {
	fmt.Println(mostFrequentIDs([]int{1, 2, 3, 2, 2, 1, 3}, []int{-2, 3, -3, 5, 1, -3, 1}))
	fmt.Println(mostFrequentIDs([]int{2, 3, 2, 1}, []int{3, 2, -3, 1}))
}
```
