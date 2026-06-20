# 3101 — Count Alternating Subarrays

## Deskripsi

**Soal:** [3101. Count Alternating Subarrays](https://leetcode.com/problems/count-alternating-subarrays/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func countAlternatingSubarrays(nums []int) int64`

## Solusi Go

```go
package main

// LeetCode #3101: Count Alternating Subarrays
// https://leetcode.com/problems/count-alternating-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func countAlternatingSubarrays(nums []int) int64 {
	var ans int64
	n := len(nums)
	left := 0

	for right := 0; right < n; right++ {
		if right > 0 && nums[right] == nums[right-1] {
			left = right
		}
		ans += int64(right - left + 1)
	}

	return ans
}

func main() {
	fmt.Println(countAlternatingSubarrays([]int{0, 1, 1, 1})) // Expected: 5
	fmt.Println(countAlternatingSubarrays([]int{1, 0, 1, 0})) // Expected: 10
}
```
