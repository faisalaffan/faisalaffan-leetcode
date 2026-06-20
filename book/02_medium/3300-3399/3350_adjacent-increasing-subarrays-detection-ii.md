# 3350 — Adjacent Increasing Subarrays Detection Ii

## Deskripsi

**Soal:** [3350. Adjacent Increasing Subarrays Detection Ii](https://leetcode.com/problems/adjacent-increasing-subarrays-detection-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3350: Adjacent Increasing Subarrays Detection II
// https://leetcode.com/problems/adjacent-increasing-subarrays-detection-ii/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxIncreasingSubarrays([]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1})) // 3
	fmt.Println(maxIncreasingSubarrays([]int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7})) // 2
}

func maxIncreasingSubarrays(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	// len[i] = length of strictly increasing subarray ending at i
  // Membuat slice untuk menyimpan hasil
	lenEnd := make([]int, n)
	lenEnd[0] = 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			lenEnd[i] = lenEnd[i-1] + 1
		} else {
			lenEnd[i] = 1
		}
	}

	// For each position i, find max k such that:
	// subarray [i-k+1, i] is increasing AND [i+1, i+k] is increasing
	ans := 0
	for i := 0; i < n-1; i++ {
		// Current increasing subarray ending at i has length lenEnd[i]
		// Next increasing subarray starting at i+1 has length...
		// We can compute the length of the next increasing subarray starting at i+1
		k := 1
		maxK := 0
		for k <= lenEnd[i] && i+k < n {
			if lenEnd[i+k] >= k {
				maxK = k
			}
			k++
		}
		if maxK > ans {
			ans = maxK
		}
	}

	return ans
}
```
