# 0593 — Valid Square

## Deskripsi

**Soal:** [0593. Valid Square](https://leetcode.com/problems/valid-square/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #593: Valid Square
// https://leetcode.com/problems/valid-square/
// Difficulty: Medium
// Time: O(1)
// Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(ValidSquare([]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 1}))
	fmt.Println(ValidSquare([]int{0, 0}, []int{1, 2}, []int{2, 1}, []int{1, 0}))
}

func ValidSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	points := [][]int{p1, p2, p3, p4}
	dists := []int{}

	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			dist := distSq(points[i], points[j])
			dists = append(dists, dist)
		}
	}

	sort.Ints(dists)

	// All 4 sides equal and non-zero, and 2 diagonals equal
	return dists[0] > 0 && dists[0] == dists[1] && dists[1] == dists[2] && dists[2] == dists[3] &&
		dists[4] == dists[5]
}

func distSq(a, b []int) int {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return dx*dx + dy*dy
}
```
