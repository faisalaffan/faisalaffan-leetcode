# 3529 — Count Cells In Overlapping Horizontal And Vertical Substrings

## Deskripsi

**Soal:** [3529. Count Cells In Overlapping Horizontal And Vertical Substrings](https://leetcode.com/problems/count-cells-in-overlapping-horizontal-and-vertical-substrings/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3529: Count Cells in Overlapping Horizontal and Vertical Substrings
// https://leetcode.com/problems/count-cells-in-overlapping-horizontal-and-vertical-substrings/
// Difficulty: Medium
// Complexity: O(h*w) time, O(h*w) space

import "fmt"

func main() {
	// Test case 1
	h := []int{0, 2}
	v := []int{0, 2}
	fmt.Println("Test 1:", CountCellsInOverlappingHorizontalAndVerticalSubstrings(3, 3, h, v))

	// Test case 2
	h2 := []int{0, 1}
	v2 := []int{0, 1}
	fmt.Println("Test 2:", CountCellsInOverlappingHorizontalAndVerticalSubstrings(3, 3, h2, v2))

	// Test case 3
	h3 := []int{0}
	v3 := []int{0}
	fmt.Println("Test 3:", CountCellsInOverlappingHorizontalAndVerticalSubstrings(1, 1, h3, v3))
}

func CountCellsInOverlappingHorizontalAndVerticalSubstrings(rows, cols int, horizontal, vertical []int) int {
	// Count cells that are in the intersection of selected horizontal and vertical ranges
  // Membuat map untuk pencarian O(1): key → value
	hSet := make(map[int]bool)
	for _, h := range horizontal {
		hSet[h] = true
	}
  // Membuat map untuk pencarian O(1): key → value
	vSet := make(map[int]bool)
	for _, v := range vertical {
		vSet[v] = true
	}
	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if hSet[r] && vSet[c] {
				count++
			}
		}
	}
	return count
}
```
