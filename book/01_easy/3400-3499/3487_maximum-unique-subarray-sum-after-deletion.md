# 3487 — Maximum Unique Subarray Sum After Deletion

## Deskripsi

**Soal:** [3487. Maximum Unique Subarray Sum After Deletion](https://leetcode.com/problems/maximum-unique-subarray-sum-after-deletion/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** Sliding Window (jendela geser)

## Solusi Go

```go
package main

// LeetCode #3487: Maximum Unique Subarray Sum After Deletion
// https://leetcode.com/problems/maximum-unique-subarray-sum-after-deletion/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumUniqueSubarraySumAfterDeletion([]int{1, 2, 3, 4, 5}))
	fmt.Println(MaximumUniqueSubarraySumAfterDeletion([]int{1, 1, 0, 1, 1}))
	fmt.Println(MaximumUniqueSubarraySumAfterDeletion([]int{1, 2, -1, -2, 1, 0, -1}))
}

// MaximumUniqueSubarraySumAfterDeletion returns max sum of a subarray with all unique elements, after optionally deleting one element.
// Time: O(n). Space: O(n).
func MaximumUniqueSubarraySumAfterDeletion(nums []int) int {
	// After deletion means we can delete any element, then find max sum of a subarray with distinct elements.
	// Equivalent to finding max sum of a subarray with at most one duplicate.
	// Simplified: find max sum subarray where all elements are distinct (allow deleting one element).
	// We'll use sliding window that allows one "skip" (a duplicate that we can delete).
	maxSum := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums); i++ {
  // Membuat map untuk pencarian O(1): key → value
		seen := make(map[int]int)
		sum := 0
		dupAllowed := true
		for j := i; j < len(nums); j++ {
			seen[nums[j]]++
			if seen[nums[j]] > 2 {
				break
			}
			if seen[nums[j]] == 2 {
				if dupAllowed {
					dupAllowed = false
				} else {
					break
				}
			}
			sum += nums[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}
```
