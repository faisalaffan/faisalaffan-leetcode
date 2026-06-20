# 2905 — Find Indices With Index And Value Difference Ii

## Deskripsi

**Soal:** [2905. Find Indices With Index And Value Difference Ii](https://leetcode.com/problems/find-indices-with-index-and-value-difference-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2905: Find Indices With Index and Value Difference II
// https://leetcode.com/problems/find-indices-with-index-and-value-difference-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(findIndices([]int{5, 1, 4, 1}, 2, 4))
	fmt.Println(findIndices([]int{2, 1}, 0, 0))
	fmt.Println(findIndices([]int{1, 2, 3}, 2, 4))
}

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	mi, mx := 0, 0
	for i := indexDifference; i < len(nums); i++ {
		j := i - indexDifference
		if nums[j] < nums[mi] {
			mi = j
		}
		if nums[j] > nums[mx] {
			mx = j
		}
		if nums[i]-nums[mi] >= valueDifference {
			return []int{mi, i}
		}
		if nums[mx]-nums[i] >= valueDifference {
			return []int{mx, i}
		}
	}
	return []int{-1, -1}
}
```
