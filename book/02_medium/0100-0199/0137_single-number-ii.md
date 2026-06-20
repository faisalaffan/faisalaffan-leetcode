# 0137 — Single Number Ii

## Deskripsi

**Soal:** [0137. Single Number Ii](https://leetcode.com/problems/single-number-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func singleNumber(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #137: Single Number II
// https://leetcode.com/problems/single-number-ii/
// Difficulty: Medium

import "fmt"

func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones = (ones ^ num) & ^twos
		twos = (twos ^ num) & ^ones
	}
	return ones
}

func main() {
	// Test case 1
	fmt.Println(singleNumber([]int{2, 2, 3, 2})) // 3

	// Test case 2
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99})) // 99

	// Test case 3
	fmt.Println(singleNumber([]int{30000, 500, 100, 30000, 100, 30000, 100})) // 500
}

// Time: O(n) | Space: O(1)
```
