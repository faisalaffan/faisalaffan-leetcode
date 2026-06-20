# 2295 — Replace Elements In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func arrayChange(nums []int, operations [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n + m)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2295: Replace Elements in an Array
// https://leetcode.com/problems/replace-elements-in-an-array/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n)

import "fmt"

func arrayChange(nums []int, operations [][]int) []int {
  // HashMap: O(1) lookup
	pos := make(map[int]int)
	for i, v := range nums {
		pos[v] = i
	}

	for _, op := range operations {
		oldVal, newVal := op[0], op[1]
		if idx, ok := pos[oldVal]; ok {
			nums[idx] = newVal
			delete(pos, oldVal)
			pos[newVal] = idx
		}
	}
	return nums
}

func main() {
	// Test case 1
	fmt.Println(arrayChange([]int{1, 2, 4, 6}, [][]int{{1, 3}, {4, 7}, {6, 1}}))
	// Expected: [3,2,7,1]

	// Test case 2
	fmt.Println(arrayChange([]int{1, 2}, [][]int{{1, 3}, {2, 1}, {3, 2}}))
	// Expected: [2,1]
}
```
