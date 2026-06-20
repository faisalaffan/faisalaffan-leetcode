# 2529 — Maximum Count Of Positive Integer And Negative Integer

## Deskripsi

**Soal:** [2529. Maximum Count Of Positive Integer And Negative Integer](https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2529: Maximum Count of Positive Integer and Negative Integer
// https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(MaximumCountOfPositiveIntegerAndNegativeInteger([]int{-2, -1, -1, 1, 2, 3})) // 3
	fmt.Println(MaximumCountOfPositiveIntegerAndNegativeInteger([]int{-3, -2, -1, 0, 0, 1, 2})) // 3
}

func MaximumCountOfPositiveIntegerAndNegativeInteger(nums []int) int {
	pos, neg := 0, 0
	for _, n := range nums {
		if n > 0 {
			pos++
		} else if n < 0 {
			neg++
		}
	}
	if pos > neg {
		return pos
	}
	return neg
}
```
