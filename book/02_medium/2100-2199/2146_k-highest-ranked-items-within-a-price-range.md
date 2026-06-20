# 2146 — K Highest Ranked Items Within A Price Range

## Deskripsi

**Soal:** [2146. K Highest Ranked Items Within A Price Range](https://leetcode.com/problems/k-highest-ranked-items-within-a-price-range/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m*n log(m*n))  
**Kompleksitas Ruang:** O(m*n)

**Algoritma:** BFS (Breadth-First Search / pencarian lebar)

**Fungsi Solusi:** `func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int`

## Solusi Go

```go
package main

// LeetCode #2146: K Highest Ranked Items Within a Price Range
// https://leetcode.com/problems/k-highest-ranked-items-within-a-price-range/
// Difficulty: Medium
// Time: O(m*n log(m*n)) | Space: O(m*n)

import (
	"fmt"
	"sort"
)

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	low, high := pricing[0], pricing[1]

	// BFS
	type Item struct {
		dist, price, row, col int
	}
  // Membuat slice 2D untuk DP/tabel
	visited := make([][]bool, m)
  // Iterasi seluruh elemen
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := [][2]int{{start[0], start[1]}}
	visited[start[0]][start[1]] = true
	items := []Item{}
	dist := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			r, c := queue[i][0], queue[i][1]
			if grid[r][c] >= low && grid[r][c] <= high {
				items = append(items, Item{dist, grid[r][c], r, c})
			}
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && !visited[nr][nc] && grid[nr][nc] != 0 {
					visited[nr][nc] = true
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
		queue = queue[size:]
		dist++
	}

	// Sort by rank: distance, price, row, col
	sort.Slice(items, func(i, j int) bool {
		if items[i].dist != items[j].dist {
			return items[i].dist < items[j].dist
		}
		if items[i].price != items[j].price {
			return items[i].price < items[j].price
		}
		if items[i].row != items[j].row {
			return items[i].row < items[j].row
		}
		return items[i].col < items[j].col
	})

	// Take first k
  // Membuat slice 2D untuk DP/tabel
	result := make([][]int, 0, k)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(items) && i < k; i++ {
		result = append(result, []int{items[i].row, items[i].col})
	}
	return result
}

func main() {
	// Test case 1
	grid1 := [][]int{{1, 2, 0, 1}, {1, 3, 0, 1}, {0, 2, 5, 1}}
	fmt.Println("Test 1:", highestRankedKItems(grid1, []int{2, 5}, []int{0, 0}, 3))
	// Expected: [[0,1],[1,1],[2,1]]

	// Test case 2
	grid2 := [][]int{{1, 1, 1}, {0, 0, 1}, {2, 3, 4}}
	fmt.Println("Test 2:", highestRankedKItems(grid2, []int{2, 3}, []int{0, 0}, 3))
	// Expected: [[2,1],[2,0]]

	// Test case 3
	grid3 := [][]int{{1, 2}, {3, 4}}
	fmt.Println("Test 3:", highestRankedKItems(grid3, []int{1, 4}, []int{0, 0}, 10))
	// Expected: [[0,0],[0,1],[1,0],[1,1]]
}
```
