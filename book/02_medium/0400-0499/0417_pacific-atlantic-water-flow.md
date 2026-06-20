# 0417 — Pacific Atlantic Water Flow

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func pacificAtlantic(heights [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, DFS

**Waktu:** O(m*n)  |  **Ruang:** O(m*n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #417: Pacific Atlantic Water Flow
// https://leetcode.com/problems/pacific-atlantic-water-flow/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 {
		return [][]int{}
	}
	m, n := len(heights), len(heights[0])
  // Matriks 2D
	pacific := make([][]bool, m)
  // Matriks 2D
	atlantic := make([][]bool, m)
  // Range loop
	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	var dfs func(i, j int, visited [][]bool)
	dfs = func(i, j int, visited [][]bool) {
		visited[i][j] = true
		dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && !visited[ni][nj] && heights[ni][nj] >= heights[i][j] {
				dfs(ni, nj, visited)
			}
		}
	}

	// Pacific: top and left edges
	for i := 0; i < m; i++ {
		dfs(i, 0, pacific)
	}
	for j := 0; j < n; j++ {
		dfs(0, j, pacific)
	}

	// Atlantic: bottom and right edges
	for i := 0; i < m; i++ {
		dfs(i, n-1, atlantic)
	}
	for j := 0; j < n; j++ {
		dfs(m-1, j, atlantic)
	}

	result := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}

func main() {
	// Test case 1
	h1 := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	fmt.Println("Test 1:", pacificAtlantic(h1))
	// Expected: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]] (in any order)

	// Test case 2
	h2 := [][]int{{1}}
	fmt.Println("Test 2:", pacificAtlantic(h2))
	// Expected: [[0,0]]
}
```
