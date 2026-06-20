# 0570 — Managers With At Least 5 Direct Reports

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func FindManagers(employees [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #570: Managers with at Least 5 Direct Reports
// https://leetcode.com/problems/managers-with-at-least-5-direct-reports/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Employees: {id, name, department, managerId}
	// managerId == -1 means top-level manager
	employees := [][]int{
		{101, 1, 0, -1},  // John, manager
		{102, 2, 0, 101}, // Dan
		{103, 3, 0, 101}, // James
		{104, 4, 0, 101}, // Amy
		{105, 5, 0, 101}, // Ben
		{106, 6, 0, 101}, // Sam
		{107, 7, 1, -1},  // Ron, another manager
		{108, 8, 1, 107}, // Tom
	}
	fmt.Println(FindManagers(employees))
}

func FindManagers(employees [][]int) []int {
  // HashMap: O(1) lookup
	reportCount := make(map[int]int)
  // HashMap: O(1) lookup
	managerNames := make(map[int]int) // managerId -> manager name (for simplicity, just id)

	for _, emp := range employees {
		id, name, _, managerId := emp[0], emp[1], emp[2], emp[3]
		managerNames[id] = name
		if managerId != -1 {
			reportCount[managerId]++
		}
	}

	result := []int{}
	for mid, count := range reportCount {
		if count >= 5 {
			result = append(result, managerNames[mid])
		}
	}

	return result
}
```
