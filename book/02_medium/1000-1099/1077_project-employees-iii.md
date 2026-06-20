# 1077 — Project Employees Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func projectEmployeesIII(project [][]int, employee [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n log n) where n = len(project)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1077: Project Employees III
// https://leetcode.com/problems/project-employees-iii/
// Difficulty: Medium
//
// Approach: Group employee experience by project, find max per project
// Time: O(n log n) where n = len(project)
// Space: O(n)

import "fmt"

func main() {
	// (project_id, employee_id, experience_years)
	project := [][]int{{1, 1}, {1, 2}, {2, 3}, {2, 4}}
	employee := [][]int{{1, 5}, {2, 3}, {3, 7}, {4, 2}}
	fmt.Println(projectEmployeesIII(project, employee))
}

func projectEmployeesIII(project [][]int, employee [][]int) [][]int {
  // HashMap: O(1) lookup
	expMap := make(map[int]int)
	for _, e := range employee {
		expMap[e[0]] = e[1]
	}

	// For each project, find max experience and which employees have it
	type projInfo struct {
		maxExp int
		empID  int
	}
  // HashMap: O(1) lookup
	projMax := make(map[int]projInfo)

	for _, p := range project {
		projID, empID := p[0], p[1]
		exp := expMap[empID]

		if info, ok := projMax[projID]; !ok || exp > info.maxExp {
			projMax[projID] = projInfo{exp, empID}
		}
	}

  // Matriks 2D
	result := make([][]int, 0)
	for _, p := range project {
		projID, empID := p[0], p[1]
		info := projMax[projID]
		if empID == info.empID {
			result = append(result, []int{projID, empID})
		}
	}

	return result
}
```
