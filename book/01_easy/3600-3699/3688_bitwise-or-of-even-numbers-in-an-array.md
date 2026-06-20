# 3688 — Bitwise Or Of Even Numbers In An Array

## Deskripsi

**Soal:** [3688. Bitwise Or Of Even Numbers In An Array](https://leetcode.com/problems/bitwise-or-of-even-numbers-in-an-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3688: Bitwise OR of Even Numbers in an Array
// https://leetcode.com/problems/bitwise-or-of-even-numbers-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(BitwiseOrOfEvenNumbersInAnArray([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(BitwiseOrOfEvenNumbersInAnArray([]int{7, 9, 11}))
	fmt.Println(BitwiseOrOfEvenNumbersInAnArray([]int{1, 8, 16}))
}

// Time: O(n)
// Space: O(1)
func BitwiseOrOfEvenNumbersInAnArray(nums []int) int {
	res := 0
	for _, n := range nums {
		if n%2 == 0 {
			res |= n
		}
	}
	return res
}
```
