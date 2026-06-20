# 2386 — Find The K Sum Of An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func kSum(nums []int, k int) int64
```

> **💡 Hint:** // 1. Compute the maximum sum = sum of all positive numbers.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2386: Find the K-Sum of an Array
// https://leetcode.com/problems/find-the-k-sum-of-an-array/
// Difficulty: Hard
//
// Given an array nums and integer k, return the k-th largest sum of any
// subsequence of nums (non-empty allowed).
//
// Approach:
// 1. Compute the maximum sum = sum of all positive numbers.
// 2. Take absolute values of all numbers and sort them.
// 3. The k-th largest subsequence sum is obtained by subtracting some
//    combination of absolute values from the max sum.
// 4. Use a min-heap to generate sums in decreasing order.
//    Each state is (currentSum, index). From each state, we can either
//    subtract the next absolute value, or replace the current subtraction
//    with the next one.

import (
	"container/heap"
	"fmt"
	"sort"
)

type Item struct {
	sum int64
	idx int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kSum(nums []int, k int) int64 {
	n := len(nums)

	// Sum of all positive numbers
	var maxSum int64
  // Alokasi slice integer
	absVals := make([]int64, 0, n)
	for _, v := range nums {
		if v > 0 {
			maxSum += int64(v)
		}
		av := int64(v)
		if av < 0 {
			av = -av
		}
		absVals = append(absVals, av)
	}

	// Sort by absolute value
  // Custom sort dengan comparator
	sort.Slice(absVals, func(i, j int) bool {
		return absVals[i] < absVals[j]
	})

	// The k-th largest sum
	// ans[0] = maxSum (largest subsequence sum = sum of all positives)
	// ans[1..k-1] generated from heap
	h := &MinHeap{}
  // Masukkan elemen ke priority queue
	heap.Push(h, Item{sum: maxSum - absVals[0], idx: 0})

	var ans int64 = maxSum

	for i := 1; i < k; i++ {
  // Ambil elemen terkecil/terbesar dari heap
		it := heap.Pop(h).(Item)
		ans = it.sum

		if it.idx+1 < n {
			// Extend: subtract the next value
			nextSum := it.sum - absVals[it.idx+1]
  // Masukkan elemen ke priority queue
			heap.Push(h, Item{sum: nextSum, idx: it.idx + 1})

			// Replace: undo current subtraction and subtract the next
			replaceSum := it.sum + absVals[it.idx] - absVals[it.idx+1]
  // Masukkan elemen ke priority queue
			heap.Push(h, Item{sum: replaceSum, idx: it.idx + 1})
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(kSum([]int{2, 4, -2}, 5))
	// Example 2
	fmt.Println(kSum([]int{3, -1, -2}, 4))
	// Single positive
	fmt.Println(kSum([]int{5}, 1))
	// All negatives
	fmt.Println(kSum([]int{-1, -2, -3}, 3))
	// Mixed
	fmt.Println(kSum([]int{10, -5, 3, -2}, 6))
}
```
