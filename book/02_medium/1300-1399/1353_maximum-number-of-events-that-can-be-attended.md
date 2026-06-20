# 1353 — Maximum Number Of Events That Can Be Attended

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxEvents(events [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O(n log n) for sorting and heap operations  
**Kompleksitas Ruang:** O(n) for heap

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1353: Maximum Number of Events That Can Be Attended
// https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended/
// Difficulty: Medium

import "fmt"
import "sort"
import "container/heap"

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	// Test case 1
	fmt.Println(maxEvents([][]int{{1, 2}, {2, 3}, {3, 4}})) // 3

	// Test case 2
	fmt.Println(maxEvents([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}})) // 4

	// Test case 3
	fmt.Println(maxEvents([][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {2, 2}})) // 4

	// Test case 4 - single day
	fmt.Println(maxEvents([][]int{{1, 5}, {1, 5}, {1, 5}, {2, 3}, {2, 3}})) // 5
}

// Time: O(n log n) for sorting and heap operations
// Space: O(n) for heap
func maxEvents(events [][]int) int {
	if len(events) == 0 {
		return 0
	}

	// Sort by start day
  // Custom sort dengan comparator
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	h := &minHeap{}
	heap.Init(h)

	i := 0
	day := events[0][0]
	count := 0
	n := len(events)

	for i < n || h.Len() > 0 {
		// Add all events starting today
		for i < n && events[i][0] == day {
  // Masukkan elemen ke priority queue
			heap.Push(h, events[i][1])
			i++
		}

		// Remove expired events (end day < today)
		for h.Len() > 0 && (*h)[0] < day {
  // Ambil elemen terkecil/terbesar dari heap
			heap.Pop(h)
		}

		// Attend one event today (earliest ending)
		if h.Len() > 0 {
  // Ambil elemen terkecil/terbesar dari heap
			heap.Pop(h)
			count++
			day++
		} else if i < n {
			day = events[i][0]
		}
	}

	return count
}
```
