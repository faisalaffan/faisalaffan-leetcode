# 3906 — Count Good Integers On A Grid Path

## Deskripsi

**Soal:** [3906. Count Good Integers On A Grid Path](https://leetcode.com/problems/count-good-integers-on-a-grid-path/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** Generate all possible numbers along the path. For each

## Solusi Go

```go
package main

// LeetCode #3906: Count Good Integers on a Grid Path
// https://leetcode.com/problems/count-good-integers-on-a-grid-path/
// Difficulty: Hard
//
// Given range [l, r] and directions string (3 D's and 3 R's),
// trace a path on a 4x4 grid starting from (0,0). At each step,
// form a number by reading row and column indices as digits.
// Count numbers in [l, r] that appear at any position on the path.
//
// Approach: Generate all possible numbers along the path. For each
// position on the path, construct the number from its coordinates.
// Check if it falls in [l, r]. Return count.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countGoodIntegersOnPath(10, 50, "DDRRDR"))
	// Example 2
	fmt.Println(countGoodIntegersOnPath(1, 100, "DRDDRR"))
	// Edge: single number
	fmt.Println(countGoodIntegersOnPath(0, 5, "DDDRRR"))
}

func countGoodIntegersOnPath(l int64, r int64, directions string) int64 {
	// Generate path coordinates
	x, y := 0, 0
	coords := [][2]int{{x, y}}
	for _, ch := range directions {
		if ch == 'D' {
			x++
		} else if ch == 'R' {
			y++
		}
		coords = append(coords, [2]int{x, y})
	}

	// Collect all numbers formed at each position
	// A number is formed by concatenating row and column (e.g., (2,3) -> 23)
	var count int64
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[int64]bool)
	for _, c := range coords {
		num := int64(c[0]*10 + c[1]) // 2-digit number from coordinates
		if num >= l && num <= r && !seen[num] {
			seen[num] = true
			count++
		}
		// Also consider the number formed by column then row
		num2 := int64(c[1]*10 + c[0])
		if num2 >= l && num2 <= r && !seen[num2] {
			seen[num2] = true
			count++
		}
	}

	return count
}
```
