# 1197 — Minimum Knight Moves

## Deskripsi

**Soal:** [1197. Minimum Knight Moves](https://leetcode.com/problems/minimum-knight-moves/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(  
**Kompleksitas Ruang:** O(

**Algoritma:** BFS (Breadth-First Search / pencarian lebar)

**Fungsi Solusi:** `func minKnightMoves(x int, y int) int`

## Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1197: Minimum Knight Moves
// https://leetcode.com/problems/minimum-knight-moves/
// Difficulty: Medium [Paid]

// BFS from (0,0) to (x,y). Use symmetry: 8 symmetric quadrants.
// Constrain BFS to non-negative coordinates for efficiency.

// Time: O(|x|*|y|) worst case
// Space: O(|x|*|y|)

func minKnightMoves(x int, y int) int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	if x < y {
		x, y = y, x
	}

	// BFS
	dirs := [][]int{{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
		{-2, -1}, {-1, -2}, {1, -2}, {2, -1}}

  // Membuat map untuk pencarian O(1): key → value
	visited := make(map[[2]int]bool)
	queue := [][2]int{{0, 0}}
	visited[[2]int{0, 0}] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur[0] == x && cur[1] == y {
				return steps
			}
			for _, d := range dirs {
				nx, ny := cur[0]+d[0], cur[1]+d[1]
				// Prune: only explore positive but not too far
				if nx >= -2 && ny >= -2 && nx <= x+2 && ny <= y+2 {
					key := [2]int{nx, ny}
					if !visited[key] {
						visited[key] = true
						queue = append(queue, key)
					}
				}
			}
		}
		steps++
	}

	return -1
}

func main() {
	fmt.Printf("%d (expected: 1)\n", minKnightMoves(2, 1))
	fmt.Printf("%d (expected: 2)\n", minKnightMoves(5, 5))
	fmt.Printf("%d (expected: 0)\n", minKnightMoves(0, 0))
}
```
