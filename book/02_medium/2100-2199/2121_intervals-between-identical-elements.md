# 2121 — Intervals Between Identical Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func getDistances(arr []int) []int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer, Prefix Sum

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2121: Intervals Between Identical Elements
// https://leetcode.com/problems/intervals-between-identical-elements/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func getDistances(arr []int) []int64 {
	n := len(arr)
	// Group indices by value
  // HashMap: O(1) lookup
	groups := make(map[int][]int)
	for i, v := range arr {
		groups[v] = append(groups[v], i)
	}

  // Alokasi slice
	result := make([]int64, n)
	for _, indices := range groups {
		m := len(indices)
		if m <= 1 {
			continue
		}
		// Prefix sum of distances
  // Alokasi slice
		prefix := make([]int64, m+1)
		for i := 0; i < m; i++ {
			prefix[i+1] = prefix[i] + int64(indices[i])
		}
		for i, pos := range indices {
			// Sum of distances to all other same-value elements
			// Left side: pos * i - prefix[i]
			// Right side: (prefix[m] - prefix[i+1]) - pos * (m-1-i)
			left := int64(pos)*int64(i) - prefix[i]
			right := (prefix[m] - prefix[i+1]) - int64(pos)*int64(m-1-i)
			result[pos] = left + right
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", getDistances([]int{2, 1, 3, 1, 2, 3, 3}))
	// Expected: [4, 2, 7, 2, 4, 4, 5]

	// Test case 2
	fmt.Println("Test 2:", getDistances([]int{10, 5, 10, 10}))
	// Expected: [5, 0, 3, 4]

	// Test case 3
	fmt.Println("Test 3:", getDistances([]int{1, 2, 3}))
	// Expected: [0, 0, 0]
}
```
