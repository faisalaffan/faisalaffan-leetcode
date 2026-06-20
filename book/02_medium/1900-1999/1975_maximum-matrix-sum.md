# 1975 — Maximum Matrix Sum

## Deskripsi

**Soal:** [1975. Maximum Matrix Sum](https://leetcode.com/problems/maximum-matrix-sum/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m*n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1975: Maximum Matrix Sum
// https://leetcode.com/problems/maximum-matrix-sum/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxMatrixSum([][]int{{1, -1}, {-1, 1}}))
	fmt.Println(MaxMatrixSum([][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}))
}

// Time: O(m*n), Space: O(1)
func MaxMatrixSum(matrix [][]int) int64 {
	total := int64(0)
	negCount := 0
	minAbs := int64(1 << 31)

	for _, row := range matrix {
		for _, val := range row {
			if val < 0 {
				negCount++
			}
			abs := int64(val)
			if abs < 0 {
				abs = -abs
			}
			total += abs
			if abs < minAbs {
				minAbs = abs
			}
		}
	}

	if negCount%2 == 1 {
		total -= 2 * minAbs
	}
	return total
}
```
