# 2249 — Count Lattice Points Inside A Circle

## Deskripsi

**Soal:** [2249. Count Lattice Points Inside A Circle](https://leetcode.com/problems/count-lattice-points-inside-a-circle/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * r^2)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func countLatticePoints(circles [][]int) int`

## Solusi Go

```go
package main

// LeetCode #2249: Count Lattice Points Inside a Circle
// https://leetcode.com/problems/count-lattice-points-inside-a-circle/
// Difficulty: Medium
// Time: O(n * r^2) | Space: O(n)

import "fmt"

func countLatticePoints(circles [][]int) int {
  // Membuat map untuk pencarian O(1): key → value
	points := make(map[[2]int]bool)

	for _, c := range circles {
		x, y, r := c[0], c[1], c[2]
		rr := r * r
		for dx := -r; dx <= r; dx++ {
			for dy := -r; dy <= r; dy++ {
				if dx*dx+dy*dy <= rr {
					points[[2]int{x + dx, y + dy}] = true
				}
			}
		}
	}
	return len(points)
}

func main() {
	// Test case 1
	fmt.Println(countLatticePoints([][]int{{2, 2, 1}}))
	// Expected: 5

	// Test case 2
	fmt.Println(countLatticePoints([][]int{{2, 2, 2}, {3, 4, 1}}))
	// Expected: 16
}
```
