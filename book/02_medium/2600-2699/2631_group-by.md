# 2631 — Group By

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2631: Group By
// https://leetcode.com/problems/group-by/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func groupBy[T any, K comparable](items []T, keyFn func(T) K) map[K][]T {
  // HashMap: O(1) lookup
	result := make(map[K][]T)
	for _, item := range items {
		key := keyFn(item)
		result[key] = append(result[key], item)
	}
	return result
}

func main() {
	// Test case 1: group integers by even/odd
	nums := []int{1, 2, 3, 4, 5, 6}
	grouped := groupBy(nums, func(n int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	})
	fmt.Println("Test 1:", grouped)
	// Expected: map[even:[2 4 6] odd:[1 3 5]]

	// Test case 2: group strings by length
	words := []string{"one", "two", "three", "four"}
	byLen := groupBy(words, func(s string) int {
		return len(s)
	})
	fmt.Println("Test 2:", byLen)
	// Expected: map[3:[one two] 5:[three four]]

	// Test case 3
	single := []int{1}
	result := groupBy(single, func(n int) string { return "a" })
	fmt.Println("Test 3:", result)
}
```
