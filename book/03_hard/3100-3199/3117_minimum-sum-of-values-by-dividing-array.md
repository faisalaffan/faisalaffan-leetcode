# 3117 — Minimum Sum Of Values By Dividing Array

## Deskripsi

**Soal:** [3117. Minimum Sum Of Values By Dividing Array](https://leetcode.com/problems/minimum-sum-of-values-by-dividing-array/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func minimumSumOfValuesByDividingArray(nums []int, andValues []int) int`

## Solusi Go

```go
package main

// LeetCode #3117: Minimum Sum of Values by Dividing Array
// https://leetcode.com/problems/minimum-sum-of-values-by-dividing-array/
// Difficulty: Hard
//
// Partition nums into m contiguous subarrays such that the AND of the i-th
// subarray equals andValues[i]. Minimize the sum of the last elements of each
// subarray. Use DP with maps tracking (completed_segments, current_AND) -> min_sum.

import (
	"fmt"
)

const ALL_ONES = (1 << 20) - 1

func minimumSumOfValuesByDividingArray(nums []int, andValues []int) int {
	m := len(andValues)
  // Membuat slice untuk menyimpan hasil
	dp := make([]map[int]int, m+1)
	for j := 0; j <= m; j++ {
		dp[j] = make(map[int]int)
	}
	dp[0][ALL_ONES] = 0

	for _, x := range nums {
  // Membuat slice untuk menyimpan hasil
		ndp := make([]map[int]int, m+1)
		for j := 0; j <= m; j++ {
			ndp[j] = make(map[int]int)
		}
		for j := 0; j <= m; j++ {
			for andVal, sum := range dp[j] {
				if andVal == ALL_ONES {
					// Start new segment at x, don't close
					if val, ok := ndp[j][x]; !ok || sum < val {
						ndp[j][x] = sum
					}
					// Start and immediately close (single element segment)
					if j < m && x == andValues[j] {
						if val, ok := ndp[j+1][ALL_ONES]; !ok || sum+x < val {
							ndp[j+1][ALL_ONES] = sum + x
						}
					}
				} else {
					newAnd := andVal & x
					// Extend, don't close
					if val, ok := ndp[j][newAnd]; !ok || sum < val {
						ndp[j][newAnd] = sum
					}
					// Extend and close
					if j < m && newAnd == andValues[j] {
						if val, ok := ndp[j+1][ALL_ONES]; !ok || sum+x < val {
							ndp[j+1][ALL_ONES] = sum + x
						}
					}
				}
			}
		}
		dp = ndp
	}

	if ans, ok := dp[m][ALL_ONES]; ok {
		return ans
	}
	return -1
}

func main() {
	// Test case 1
	nums := []int{1, 4, 3, 3, 2}
	andValues := []int{0, 3, 3, 2}
	fmt.Println("Test 1:", minimumSumOfValuesByDividingArray(nums, andValues))
	// Expected: 12

	// Test case 2: single segment
	nums2 := []int{1, 2, 3}
	andValues2 := []int{0}
	fmt.Println("Test 2:", minimumSumOfValuesByDividingArray(nums2, andValues2))
	// Expected: 3

	// Test case 3: impossible
	nums3 := []int{1, 2}
	andValues3 := []int{5}
	fmt.Println("Test 3:", minimumSumOfValuesByDividingArray(nums3, andValues3))
	// Expected: -1

	// Test case 4: all same
	nums4 := []int{7, 7, 7, 7}
	andValues4 := []int{7, 7}
	fmt.Println("Test 4:", minimumSumOfValuesByDividingArray(nums4, andValues4))
	// Expected: 7+7=14
}
```
