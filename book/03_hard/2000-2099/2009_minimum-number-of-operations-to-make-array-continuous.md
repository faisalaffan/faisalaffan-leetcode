# 2009 — Minimum Number Of Operations To Make Array Continuous

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minOperations(nums []int) int
```

> **💡 Hint:** Sort + sliding window on unique elements.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2009: Minimum Number of Operations to Make Array Continuous
// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-continuous/
// Difficulty: Hard
// Approach: Sort + sliding window on unique elements.
// A continuous array has max-min < n and distinct elements.
// For each element as left bound, find rightmost element with value < left+n.
// Answer = n - max window size.

import (
	"fmt"
	"sort"
)

func minOperations(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// Sort and deduplicate
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
  // Alokasi slice integer
	uniq := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		uniq = append(uniq, nums[i])
	}

	maxWindow := 0
	right := 0
	for left := 0; left < len(uniq); left++ {
		for right < len(uniq) && uniq[right] < uniq[left]+n {
			right++
		}
		windowSize := right - left
		if windowSize > maxWindow {
			maxWindow = windowSize
		}
	}

	return n - maxWindow
}

func main() {
	// Example: [4,2,5,3] -> 0 (already continuous)
	fmt.Println(minOperations([]int{4, 2, 5, 3}))

	// Additional tests
	fmt.Println(minOperations([]int{1, 2, 3, 5, 6})) // [1,2,3,5,6] -> need to make [2,3,4,5,6] or [1,2,3,4,5], ops = 1
	fmt.Println(minOperations([]int{1, 10, 100, 1000}))
	fmt.Println(minOperations([]int{8, 5, 9, 9, 5, 7, 2, 3}))
}
```
