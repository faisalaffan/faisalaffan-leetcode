# 2670 — Find The Distinct Difference Array

## Deskripsi

**Soal:** [2670. Find The Distinct Difference Array](https://leetcode.com/problems/find-the-distinct-difference-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2670: Find the Distinct Difference Array
// https://leetcode.com/problems/find-the-distinct-difference-array/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(FindTheDistinctDifferenceArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(FindTheDistinctDifferenceArray([]int{3, 2, 3, 4, 2}))
}

func FindTheDistinctDifferenceArray(nums []int) []int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	suffixDistinct := make([]int, n+1)
	seen := map[int]bool{}

	for i := n - 1; i >= 0; i-- {
		suffixDistinct[i] = suffixDistinct[i+1]
		if !seen[nums[i]] {
			seen[nums[i]] = true
			suffixDistinct[i]++
		}
	}

	seen = map[int]bool{}
	prefixDistinct := 0
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		if !seen[nums[i]] {
			seen[nums[i]] = true
			prefixDistinct++
		}
		ans[i] = prefixDistinct - suffixDistinct[i+1]
	}

	return ans
}
```
