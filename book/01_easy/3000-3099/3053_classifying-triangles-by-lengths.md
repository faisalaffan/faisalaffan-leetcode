# 3053 — Classifying Triangles By Lengths

## Deskripsi

**Soal:** [3053. Classifying Triangles By Lengths](https://leetcode.com/problems/classifying-triangles-by-lengths/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3053: Classifying Triangles by Lengths
// https://leetcode.com/problems/classifying-triangles-by-lengths/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent logic: classify triangles by side lengths.

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: triangleType
	// Input: [side_a, side_b, side_c]
	fmt.Println(ClassifyingTrianglesByLengths([][]int{{3, 3, 3}, {3, 4, 5}, {3, 3, 5}, {1, 2, 3}}))
	// [Equilateral Scalene Isosceles None]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: triangleType
func ClassifyingTrianglesByLengths(triangles [][]int) []string {
  // Membuat slice untuk menyimpan hasil
	result := make([]string, len(triangles))
	for i, t := range triangles {
		sides := []int{t[0], t[1], t[2]}
		sort.Ints(sides)
		a, b, c := sides[0], sides[1], sides[2]

		if a+b <= c {
			result[i] = "None"
		} else if a == b && b == c {
			result[i] = "Equilateral"
		} else if a == b || b == c || a == c {
			result[i] = "Isosceles"
		} else {
			result[i] = "Scalene"
		}
	}
	return result
}
```
