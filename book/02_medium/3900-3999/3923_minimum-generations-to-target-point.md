# 3923 — Minimum Generations To Target Point

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func MinimumGenerationsToTargetPoint(points [][]int, target []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(K * N^2)  |  **Ruang:** O(7^3) where N = seen points ≤ 343, K = max generations


## 💻 Solusi Go

```go
package main

// LeetCode #3923: Minimum Generations to Target Point
// https://leetcode.com/problems/minimum-generations-to-target-point/
// Difficulty: Medium
// Time: O(K * N^2) | Space: O(7^3) where N = seen points ≤ 343, K = max generations
// Approach: BFS simulation. Generate new points by pairing all distinct seen points,
// computing floor midpoint. Coordinate space is [0,6]^3 (max 343 points).
// Track seen points in 3D boolean grid. Return generation when target appears.

import "fmt"

func MinimumGenerationsToTargetPoint(points [][]int, target []int) int {
	// 7x7x7 grid for coordinate space [0,6]
	var seen [7][7][7]bool
	var curGen [][3]int

	// Initialize with generation 0
	for _, p := range points {
		x, y, z := p[0], p[1], p[2]
		if !seen[x][y][z] {
			seen[x][y][z] = true
			curGen = append(curGen, [3]int{x, y, z})
		}
	}

	tx, ty, tz := target[0], target[1], target[2]
	if seen[tx][ty][tz] {
		return 0
	}

	// BFS generation by generation
	for gen := 1; gen <= 7; gen++ {
		// Collect all points seen so far
  // Alokasi slice
		allPoints := make([][3]int, 0, 343)
		for x := 0; x <= 6; x++ {
			for y := 0; y <= 6; y++ {
				for z := 0; z <= 6; z++ {
					if seen[x][y][z] {
						allPoints = append(allPoints, [3]int{x, y, z})
					}
				}
			}
		}

		if len(allPoints) <= 1 {
			break // need at least 2 distinct points to generate
		}

		var nextGen [][3]int
  // Linear scan O(n)
		for i := 0; i < len(allPoints); i++ {
			for j := i + 1; j < len(allPoints); j++ {
				// Compute floor midpoint
				nx := (allPoints[i][0] + allPoints[j][0]) / 2
				ny := (allPoints[i][1] + allPoints[j][1]) / 2
				nz := (allPoints[i][2] + allPoints[j][2]) / 2

				if !seen[nx][ny][nz] {
					seen[nx][ny][nz] = true
					nextGen = append(nextGen, [3]int{nx, ny, nz})
				}
			}
		}

		if seen[tx][ty][tz] {
			return gen
		}

		if len(nextGen) == 0 {
			break
		}

		curGen = nextGen
	}

	return -1
}

func main() {
	// Example 1
	fmt.Println(MinimumGenerationsToTargetPoint([][]int{{0, 0, 0}, {6, 6, 6}}, []int{3, 3, 3})) // Expected: 1

	// Example 2
	fmt.Println(MinimumGenerationsToTargetPoint([][]int{{0, 0, 0}, {5, 5, 5}}, []int{1, 1, 1})) // Expected: 2

	// Example 3
	fmt.Println(MinimumGenerationsToTargetPoint([][]int{{0, 0, 0}, {2, 2, 2}, {3, 3, 3}}, []int{2, 2, 2})) // Expected: 0

	// Example 4
	fmt.Println(MinimumGenerationsToTargetPoint([][]int{{1, 2, 3}}, []int{5, 5, 5})) // Expected: -1
}
```
