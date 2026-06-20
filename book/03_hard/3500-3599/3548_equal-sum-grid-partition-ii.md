# 3548 — Equal Sum Grid Partition Ii

## Deskripsi

**Soal:** [3548. Equal Sum Grid Partition Ii](https://leetcode.com/problems/equal-sum-grid-partition-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), Prefix Sum (jumlah kumulatif)

> **Ide Kunci:** Compute prefix sums. Try different partition sizes and check

## Solusi Go

```go
package main

// LeetCode #3548: Equal Sum Grid Partition II
// https://leetcode.com/problems/equal-sum-grid-partition-ii/
// Difficulty: Hard
//
// Given a grid, partition it into sub-rectangles all with equal sum.
// Find the maximum number of partitions possible.
//
// Approach: Compute prefix sums. Try different partition sizes and check
// if grid can be divided into equal-sum sub-rectangles of that size.

import "fmt"

func main() {
	// Example 1
	fmt.Println(equalSumGridPartition([][]int{{1, 2}, {3, 4}}))
	// Example 2: all same
	fmt.Println(equalSumGridPartition([][]int{{1, 1}, {1, 1}}))
	// Example 3: single cell
	fmt.Println(equalSumGridPartition([][]int{{5}}))
	// Edge: 1xn grid
	fmt.Println(equalSumGridPartition([][]int{{1, 2, 3, 4}}))
}

func equalSumGridPartition(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	// Compute prefix sum
  // Membuat slice 2D untuk DP/tabel
	pref := make([][]int, m+1)
  // Iterasi seluruh elemen
	for i := range pref {
		pref[i] = make([]int, n+1)
	}
	total := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pref[i+1][j+1] = pref[i][j+1] + pref[i+1][j] - pref[i][j] + grid[i][j]
			total += grid[i][j]
		}
	}

	sumRegion := func(r1, c1, r2, c2 int) int {
		return pref[r2+1][c2+1] - pref[r1][c2+1] - pref[r2+1][c1] + pref[r1][c1]
	}

	// Try all possible target sums (divisors of total)
	best := 1
	for target := 1; target <= total; target++ {
		if total%target != 0 {
			continue
		}
		// Check if grid can be partitioned into rectangles each summing to target
		count := 0
  // Membuat slice 2D untuk DP/tabel
		visited := make([][]bool, m)
  // Iterasi seluruh elemen
		for i := range visited {
			visited[i] = make([]bool, n)
		}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if visited[i][j] {
					continue
				}
				// Find a rectangle starting at (i,j) that sums to target
				found := false
				for r := i; r < m && !found; r++ {
					for c := j; c < n && !found; c++ {
						if sumRegion(i, j, r, c) == target {
							// Mark this rectangle as visited
							allFree := true
							for x := i; x <= r && allFree; x++ {
								for y := j; y <= c && allFree; y++ {
									if visited[x][y] {
										allFree = false
									}
								}
							}
							if allFree {
								for x := i; x <= r; x++ {
									for y := j; y <= c; y++ {
										visited[x][y] = true
									}
								}
								count++
								found = true
							}
						}
					}
				}
				if !found {
					return best
				}
			}
		}
		if count > best {
			best = count
		}
	}

	return best
}
```
