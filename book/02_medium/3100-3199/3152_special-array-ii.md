# 3152 — Special Array Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func isArraySpecial(nums []int, queries [][]int) []bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n + q)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3152: Special Array II
// https://leetcode.com/problems/special-array-ii/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
  // Alokasi slice
	prefix := make([]int, n)
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1]
		if nums[i]%2 == nums[i-1]%2 {
			prefix[i]++
		}
	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		from, to := q[0], q[1]
		ans[i] = prefix[from] == prefix[to]
	}
	return ans
}

func main() {
	fmt.Println(isArraySpecial([]int{3, 4, 1, 2, 6}, [][]int{{0, 4}}))        // Expected: [false]
	fmt.Println(isArraySpecial([]int{4, 3, 1, 6}, [][]int{{0, 2}, {2, 3}}))   // Expected: [false, true]
}
```
