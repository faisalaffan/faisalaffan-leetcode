# 3873 — Maximum Points Activated With One Addition

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func maxActivated(points [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3873: Maximum Points Activated with One Addition
// https://leetcode.com/problems/maximum-points-activated-with-one-addition/
// Difficulty: Hard
//
// Given integer points on a grid, add one new point at integer
// coordinates. A point is "activated" if it has another point
// (original or added) within Manhattan distance 1 (sharing a side).
// Maximize activated points after addition.
//
// Approach: Count already-activated points (have a neighbor). For
// each existing point, try adding adjacent to it. Count new
// activations from this candidate.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maxActivated([][]int{{0, 0}, {1, 0}, {0, 1}}))
	// Example 2
	fmt.Println(maxActivated([][]int{{0, 0}, {2, 0}}))
	// Edge: single point
	fmt.Println(maxActivated([][]int{{5, 5}}))
	// Edge: two adjacent
	fmt.Println(maxActivated([][]int{{0, 0}, {0, 1}}))
}

func maxActivated(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	if len(points) == 1 {
		return 0
	}

  // HashMap: O(1) lookup
	pointSet := make(map[[2]int]bool)
	for _, p := range points {
		pointSet[[2]int{p[0], p[1]}] = true
	}

	// Count already activated points (have a neighbor)
  // Alokasi slice
	isActivated := make([]bool, len(points))
	activatedCnt := 0
	for i, p := range points {
		x, y := p[0], p[1]
		if pointSet[[2]int{x - 1, y}] || pointSet[[2]int{x + 1, y}] ||
			pointSet[[2]int{x, y - 1}] || pointSet[[2]int{x, y + 1}] {
			isActivated[i] = true
			activatedCnt++
		}
	}

	if activatedCnt == len(points) {
		return activatedCnt
	}

	// For each candidate adjacent to each point, compute score
  // HashMap: O(1) lookup
	candidateScore := make(map[[2]int]int)
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for _, p := range points {
		x, y := p[0], p[1]
		for _, d := range dirs {
			cx, cy := x+d[0], y+d[1]
			if pointSet[[2]int{cx, cy}] {
				continue
			}
			score := countNewActivated(cx, cy, pointSet, isActivated)
			if score > candidateScore[[2]int{cx, cy}] {
				candidateScore[[2]int{cx, cy}] = score
			}
		}
	}

	best := activatedCnt
	for _, cnt := range candidateScore {
		if cnt > best {
			best = cnt
		}
	}
	return best
}

func countNewActivated(cx, cy int, pointSet map[[2]int]bool, isActivated []bool) int {
	// Count newly activated: the candidate itself + neighbors that
	// get activated by being adjacent to the candidate
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	hasNeighbor := false
	newCnt := 0

	for _, d := range dirs {
		nx, ny := cx+d[0], cy+d[1]
		if pointSet[[2]int{nx, ny}] {
			hasNeighbor = true
			// Check if this neighbor becomes activated
			// (it already is, or gets activated by candidate)
		}
	}

	if hasNeighbor {
		newCnt++ // candidate itself activated
	}

	// Count neighbor points that get newly activated
	for _, d := range dirs {
		nx, ny := cx+d[0], cy+d[1]
		if pointSet[[2]int{nx, ny}] {
			// Check if this point was already activated
			alreadyActivated := false
			px, py := nx, ny
			if pointSet[[2]int{px - 1, py}] && (px-1 != cx || py != cy) {
				alreadyActivated = true
			}
			if pointSet[[2]int{px + 1, py}] && (px+1 != cx || py != cy) {
				alreadyActivated = true
			}
			if pointSet[[2]int{px, py - 1}] && (px != cx || py-1 != cy) {
				alreadyActivated = true
			}
			if pointSet[[2]int{px, py + 1}] && (px != cx || py+1 != cy) {
				alreadyActivated = true
			}
			if !alreadyActivated {
				newCnt++
			}
		}
	}

	return newCnt
}
```
