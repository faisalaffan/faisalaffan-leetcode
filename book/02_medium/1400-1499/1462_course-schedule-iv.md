# 1462 — Course Schedule Iv

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^3) for Floyd-Warshall  |  **Ruang:** O(n^2) for reachability matrix


## 💻 Solusi Go

```go
package main

// LeetCode #1462: Course Schedule IV
// https://leetcode.com/problems/course-schedule-iv/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(checkIfPrerequisite(2, [][]int{{1, 0}}, [][]int{{0, 1}, {1, 0}}))
	// [false, true]

	// Test case 2
	fmt.Println(checkIfPrerequisite(2, [][]int{}, [][]int{{1, 0}, {0, 1}}))
	// [false, false]

	// Test case 3
	fmt.Println(checkIfPrerequisite(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, [][]int{{0, 4}, {4, 0}, {1, 3}, {3, 0}}))
	// [true, false, true, false]
}

// Time: O(n^3) for Floyd-Warshall
// Space: O(n^2) for reachability matrix
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	// Build adjacency list
  // Matriks 2D
	adj := make([][]int, numCourses)
	for _, p := range prerequisites {
		adj[p[0]] = append(adj[p[0]], p[1])
	}

	// Floyd-Warshall for reachability
  // Matriks 2D
	reachable := make([][]bool, numCourses)
  // Range loop
	for i := range reachable {
		reachable[i] = make([]bool, numCourses)
	}

	for _, p := range prerequisites {
		reachable[p[0]][p[1]] = true
	}

	for k := 0; k < numCourses; k++ {
		for i := 0; i < numCourses; i++ {
			for j := 0; j < numCourses; j++ {
				if reachable[i][k] && reachable[k][j] {
					reachable[i][j] = true
				}
			}
		}
	}

	result := make([]bool, len(queries))
	for i, q := range queries {
		result[i] = reachable[q[0]][q[1]]
	}

	return result
}
```
