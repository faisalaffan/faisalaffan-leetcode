# 3349 — Adjacent Increasing Subarrays Detection I

## Deskripsi

**Soal:** [3349. Adjacent Increasing Subarrays Detection I](https://leetcode.com/problems/adjacent-increasing-subarrays-detection-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3349: Adjacent Increasing Subarrays Detection I
// https://leetcode.com/problems/adjacent-increasing-subarrays-detection-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(AdjacentIncreasingSubarraysDetectionI([]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}, 3))
	fmt.Println(AdjacentIncreasingSubarraysDetectionI([]int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}, 5))
}

// AdjacentIncreasingSubarraysDetectionI returns true if there exist two adjacent k-length increasing subarrays.
// Time: O(n). Space: O(n).
func AdjacentIncreasingSubarraysDetectionI(nums []int, k int) bool {
	n := len(nums)
	if n < 2*k {
		return false
	}

	// inc[i] = true if subarray starting at i of length k is strictly increasing
  // Membuat slice untuk menyimpan hasil
	inc := make([]bool, n-k+1)
	for i := 0; i <= n-k; i++ {
		isInc := true
		for j := i; j < i+k-1; j++ {
			if nums[j] >= nums[j+1] {
				isInc = false
				break
			}
		}
		inc[i] = isInc
	}

	for i := 0; i <= n-2*k; i++ {
		if inc[i] && inc[i+k] {
			return true
		}
	}
	return false
}
```
