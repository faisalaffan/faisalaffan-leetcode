# 0896 — Monotonic Array

## Deskripsi

**Soal:** [0896. Monotonic Array](https://leetcode.com/problems/monotonic-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #896: Monotonic Array
// https://leetcode.com/problems/monotonic-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))   // true
	fmt.Println(isMonotonic([]int{6, 5, 4, 4}))   // true
	fmt.Println(isMonotonic([]int{1, 3, 2}))      // false
	fmt.Println(isMonotonic([]int{1, 1, 1}))      // true
}

// isMonotonic checks if the array is monotonic (either non-decreasing or non-increasing).
// Time: O(n). Space: O(1).
func isMonotonic(nums []int) bool {
	inc, dec := true, true
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dec = false
		}
		if nums[i] < nums[i-1] {
			inc = false
		}
	}
	return inc || dec
}
```
