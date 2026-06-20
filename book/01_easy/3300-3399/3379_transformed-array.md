# 3379 — Transformed Array

## Deskripsi

**Soal:** [3379. Transformed Array](https://leetcode.com/problems/transformed-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3379: Transformed Array
// https://leetcode.com/problems/transformed-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TransformedArray([]int{3, -2, 1, 1}))
	fmt.Println(TransformedArray([]int{-1, 4, -1}))
}

// TransformedArray constructs a new array where result[i] = nums[(i + nums[i]) mod n], handling negative wrap-around.
// Time: O(n). Space: O(n).
func TransformedArray(nums []int) []int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	result := make([]int, n)
	for i, val := range nums {
		idx := (i + val) % n
		if idx < 0 {
			idx += n
		}
		result[i] = nums[idx]
	}
	return result
}
```
