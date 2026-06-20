# 2760 — Longest Even Odd Subarray With Threshold

## Deskripsi

**Soal:** [2760. Longest Even Odd Subarray With Threshold](https://leetcode.com/problems/longest-even-odd-subarray-with-threshold/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2760: Longest Even Odd Subarray With Threshold
// https://leetcode.com/problems/longest-even-odd-subarray-with-threshold/
// Difficulty: Easy
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(LongestEvenOddSubarrayWithThreshold([]int{3, 2, 5, 4}, 5))
	fmt.Println(LongestEvenOddSubarrayWithThreshold([]int{4, 5, 2, 1}, 4))
}

func LongestEvenOddSubarrayWithThreshold(nums []int, threshold int) int {
	maxLen := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 || nums[i] > threshold {
			continue
		}
		length := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > threshold {
				break
			}
			if nums[j]%2 == nums[j-1]%2 {
				break
			}
			length++
		}
		if length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}
```
