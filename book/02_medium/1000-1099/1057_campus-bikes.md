# 1057 — Campus Bikes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func assignBikes(workers [][]int, bikes [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(W * B) where W = workers, B = bikes  |  **Ruang:** O(W * B)


## 💻 Solusi Go

```go
package main

// LeetCode #1057: Campus Bikes
// https://leetcode.com/problems/campus-bikes/
// Difficulty: Medium
//
// Approach: Bucket sort by Manhattan distance. Assign closest pairs.
// Time: O(W * B) where W = workers, B = bikes
// Space: O(W * B)

import "fmt"

func main() {
	fmt.Println(assignBikes([][]int{{0, 0}, {2, 1}}, [][]int{{1, 2}, {3, 3}})) // [1,0]
	fmt.Println(assignBikes([][]int{{0, 0}, {1, 1}, {2, 0}}, [][]int{{1, 0}, {2, 2}, {2, 1}})) // [0,2,1]
}

func assignBikes(workers [][]int, bikes [][]int) []int {
	w, b := len(workers), len(bikes)
	// Buckets of distances: max distance is 2000 (0-1000, 0-1000)
	maxDist := 2000
  // Matriks 2D
	buckets := make([][][2]int, maxDist+1)

	for i := 0; i < w; i++ {
		for j := 0; j < b; j++ {
			dist := abs(workers[i][0]-bikes[j][0]) + abs(workers[i][1]-bikes[j][1])
			buckets[dist] = append(buckets[dist], [2]int{i, j})
		}
	}

  // Alokasi slice
	result := make([]int, w)
  // Range loop
	for i := range result {
		result[i] = -1
	}
	bikeUsed := make([]bool, b)

	for d := 0; d <= maxDist; d++ {
		for _, pair := range buckets[d] {
			wi, bi := pair[0], pair[1]
			if result[wi] == -1 && !bikeUsed[bi] {
				result[wi] = bi
				bikeUsed[bi] = true
			}
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
