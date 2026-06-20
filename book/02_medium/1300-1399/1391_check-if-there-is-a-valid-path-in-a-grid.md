# 1391 — Check If There Is A Valid Path In A Grid

## Deskripsi

**Soal:** [1391. Check If There Is A Valid Path In A Grid](https://leetcode.com/problems/check-if-there-is-a-valid-path-in-a-grid/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m*n) where m,n = grid dimensions  
**Kompleksitas Ruang:** O(m*n) for visited array

**Algoritma:** Dynamic Programming (DP), LIS (Longest Increasing Subsequence)

## Solusi Go

```go
package main

// LeetCode #1391: Check if There is a Valid Path in a Grid
// https://leetcode.com/problems/check-if-there-is-a-valid-path-in-a-grid/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(hasValidPath([][]int{{2, 4, 3}, {6, 5, 2}})) // true

	// Test case 2
	fmt.Println(hasValidPath([][]int{{1, 2, 3}, {4, 5, 6}})) // false

	// Test case 3
	fmt.Println(hasValidPath([][]int{{4, 1}, {6, 1}})) // true

	// Test case 4
	fmt.Println(hasValidPath([][]int{{2}, {2}, {2}, {2}, {2}})) // false
}

// Directions: 0=right, 1=down, 2=left, 3=up
// Each street type defines which directions it connects to
// Street 1: left(2)-right(0)
// Street 2: up(3)-down(1)
// Street 3: left(2)-down(1)
// Street 4: right(0)-down(1)
// Street 5: left(2)-up(3)
// Street 6: right(0)-up(3)

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // right, down, left, up

// street[type][direction] = list of compatible outgoing directions
var street = [7][][]int{
	{},
	{{}, {0, 2}, {}, {0, 2}, {}},    // type 1
	{{}, {1, 3}, {1, 3}, {}, {}},    // type 2
	{{}, {}, {1, 2}, {2, 1}, {}},    // type 3: left(2)->down(1), down(1)->left(2)
	{{}, {}, {0, 1}, {1, 0}, {}},    // type 4: right(0)->down(1), down(1)->right(0)
	{{}, {}, {2, 3}, {3, 2}, {}},    // type 5: left(2)->up(3), up(3)->left(2)
	{{}, {}, {0, 3}, {3, 0}, {}},    // type 6: right(0)->up(3), up(3)->right(0)
}

// Time: O(m*n) where m,n = grid dimensions
// Space: O(m*n) for visited array
func hasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
  // Membuat slice 2D untuk DP/tabel
	visited := make([][]bool, m)
  // Iterasi seluruh elemen
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var dfs func(int, int) bool
	dfs = func(r, c int) bool {
		if r == m-1 && c == n-1 {
			return true
		}
		visited[r][c] = true

		st := grid[r][c]
		// Try each direction the current street allows
		for _, dir := range []int{0, 1, 2, 3} {
			nr, nc := r+dirs[dir][0], c+dirs[dir][1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n || visited[nr][nc] {
				continue
			}
			// Check if current street allows this direction
			if !connects(st, dir) {
				continue
			}
			// Check if next street can connect back
			opp := (dir + 2) % 4
			if !connects(grid[nr][nc], opp) {
				continue
			}
			if dfs(nr, nc) {
				return true
			}
		}
		return false
	}

	return dfs(0, 0)
}

func connects(streetType, dir int) bool {
	switch streetType {
	case 1:
		return dir == 0 || dir == 2
	case 2:
		return dir == 1 || dir == 3
	case 3:
		return dir == 1 || dir == 2
	case 4:
		return dir == 0 || dir == 1
	case 5:
		return dir == 2 || dir == 3
	case 6:
		return dir == 0 || dir == 3
	}
	return false
}
```
