# 3012 — Minimize Length Of Array Using Operations

## Deskripsi

**Soal:** [3012. Minimize Length Of Array Using Operations](https://leetcode.com/problems/minimize-length-of-array-using-operations/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3012: Minimize Length of Array Using Operations
// https://leetcode.com/problems/minimize-length-of-array-using-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minimumArrayLength([]int{1, 4, 3, 1}))
	fmt.Println(minimumArrayLength([]int{5, 5, 5, 10, 5}))
	fmt.Println(minimumArrayLength([]int{3, 5}))
}

func minimumArrayLength(nums []int) int {
	minVal := nums[0]
	for _, x := range nums[1:] {
		if x < minVal {
			minVal = x
		}
	}
	cnt := 0
	for _, x := range nums {
		if x%minVal != 0 {
			return 1
		}
		if x == minVal {
			cnt++
		}
	}
	return (cnt + 1) / 2
}
```
