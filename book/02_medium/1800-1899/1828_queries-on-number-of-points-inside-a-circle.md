# 1828 — Queries On Number Of Points Inside A Circle

## Deskripsi

**Soal:** [1828. Queries On Number Of Points Inside A Circle](https://leetcode.com/problems/queries-on-number-of-points-inside-a-circle/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * q), Space: O(q)  
**Kompleksitas Ruang:** O(q)

**Algoritma:** —

**Fungsi Solusi:** `func countPoints(points [][]int, queries [][]int) []int`

## Solusi Go

```go
package main

// LeetCode #1828: Queries on Number of Points Inside a Circle
// https://leetcode.com/problems/queries-on-number-of-points-inside-a-circle/
// Difficulty: Medium
// Time: O(n * q), Space: O(q)

import "fmt"

func countPoints(points [][]int, queries [][]int) []int {
  // Membuat slice untuk menyimpan hasil
	result := make([]int, len(queries))

	for i, q := range queries {
		cx, cy, r := q[0], q[1], q[2]
		r2 := r * r
		count := 0
		for _, p := range points {
			dx := p[0] - cx
			dy := p[1] - cy
			if dx*dx+dy*dy <= r2 {
				count++
			}
		}
		result[i] = count
	}
	return result
}

func main() {
	fmt.Println(countPoints([][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}, [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}))
	// Expected: [3, 2, 2]

	fmt.Println(countPoints([][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}, [][]int{{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3}}))
	// Expected: [2, 3, 2, 3]
}
```
