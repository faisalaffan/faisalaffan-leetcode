# 0207 — Course Schedule

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func canFinish(numCourses int, prerequisites [][]int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS

**Waktu:** O(V+E), Space: O(V+E)  |  **Ruang:** O(V+E)

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #207: Course Schedule
// https://leetcode.com/problems/course-schedule/
// Difficulty: Medium
// Time: O(V+E), Space: O(V+E)

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
  // Matriks 2D
	graph := make([][]int, numCourses)
  // Alokasi slice
	inDegree := make([]int, numCourses)

	for _, pre := range prerequisites {
		course, prereq := pre[0], pre[1]
		graph[prereq] = append(graph[prereq], course)
		inDegree[course]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	count := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		count++

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return count == numCourses
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(canFinish(5, [][]int{{1, 0}, {2, 1}, {3, 2}}))
}
```
