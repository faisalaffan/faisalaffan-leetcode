# 2869 — Minimum Operations To Collect Elements

## Deskripsi

**Soal:** [2869. Minimum Operations To Collect Elements](https://leetcode.com/problems/minimum-operations-to-collect-elements/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(k)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2869: Minimum Operations to Collect Elements
// https://leetcode.com/problems/minimum-operations-to-collect-elements/
// Difficulty: Easy
// Time: O(n) | Space: O(k)

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToCollectElements([]int{3, 1, 5, 4, 2}, 2))
	fmt.Println(MinimumOperationsToCollectElements([]int{3, 1, 5, 4, 2}, 5))
}

func MinimumOperationsToCollectElements(nums []int, k int) int {
  // Membuat slice untuk menyimpan hasil
	seen := make([]bool, k+1)
	collected := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] >= 1 && nums[i] <= k && !seen[nums[i]] {
			seen[nums[i]] = true
			collected++
		}
		if collected == k {
			return len(nums) - i
		}
	}
	return len(nums)
}
```
