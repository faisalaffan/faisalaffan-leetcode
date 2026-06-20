# 2122 — Recover The Original Array

## Deskripsi

**Soal:** [2122. Recover The Original Array](https://leetcode.com/problems/recover-the-original-array/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** Sort + frequency map.

## Solusi Go

```go
package main

// LeetCode #2122: Recover the Original Array
// https://leetcode.com/problems/recover-the-original-array/
// Difficulty: Hard
//
// Approach: Sort + frequency map.
// Sort the array. The smallest element must be original[0] - k.
// Try each possible partner for nums[0] as the "high" pair (original[0] + k).
// For each valid 2k = diff, greedily pair elements using a frequency map.
// Return the first valid original array found.

import (
	"fmt"
	"sort"
)

func main() {
	// Example from problem statement
	nums1 := []int{2, 10, 6, 4, 8, 12}
	fmt.Printf("recoverArray(%v) = %v (expected [3 7 11] or similar)\n", nums1, recoverArray(nums1))

	// Additional tests
	nums2 := []int{1, 1, 3, 3}
	fmt.Printf("recoverArray(%v) = %v (expected [2 2])\n", nums2, recoverArray(nums2))

	nums3 := []int{5, 5, 9, 9}
	fmt.Printf("recoverArray(%v) = %v\n", nums3, recoverArray(nums3))
}

func recoverArray(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)

	// nums[0] is always a "low" element (original[0] - k)
	// Try each possible partner for nums[0] to determine 2k = diff
	for i := 1; i < n; i++ {
		diff := nums[i] - nums[0]
		if diff == 0 || diff%2 != 0 {
			continue
		}

		// Try to reconstruct with this 2k value
  // Membuat map untuk pencarian O(1): key → value
		freq := make(map[int]int)
		for _, v := range nums {
			freq[v]++
		}

  // Membuat slice untuk menyimpan hasil
		result := make([]int, 0, n/2)
		valid := true

		for _, v := range nums {
			if freq[v] == 0 {
				continue
			}
			// v must be a "low" element; pair with v + diff (if diff = 2k)
			high := v + diff
			if freq[high] == 0 {
				valid = false
				break
			}
			freq[v]--
			freq[high]--
			result = append(result, v+diff/2)
		}

		if valid && len(result) == n/2 {
			return result
		}
	}

	return nil
}
```
