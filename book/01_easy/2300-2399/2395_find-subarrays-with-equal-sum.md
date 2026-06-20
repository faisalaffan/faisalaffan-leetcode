# 2395 — Find Subarrays With Equal Sum

## Deskripsi

**Soal:** [2395. Find Subarrays With Equal Sum](https://leetcode.com/problems/find-subarrays-with-equal-sum/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2395: Find Subarrays With Equal Sum
// https://leetcode.com/problems/find-subarrays-with-equal-sum/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(FindSubarraysWithEqualSum([]int{4, 2, 4}))   // true
	fmt.Println(FindSubarraysWithEqualSum([]int{1, 2, 3, 4, 5})) // false
}

func FindSubarraysWithEqualSum(nums []int) bool {
	seen := map[int]bool{}
	for i := 1; i < len(nums); i++ {
		sum := nums[i-1] + nums[i]
		if seen[sum] {
			return true
		}
		seen[sum] = true
	}
	return false
}
```
