# 0356 — Line Reflection

## Deskripsi

**Soal:** [0356. Line Reflection](https://leetcode.com/problems/line-reflection/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func isReflected(points [][]int) bool`

## Solusi Go

```go
package main

// LeetCode #356: Line Reflection
// https://leetcode.com/problems/line-reflection/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func isReflected(points [][]int) bool {
	if len(points) == 0 {
		return true
	}

  // Membuat map untuk pencarian O(1): key → value
	set := make(map[[2]int]bool)
	minX, maxX := points[0][0], points[0][0]

	for _, p := range points {
		set[[2]int{p[0], p[1]}] = true
		if p[0] < minX {
			minX = p[0]
		}
		if p[0] > maxX {
			maxX = p[0]
		}
	}

	sum := minX + maxX // 2 * mid
	for _, p := range points {
		reflect := [2]int{sum - p[0], p[1]}
		if !set[reflect] {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1: Points symmetric across vertical line x=1
	fmt.Println("Test 1:", isReflected([][]int{{1, 1}, {-1, 1}}))
	// Expected: true

	// Test case 2: Not symmetric
	fmt.Println("Test 2:", isReflected([][]int{{1, 1}, {-1, -1}}))
	// Expected: false

	// Test case 3: Points on same line
	fmt.Println("Test 3:", isReflected([][]int{{0, 0}, {1, 0}, {2, 0}}))
	// Expected: true (midline at x=1, 0 reflects to 2)
}
```
