# 1036 — Escape A Large Maze

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func isEscapePossible(blocked [][]int, source []int, target []int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, BFS

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1036: Escape a Large Maze
// https://leetcode.com/problems/escape-a-large-maze/
// Difficulty: Hard
//
// Approach: BFS limited by blocked cells.
//   The grid is 1M x 1M, too large for full BFS. But there are at most 200 blocked cells.
//   If source and target aren't fully enclosed by blocked cells, BFS from source
//   visiting at most len(blocked)*(len(blocked)-1)/2 cells will either reach target
//   or escape the bounded area. We do the same from target to source.

import "fmt"

func main() {
	blocked := [][]int{{0, 1}, {1, 0}}
	source := []int{0, 0}
	target := []int{0, 2}
	fmt.Println(isEscapePossible(blocked, source, target)) // false

	blocked2 := [][]int{}
	fmt.Println(isEscapePossible(blocked2, []int{0, 0}, []int{999999, 999999})) // true
}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	n := len(blocked)
  // Edge case: input kosong
	if n == 0 {
		return true
	}
	// Maximum cells we need to explore before determining escape
	limit := n * (n + 1) / 2

  // HashMap: O(1) lookup
	blockedSet := make(map[[2]int]bool)
	for _, b := range blocked {
		blockedSet[[2]int{b[0], b[1]}] = true
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	bfs := func(start, target []int) bool {
  // HashMap: O(1) lookup
		visited := make(map[[2]int]bool)
		queue := [][2]int{{start[0], start[1]}}
		visited[[2]int{start[0], start[1]}] = true

		for len(queue) > 0 && len(visited) <= limit {
			cur := queue[0]
			queue = queue[1:]

			if cur[0] == target[0] && cur[1] == target[1] {
				return true
			}

			for _, d := range dirs {
				nr, nc := cur[0]+d[0], cur[1]+d[1]
				key := [2]int{nr, nc}
				if nr < 0 || nr >= 1000000 || nc < 0 || nc >= 1000000 {
					continue
				}
				if visited[key] || blockedSet[key] {
					continue
				}
				visited[key] = true
				queue = append(queue, key)
			}
		}
		// If we visited more than limit cells, we have escaped the enclosed area
		return len(visited) > limit
	}

	return bfs(source, target) && bfs(target, source)
}
```
