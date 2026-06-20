# 3061 — Calculate Trapping Rain Water

## Deskripsi

**Soal:** [3061. Calculate Trapping Rain Water](https://leetcode.com/problems/calculate-trapping-rain-water/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func calculateTrappingRainWater(heights []Height) int`

> **Ide Kunci:** Classic two-pass prefix/suffix max algorithm.

## Solusi Go

```go
package main

// LeetCode #3061: Calculate Trapping Rain Water (SQL simulation)
// https://leetcode.com/problems/calculate-trapping-rain-water/
// Difficulty: Hard [Paid]
//
// Approach: Classic two-pass prefix/suffix max algorithm.
// For each position, water trapped = min(leftMax, rightMax) - height.

import "fmt"

type Height struct {
	ID     int
	Height int
}

func calculateTrappingRainWater(heights []Height) int {
	n := len(heights)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}
  // Membuat slice untuk menyimpan hasil
	h := make([]int, n)
	for i, ht := range heights {
		h[i] = ht.Height
	}
  // Membuat slice untuk menyimpan hasil
	leftMax := make([]int, n)
  // Membuat slice untuk menyimpan hasil
	rightMax := make([]int, n)
	leftMax[0] = h[0]
	for i := 1; i < n; i++ {
		if h[i] > leftMax[i-1] {
			leftMax[i] = h[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}
	rightMax[n-1] = h[n-1]
	for i := n - 2; i >= 0; i-- {
		if h[i] > rightMax[i+1] {
			rightMax[i] = h[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}
	total := 0
	for i := 0; i < n; i++ {
		minBound := leftMax[i]
		if rightMax[i] < minBound {
			minBound = rightMax[i]
		}
		total += minBound - h[i]
	}
	return total
}

func main() {
	// Example 1: classic LeetCode test
	heights1 := []Height{
		{1, 0}, {2, 1}, {3, 0}, {4, 2},
		{5, 1}, {6, 0}, {7, 1}, {8, 3},
		{9, 2}, {10, 1}, {11, 2}, {12, 1},
	}
	fmt.Println("Example 1 (classic):", calculateTrappingRainWater(heights1))
	// Expected: 6

	// Example 2: increasing
	heights2 := []Height{
		{1, 1}, {2, 2}, {3, 3}, {4, 4},
	}
	fmt.Println("Example 2 (increasing):", calculateTrappingRainWater(heights2))
	// Expected: 0

	// Example 3: valley
	heights3 := []Height{
		{1, 4}, {2, 0}, {3, 0}, {4, 4},
	}
	fmt.Println("Example 3 (valley):", calculateTrappingRainWater(heights3))
	// Expected: 8 (4 per middle column)

	// Example 4: single element
	heights4 := []Height{
		{1, 5},
	}
	fmt.Println("Example 4 (single):", calculateTrappingRainWater(heights4))
	// Expected: 0

	// Example 5: two elements
	heights5 := []Height{
		{1, 3}, {2, 5},
	}
	fmt.Println("Example 5 (two):", calculateTrappingRainWater(heights5))
	// Expected: 0

	// Example 6: empty
	heights6 := []Height{}
	fmt.Println("Example 6 (empty):", calculateTrappingRainWater(heights6))
	// Expected: 0

	// Example 7: descending
	heights7 := []Height{
		{1, 5}, {2, 4}, {3, 3}, {4, 2}, {5, 1},
	}
	fmt.Println("Example 7 (descending):", calculateTrappingRainWater(heights7))
	// Expected: 0
}
```
