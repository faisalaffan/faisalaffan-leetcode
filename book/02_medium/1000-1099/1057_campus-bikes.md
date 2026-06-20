# 1057 — Campus Bikes

## Deskripsi

**Soal:** [1057. Campus Bikes](https://leetcode.com/problems/campus-bikes/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(W * B) where W = workers, B = bikes  
**Kompleksitas Ruang:** O(W * B)

**Algoritma:** —

> **Ide Kunci:** Bucket sort by Manhattan distance. Assign closest pairs.

## Solusi Go

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
  // Membuat slice 2D untuk DP/tabel
	buckets := make([][][2]int, maxDist+1)

	for i := 0; i < w; i++ {
		for j := 0; j < b; j++ {
			dist := abs(workers[i][0]-bikes[j][0]) + abs(workers[i][1]-bikes[j][1])
			buckets[dist] = append(buckets[dist], [2]int{i, j})
		}
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int, w)
  // Iterasi seluruh elemen
	for i := range result {
		result[i] = -1
	}
  // Membuat slice untuk menyimpan hasil
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
