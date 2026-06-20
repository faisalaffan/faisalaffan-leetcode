# 1875 — Group Employees Of The Same Salary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func GroupEmployees(employees [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1875: Group Employees of the Same Salary
// https://leetcode.com/problems/group-employees-of-the-same-salary/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	// employee: [id, salary]
	employees := [][]int{{1, 50000}, {2, 60000}, {3, 50000}, {4, 70000}, {5, 60000}}
	fmt.Println(GroupEmployees(employees))
}

// Time: O(n log n), Space: O(n)
func GroupEmployees(employees [][]int) [][]int {
  // HashMap: O(1) lookup
	salaryMap := make(map[int][]int)
	for _, emp := range employees {
		id, salary := emp[0], emp[1]
		salaryMap[salary] = append(salaryMap[salary], id)
	}

  // Matriks 2D
	result := make([][]int, 0)
	for _, ids := range salaryMap {
		if len(ids) >= 2 {
  // Sort O(n log n)
			sort.Ints(ids)
			result = append(result, ids)
		}
	}

	// Sort by first employee ID for deterministic output
  // Custom sort
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}
```
