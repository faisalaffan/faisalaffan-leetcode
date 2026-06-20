# 1991 — Find The Middle Index In Array

## Deskripsi

**Soal:** [1991. Find The Middle Index In Array](https://leetcode.com/problems/find-the-middle-index-in-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1991: Find the Middle Index in Array
// https://leetcode.com/problems/find-the-middle-index-in-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheMiddleIndexInArray([]int{2, 3, -1, 8, 4}))   // 3
	fmt.Println(FindTheMiddleIndexInArray([]int{1, -1, 4}))          // 2
	fmt.Println(FindTheMiddleIndexInArray([]int{2, 5}))              // -1
}

// Time: O(n), Space: O(1)
func FindTheMiddleIndexInArray(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	leftSum := 0
	for i, v := range nums {
		if leftSum == total-leftSum-v {
			return i
		}
		leftSum += v
	}
	return -1
}
```
