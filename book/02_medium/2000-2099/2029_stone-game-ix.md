# 2029 — Stone Game Ix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func stoneGameIX(stones []int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2029: Stone Game IX
// https://leetcode.com/problems/stone-game-ix/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func stoneGameIX(stones []int) bool {
	count := [3]int{}
	for _, v := range stones {
		count[v%3]++
	}

	// If no stone with value % 3 == 0 and the other two counts differ by at most 1
	if count[0]%2 == 0 {
		return count[1] > 0 && count[2] > 0
	}
	// If count[0] is odd
	return abs(count[1]-count[2]) > 2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", stoneGameIX([]int{2, 1}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", stoneGameIX([]int{2}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", stoneGameIX([]int{5, 1, 2, 4, 3}))
	// Expected: false
}
```
