# 2126 — Destroying Asteroids

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func asteroidsDestroyed(mass int, asteroids []int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2126: Destroying Asteroids
// https://leetcode.com/problems/destroying-asteroids/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func asteroidsDestroyed(mass int, asteroids []int) bool {
  // Sort O(n log n)
	sort.Ints(asteroids)
	current := int64(mass)

	for _, a := range asteroids {
		if current < int64(a) {
			return false
		}
		current += int64(a)
	}

	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", asteroidsDestroyed(10, []int{3, 9, 19, 5, 21}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", asteroidsDestroyed(5, []int{4, 9, 23, 4}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", asteroidsDestroyed(1, []int{1, 1, 1, 1}))
	// Expected: true
}
```
