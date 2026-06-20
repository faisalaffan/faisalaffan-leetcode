# 0633 — Sum Of Square Numbers

## Deskripsi

**Soal:** [0633. Sum Of Square Numbers](https://leetcode.com/problems/sum-of-square-numbers/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(sqrt(c))  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #633: Sum of Square Numbers
// https://leetcode.com/problems/sum-of-square-numbers/
// Difficulty: Medium
// Time: O(sqrt(c))
// Space: O(1)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(JudgeSquareSum(5))
	fmt.Println(JudgeSquareSum(3))
	fmt.Println(JudgeSquareSum(4))
	fmt.Println(JudgeSquareSum(2))
}

func JudgeSquareSum(c int) bool {
	left := 0
	right := int(math.Sqrt(float64(c)))

	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum < c {
			left++
		} else {
			right--
		}
	}

	return false
}
```
