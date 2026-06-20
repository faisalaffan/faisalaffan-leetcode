# 3355 — Zero Array Transformation I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func isZeroArray(nums []int, queries [][]int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n + q) Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #3355: Zero Array Transformation I
// https://leetcode.com/problems/zero-array-transformation-i/
// Difficulty: Medium
// Time: O(n + q) Space: O(n)

import "fmt"

func main() {
	fmt.Println(isZeroArray([]int{1, 0, 1}, [][]int{{0, 2}, {0, 2}})) // true
	fmt.Println(isZeroArray([]int{2, 0, 2}, [][]int{{0, 2}, {0, 2}, {1, 1}})) // true
}

func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
  // Alokasi slice
	diff := make([]int, n+1)

	for _, q := range queries {
		l, r := q[0], q[1]
		diff[l]++
		diff[r+1]--
	}

	cur := 0
	for i := 0; i < n; i++ {
		cur += diff[i]
		if cur < nums[i] {
			return false
		}
	}
	return true
}
```
