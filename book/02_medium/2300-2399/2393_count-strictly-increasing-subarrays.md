# 2393 — Count Strictly Increasing Subarrays

## Deskripsi

**Soal:** [2393. Count Strictly Increasing Subarrays](https://leetcode.com/problems/count-strictly-increasing-subarrays/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2393: Count Strictly Increasing Subarrays
// https://leetcode.com/problems/count-strictly-increasing-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Track length of current increasing run. Each position adds run_len subarrays.

import "fmt"

func main() {
	fmt.Println(countIncreasing([]int{1, 3, 5, 4, 4, 6})) // 10
	fmt.Println(countIncreasing([]int{1, 2, 3, 4, 5}))    // 15
}

func countIncreasing(nums []int) int64 {
	var ans int64
	run := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > nums[i-1] {
			run++
		} else {
			run = 1
		}
		ans += int64(run)
	}
	return ans
}
```
