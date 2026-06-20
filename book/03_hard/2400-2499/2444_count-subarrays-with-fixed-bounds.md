# 2444 — Count Subarrays With Fixed Bounds

## Deskripsi

**Soal:** [2444. Count Subarrays With Fixed Bounds](https://leetcode.com/problems/count-subarrays-with-fixed-bounds/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Sliding Window (jendela geser)

## Solusi Go

```go
package main

// LeetCode #2444: Count Subarrays With Fixed Bounds
// https://leetcode.com/problems/count-subarrays-with-fixed-bounds/
// Difficulty: Hard
//
// Two-pointer / sliding window. For each right index, track the last positions
// of minK and maxK. Any subarray ending at right that contains both minK and
// maxK starts at (min(lastMinK, lastMaxK) ... last-bad). The valid start is
// the last index of any element outside [minK, maxK].
// Time O(N) | Space O(1)

import "fmt"

func main() {
	// Example 1
	fmt.Println(countSubarrays([]int{1, 3, 5, 2, 7, 5}, 1, 5))
	// Example 2
	fmt.Println(countSubarrays([]int{1, 1, 1, 1}, 1, 1))
	// No valid subarray
	fmt.Println(countSubarrays([]int{2, 3, 4}, 1, 5))
}

func countSubarrays(nums []int, minK, maxK int) int64 {
	var ans int64
	lastMin, lastMax := -1, -1
	bad := -1 // last index of element outside [minK, maxK]

	for i, v := range nums {
		if v < minK || v > maxK {
			bad = i
			continue
		}
		if v == minK {
			lastMin = i
		}
		if v == maxK {
			lastMax = i
		}
		start := min(lastMin, lastMax)
		if start > bad {
			ans += int64(start - bad)
		}
	}
	return ans
}
```
