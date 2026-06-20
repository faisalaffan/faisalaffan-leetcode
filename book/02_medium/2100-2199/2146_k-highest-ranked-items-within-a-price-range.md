# 2146 — K Highest Ranked Items Within A Price Range

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(m*n log(m*n))  
**Kompleksitas Ruang:** O(m*n)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

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
  // Membuat matriks/slice 2D untuk DP
	visited := make([][]bool, m)
  // Range loop: iterasi dengan indeks + nilai
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
  // Custom sort dengan comparator
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
  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, 0, k)
  // Loop linear O(n): iterasi setiap elemen
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
