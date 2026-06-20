# 0442 — Find All Duplicates In An Array

## Deskripsi

**Soal:** [0442. Find All Duplicates In An Array](https://leetcode.com/problems/find-all-duplicates-in-an-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func findDuplicates(nums []int) []int`

## Solusi Go

```go
package main

// LeetCode #442: Find All Duplicates in an Array
// https://leetcode.com/problems/find-all-duplicates-in-an-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func findDuplicates(nums []int) []int {
	result := []int{}
	for _, num := range nums {
		idx := num
		if idx < 0 {
			idx = -idx
		}
		idx--
		if nums[idx] < 0 {
			result = append(result, idx+1)
		} else {
			nums[idx] = -nums[idx]
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	// Expected: [2, 3]

	// Test case 2
	fmt.Println("Test 2:", findDuplicates([]int{1, 1, 2}))
	// Expected: [1]

	// Test case 3
	fmt.Println("Test 3:", findDuplicates([]int{1}))
	// Expected: []
}
```
