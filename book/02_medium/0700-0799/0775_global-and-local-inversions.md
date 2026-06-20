# 0775 — Global And Local Inversions

## Deskripsi

**Soal:** [0775. Global And Local Inversions](https://leetcode.com/problems/global-and-local-inversions/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #775: Global and Local Inversions
// https://leetcode.com/problems/global-and-local-inversions/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(isIdealPermutation([]int{1, 0, 2}))
	fmt.Println(isIdealPermutation([]int{1, 2, 0}))
}

func isIdealPermutation(nums []int) bool {
	maxVal := -1
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
		if maxVal > nums[i+2] {
			return false
		}
	}
	return true
}
```
