# 0268 — Missing Number

## Deskripsi

**Soal:** [0268. Missing Number](https://leetcode.com/problems/missing-number/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func MissingNumber(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #268: Missing Number
// https://leetcode.com/problems/missing-number/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func MissingNumber(nums []int) int {
	n := len(nums)
	result := n
	for i, v := range nums {
		result ^= i ^ v
	}
	return result
}

func main() {
	fmt.Println(MissingNumber([]int{3, 0, 1}))
	fmt.Println(MissingNumber([]int{0, 1}))
	fmt.Println(MissingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
```
