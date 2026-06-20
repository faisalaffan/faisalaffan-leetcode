# 2221 — Find Triangular Sum Of An Array

## Deskripsi

**Soal:** [2221. Find Triangular Sum Of An Array](https://leetcode.com/problems/find-triangular-sum-of-an-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func triangularSum(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #2221: Find Triangular Sum of an Array
// https://leetcode.com/problems/find-triangular-sum-of-an-array/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func triangularSum(nums []int) int {
	n := len(nums)
	for n > 1 {
		for i := 0; i < n-1; i++ {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
		n--
	}
	return nums[0]
}

func main() {
	// Test case 1
	fmt.Println(triangularSum([]int{1, 2, 3, 4, 5}))
	// Expected: 8

	// Test case 2
	fmt.Println(triangularSum([]int{5}))
	// Expected: 5

	// Test case 3
	fmt.Println(triangularSum([]int{2, 6, 6, 6}))
	// Expected: 4
}
```
