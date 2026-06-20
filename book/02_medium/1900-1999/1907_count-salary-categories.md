# 1907 — Count Salary Categories

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func CountSalaryCategories(accounts [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1907: Count Salary Categories
// https://leetcode.com/problems/count-salary-categories/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// accounts: [account_id, income]
	accounts := [][]int{{1, 20000}, {2, 50000}, {3, 100000}, {4, 80000}, {5, 30000}}
	fmt.Println(CountSalaryCategories(accounts))
}

// Time: O(n), Space: O(1)
func CountSalaryCategories(accounts [][]int) []int {
	low := 0
	mid := 0
	high := 0

	for _, a := range accounts {
		income := a[1]
		if income < 20000 {
			low++
		} else if income <= 50000 {
			mid++
		} else {
			high++
		}
	}
	return []int{low, mid, high}
}
```
