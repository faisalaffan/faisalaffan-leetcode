# 1051 — Height Checker

## Deskripsi

**Soal:** [1051. Height Checker](https://leetcode.com/problems/height-checker/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1051: Height Checker
// https://leetcode.com/problems/height-checker/
// Difficulty: Easy
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3})) // 3
	fmt.Println(heightChecker([]int{5, 1, 2, 3, 4}))    // 5
	fmt.Println(heightChecker([]int{1, 2, 3, 4, 5}))    // 0
}

// LeetCode submission: heightChecker
func heightChecker(heights []int) int {
  // Membuat slice untuk menyimpan hasil
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)
	count := 0
  // Iterasi seluruh elemen
	for i := range heights {
		if heights[i] != expected[i] {
			count++
		}
	}
	return count
}
```
