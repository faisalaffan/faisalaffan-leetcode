# 0295 — Find Median From Data Stream

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() MedianFinder
```

> **💡 Hint:** Two Heaps.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #295: Find Median from Data Stream
// https://leetcode.com/problems/find-median-from-data-stream/
// Difficulty: Hard
//
// Approach: Two Heaps.
//   - Max-heap (lo) stores the smaller half of numbers.
//   - Min-heap (hi) stores the larger half of numbers.
//   - Maintain invariant: lo has at most one more element than hi.
//   - Median: if lo.Len() > hi.Len(), median = lo[0]; else median = (lo[0] + hi[0]) / 2.

import (
	"container/heap"
	"fmt"
)

func main() {
	// Example usage.
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	fmt.Println("Median after [1,2]:", mf.FindMedian()) // 1.5
	mf.AddNum(3)
	fmt.Println("Median after [1,2,3]:", mf.FindMedian()) // 2.0

	// Larger test.
	mf2 := Constructor()
	for _, v := range []int{5, 2, 8, 1, 9, 3, 7} {
		mf2.AddNum(v)
	}
	fmt.Println("Median of [1,2,3,5,7,8,9]:", mf2.FindMedian()) // 5.0
}

// --- Max Heap (for the smaller half) ---

// MaxHeap implements heap.Interface (stores negative values to simulate max-heap).
type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool   { return h[i] > h[j] } // Reverse for max-heap
func (h MaxHeap) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// --- Min Heap (for the larger half) ---

// MinHeap implements heap.Interface.
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool   { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// --- MedianFinder ---

// MedianFinder maintains a data stream median.
type MedianFinder struct {
	lo *MaxHeap // max-heap for smaller half
	hi *MinHeap // min-heap for larger half
}

// Constructor initializes the MedianFinder.
func Constructor() MedianFinder {
	lo := &MaxHeap{}
	hi := &MinHeap{}
	heap.Init(lo)
	heap.Init(hi)
	return MedianFinder{lo: lo, hi: hi}
}

// AddNum adds a number to the data stream.
func (mf *MedianFinder) AddNum(num int) {
	// Step 1: Add to lo (max-heap).
  // Masukkan elemen ke priority queue
	heap.Push(mf.lo, num)

	// Step 2: Move the largest from lo to hi to maintain ordering.
  // Masukkan elemen ke priority queue
	heap.Push(mf.hi, heap.Pop(mf.lo).(int))

	// Step 3: Balance: lo size should be >= hi size.
	if mf.lo.Len() < mf.hi.Len() {
  // Masukkan elemen ke priority queue
		heap.Push(mf.lo, heap.Pop(mf.hi).(int))
	}
}

// FindMedian returns the median of all elements so far.
func (mf *MedianFinder) FindMedian() float64 {
	if mf.lo.Len() > mf.hi.Len() {
		return float64((*mf.lo)[0])
	}
	return float64((*mf.lo)[0]+(*mf.hi)[0]) / 2.0
}

// Stub compatibility.
func FindMedianFromDataStream() any {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	return mf.FindMedian() // 1.5
}

// Ensure heap.Interface is satisfied.
var _ heap.Interface = (*MaxHeap)(nil)
var _ heap.Interface = (*MinHeap)(nil)
```
