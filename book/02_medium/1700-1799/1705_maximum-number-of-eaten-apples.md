# 1705 — Maximum Number Of Eaten Apples

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func eatenApples(apples []int, days []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1705: Maximum Number of Eaten Apples
// https://leetcode.com/problems/maximum-number-of-eaten-apples/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"container/heap"
	"fmt"
)

type Apple struct {
	rottenDay int
	count     int
}

type MinHeap []Apple

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].rottenDay < h[j].rottenDay }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Apple)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func eatenApples(apples []int, days []int) int {
	h := &MinHeap{}
	heap.Init(h)
	eaten := 0
	day := 0

	for day < len(apples) || h.Len() > 0 {
		// New apples grow
		if day < len(apples) && apples[day] > 0 {
  // Masukkan elemen ke priority queue
			heap.Push(h, Apple{rottenDay: day + days[day], count: apples[day]})
		}

		// Remove rotten apples
		for h.Len() > 0 && h.Len() > 0 && (*h)[0].rottenDay <= day {
  // Ambil elemen terkecil/terbesar dari heap
			heap.Pop(h)
		}

		// Eat one apple
		if h.Len() > 0 {
			top := &(*h)[0]
			top.count--
			if top.count == 0 {
  // Ambil elemen terkecil/terbesar dari heap
				heap.Pop(h)
			}
			eaten++
		}

		day++
	}
	return eaten
}

func main() {
	fmt.Println(eatenApples([]int{1, 2, 3, 5, 2}, []int{3, 2, 1, 4, 2})) // Expected: 7
	fmt.Println(eatenApples([]int{3, 0, 0, 0, 0, 2}, []int{3, 0, 0, 0, 0, 2})) // Expected: 5
	fmt.Println(eatenApples([]int{1}, []int{2})) // Expected: 1
}
```
