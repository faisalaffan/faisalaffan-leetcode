# 1792 — Maximum Average Pass Ratio

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxAverageRatio(classes [][]int, extraStudents int) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O((n+k) log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1792: Maximum Average Pass Ratio
// https://leetcode.com/problems/maximum-average-pass-ratio/
// Difficulty: Medium
// Time: O((n+k) log n), Space: O(n)

import (
	"container/heap"
	"fmt"
)

type Class struct {
	pass, total int
	gain        float64
}

type MaxHeap []Class

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].gain > h[j].gain }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Class)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := &MaxHeap{}
	heap.Init(h)

	for _, c := range classes {
		pass, total := c[0], c[1]
		gain := float64(pass+1)/float64(total+1) - float64(pass)/float64(total)
  // Masukkan elemen ke priority queue
		heap.Push(h, Class{pass, total, gain})
	}

	for i := 0; i < extraStudents; i++ {
  // Ambil elemen terkecil/terbesar dari heap
		c := heap.Pop(h).(Class)
		c.pass++
		c.total++
		c.gain = float64(c.pass+1)/float64(c.total+1) - float64(c.pass)/float64(c.total)
  // Masukkan elemen ke priority queue
		heap.Push(h, c)
	}

	sum := 0.0
	for h.Len() > 0 {
  // Ambil elemen terkecil/terbesar dari heap
		c := heap.Pop(h).(Class)
		sum += float64(c.pass) / float64(c.total)
	}
	return sum / float64(len(classes))
}

func main() {
	fmt.Printf("%.5f\n", maxAverageRatio([][]int{{1, 2}, {3, 5}, {2, 2}}, 2)) // Expected: 0.78333
	fmt.Printf("%.5f\n", maxAverageRatio([][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}}, 4)) // Expected: 0.53485
	fmt.Printf("%.5f\n", maxAverageRatio([][]int{{1, 2}}, 1)) // Expected: 0.66667
}
```
