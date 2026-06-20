# 3782 — Last Remaining Integer After Alternating Deletion Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func lastRemaining(n int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3782: Last Remaining Integer After Alternating
// Deletion Operations
// https://leetcode.com/problems/last-remaining-integer-after-
// alternating-deletion-operations/
// Difficulty: Hard
//
// Start with list [1, 2, ..., n]. Delete elements alternately
// from left and right ends until one remains. Return survivor.
//
// Approach: Each step removes two elements (one from each end)
// alternately. The survivor position follows: f(n) = n/2 + 1.

import "fmt"

func main() {
	// Example 1
	fmt.Println(lastRemaining(5))
	// Example 2
	fmt.Println(lastRemaining(7))
	// Edge: single element
	fmt.Println(lastRemaining(1))
	// Edge: two elements
	fmt.Println(lastRemaining(2))
	// Edge: large
	fmt.Println(lastRemaining(100))
}

func lastRemaining(n int) int64 {
	if n <= 0 {
		return 0
	}
	// Each full cycle of left-right deletes 2 elements.
	// The survivor is at position floor(n/2) + 1 (1-indexed).
	return int64(n/2 + 1)
}
```
