# 2653 — Sliding Subarray Beauty

## Deskripsi

**Soal:** [2653. Sliding Subarray Beauty](https://leetcode.com/problems/sliding-subarray-beauty/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * 50)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func getSubarrayBeauty(nums []int, k int, x int) []int`

## Solusi Go

```go
package main

// LeetCode #2653: Sliding Subarray Beauty
// https://leetcode.com/problems/sliding-subarray-beauty/
// Difficulty: Medium
// Time: O(n * 50) | Space: O(1)

import "fmt"

func getSubarrayBeauty(nums []int, k int, x int) []int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, n-k+1)
  // Membuat slice untuk menyimpan hasil
	freq := make([]int, 101) // values in [-50, 50] shifted by 50

	for i := 0; i < k; i++ {
		freq[nums[i]+50]++
	}

	for i := k; i <= n; i++ {
		// Find x-th smallest
		count := 0
		val := 0
		for j := 0; j <= 100; j++ {
			count += freq[j]
			if count >= x {
				val = j - 50
				break
			}
		}
		// Only negative values qualify as "beauty"
		if val < 0 {
			ans[i-k] = val
		} else {
			ans[i-k] = 0
		}

		if i < n {
			freq[nums[i-k]+50]--
			freq[nums[i]+50]++
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", getSubarrayBeauty([]int{1, -1, -3, -2, 3}, 3, 2))
	// Expected: [-1,-2,-2]

	// Test case 2
	fmt.Println("Test 2:", getSubarrayBeauty([]int{-1, -2, -3, -4, -5}, 2, 2))
	// Expected: [-1,-2,-3,-4]

	// Test case 3
	fmt.Println("Test 3:", getSubarrayBeauty([]int{-3, 1, 2, -3, 0, -3}, 2, 1))
	// Expected: [-3,0,-3,-3,-3]
}
```
