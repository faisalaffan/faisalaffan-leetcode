# 0768 — Max Chunks To Make Sorted Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxChunksToSorted(arr []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #768: Max Chunks To Make Sorted II
// https://leetcode.com/problems/max-chunks-to-make-sorted-ii/
// Difficulty: Hard
//
// Given an array arr, split it into chunks (contiguous segments).
// Sort each chunk individually, then concatenate. The result should
// equal the sorted array. Find the maximum number of chunks.
//
// Approach: Prefix Max / Suffix Min
// A chunk can end at index i if max(arr[0..i]) <= min(arr[i+1..n-1]).
// Because after sorting each chunk, all elements in the left chunk
// will be <= all elements in the right chunk.

import "fmt"

func main() {
	fmt.Println(maxChunksToSorted([]int{5, 4, 3, 2, 1}))          // 1
	fmt.Println(maxChunksToSorted([]int{2, 1, 3, 4, 4}))          // 4
	fmt.Println(maxChunksToSorted([]int{4, 3, 2, 1, 0}))          // 1
	fmt.Println(maxChunksToSorted([]int{1, 0, 1, 3, 2}))          // 3
	fmt.Println(maxChunksToSorted([]int{0, 0, 1, 1, 1}))          // 5
	fmt.Println(maxChunksToSorted([]int{1, 1, 0, 0, 1}))          // 2
}

func maxChunksToSorted(arr []int) int {
	n := len(arr)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	// prefixMax[i] = max(arr[0..i])
  // Alokasi slice
	prefixMax := make([]int, n)
	prefixMax[0] = arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > prefixMax[i-1] {
			prefixMax[i] = arr[i]
		} else {
			prefixMax[i] = prefixMax[i-1]
		}
	}

	// suffixMin[i] = min(arr[i..n-1])
  // Alokasi slice
	suffixMin := make([]int, n)
	suffixMin[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		if arr[i] < suffixMin[i+1] {
			suffixMin[i] = arr[i]
		} else {
			suffixMin[i] = suffixMin[i+1]
		}
	}

	chunks := 1
	for i := 0; i < n-1; i++ {
		if prefixMax[i] <= suffixMin[i+1] {
			chunks++
		}
	}
	return chunks
}
```
