# 0994 — Rotting Oranges

## Deskripsi

**Soal:** [0994. Rotting Oranges](https://leetcode.com/problems/rotting-oranges/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

**Algoritma:** BFS (Breadth-First Search / pencarian lebar)

> **Ide Kunci:** BFS from all rotten oranges simultaneously

## Solusi Go

```go
package main

// LeetCode #994: Rotting Oranges
// https://leetcode.com/problems/rotting-oranges/
// Difficulty: Medium
//
// Approach: BFS from all rotten oranges simultaneously
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})) // 4
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}})) // -1
	fmt.Println(orangesRotting([][]int{{0, 2}}))                          // 0
}

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
  // Membuat slice untuk menyimpan hasil
	queue := make([][2]int, 0)
	fresh := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	minutes := 0
	for len(queue) > 0 {
		minutes++
		size := len(queue)
		for k := 0; k < size; k++ {
			cur := queue[k]
			for _, d := range dirs {
				ni, nj := cur[0]+d[0], cur[1]+d[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == 1 {
					grid[ni][nj] = 2
					fresh--
					queue = append(queue, [2]int{ni, nj})
				}
			}
		}
		queue = queue[size:]
	}

	if fresh > 0 {
		return -1
	}
	return minutes - 1
}
```
