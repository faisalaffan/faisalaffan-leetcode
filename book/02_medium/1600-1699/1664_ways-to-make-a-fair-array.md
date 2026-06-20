# 1664 — Ways To Make A Fair Array

## Deskripsi

**Soal:** [1664. Ways To Make A Fair Array](https://leetcode.com/problems/ways-to-make-a-fair-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1664: Ways to Make a Fair Array
// https://leetcode.com/problems/ways-to-make-a-fair-array/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(WaysToMakeFair([]int{2, 1, 6, 4}))
	fmt.Println(WaysToMakeFair([]int{1, 1, 1}))
	fmt.Println(WaysToMakeFair([]int{1, 2, 3, 4, 5}))
}

func WaysToMakeFair(nums []int) int {
	// Time: O(N), Space: O(1)

	// Calculate total sum at even and odd indices
	totalEven := 0
	totalOdd := 0
	for i, num := range nums {
		if i%2 == 0 {
			totalEven += num
		} else {
			totalOdd += num
		}
	}

	result := 0
	prefixEven := 0
	prefixOdd := 0

	for i, num := range nums {
		if i%2 == 0 {
			totalEven -= num
		} else {
			totalOdd -= num
		}

		// After removing nums[i], all indices shift:
		// Elements to the right of i swap parity
		// Even sum = prefixEven + totalOdd
		// Odd sum = prefixOdd + totalEven
		if prefixEven+totalOdd == prefixOdd+totalEven {
			result++
		}

		if i%2 == 0 {
			prefixEven += num
		} else {
			prefixOdd += num
		}
	}

	return result
}
```
