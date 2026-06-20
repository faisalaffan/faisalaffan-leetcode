# 1091 — Shortest Path In Binary Matrix

## Deskripsi

**Soal:** [1091. Shortest Path In Binary Matrix](https://leetcode.com/problems/shortest-path-in-binary-matrix/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

**Algoritma:** BFS (Breadth-First Search / pencarian lebar)

> **Ide Kunci:** BFS from (0,0) to (n-1,n-1) with 8-directional moves

## Solusi Go

```go
package main

// LeetCode #1091: Shortest Path in Binary Matrix
// https://leetcode.com/problems/shortest-path-in-binary-matrix/
// Difficulty: Medium
//
// Approach: BFS from (0,0) to (n-1,n-1) with 8-directional moves
// Time: O(n^2)
// Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 1}, {1, 0}})) // 2
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}})) // 4
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	if n == 1 {
		return 1
	}

	dirs := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
  // Membuat slice untuk menyimpan hasil
	queue := make([][2]int, 0, n*n)
	queue = append(queue, [2]int{0, 0})
	grid[0][0] = 1
	distance := 1

	for len(queue) > 0 {
		size := len(queue)
		distance++
		for k := 0; k < size; k++ {
			cur := queue[k]
			for _, d := range dirs {
				ni, nj := cur[0]+d[0], cur[1]+d[1]
				if ni >= 0 && ni < n && nj >= 0 && nj < n && grid[ni][nj] == 0 {
					if ni == n-1 && nj == n-1 {
						return distance
					}
					grid[ni][nj] = 1
					queue = append(queue, [2]int{ni, nj})
				}
			}
		}
		queue = queue[size:]
	}

	return -1
}
```
