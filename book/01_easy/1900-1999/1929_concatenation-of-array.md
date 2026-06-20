# 1929 — Concatenation Of Array

## Deskripsi

**Soal:** [1929. Concatenation Of Array](https://leetcode.com/problems/concatenation-of-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1929: Concatenation of Array
// https://leetcode.com/problems/concatenation-of-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ConcatenationOfArray([]int{1, 2, 1}))       // [1,2,1,1,2,1]
	fmt.Println(ConcatenationOfArray([]int{1, 3, 2, 1}))    // [1,3,2,1,1,3,2,1]
}

// Time: O(n), Space: O(n)
func ConcatenationOfArray(nums []int) []int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, 2*n)
	for i, v := range nums {
		ans[i] = v
		ans[i+n] = v
	}
	return ans
}
```
