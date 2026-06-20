# 0260 — Single Number Iii

## Deskripsi

**Soal:** [0260. Single Number Iii](https://leetcode.com/problems/single-number-iii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func singleNumber(nums []int) []int`

## Solusi Go

```go
package main

// LeetCode #260: Single Number III
// https://leetcode.com/problems/single-number-iii/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	diff := xor & -xor

	num1, num2 := 0, 0
	for _, num := range nums {
		if num&diff == 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}

	return []int{num1, num2}
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
	fmt.Println(singleNumber([]int{-1, 0}))
	fmt.Println(singleNumber([]int{0, 1}))
}
```
