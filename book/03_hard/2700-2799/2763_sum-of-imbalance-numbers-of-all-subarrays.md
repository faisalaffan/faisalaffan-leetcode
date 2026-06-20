# 2763 — Sum Of Imbalance Numbers Of All Subarrays

## Deskripsi

**Soal:** [2763. Sum Of Imbalance Numbers Of All Subarrays](https://leetcode.com/problems/sum-of-imbalance-numbers-of-all-subarrays/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** O(n^2) incremental.

## Solusi Go

```go
package main

// LeetCode #2763: Sum of Imbalance Numbers of All Subarrays
// https://leetcode.com/problems/sum-of-imbalance-numbers-of-all-subarrays/
// Difficulty: Hard
//
// Approach: O(n^2) incremental.
// For each left index, maintain a set of seen values and a running
// imbalance counter. When adding a new value v:
//   - If both v-1 and v+1 are seen: imbalance-- (v bridges a gap)
//   - If neither v-1 nor v+1 are seen: imbalance++ (v creates a new isolated point)
//   - Otherwise: no change (v fills one side of an existing gap)

import "fmt"

func main() {
	// Example 1: [2,3,1,4] -> 3
	fmt.Println(sumImbalanceNumbers([]int{2, 3, 1, 4}))
	// Example 2: [1,3,3,3,5] -> 8
	fmt.Println(sumImbalanceNumbers([]int{1, 3, 3, 3, 5}))
}

func sumImbalanceNumbers(nums []int) int {
	n := len(nums)
	ans := 0

	for i := 0; i < n; i++ {
  // Membuat map untuk pencarian O(1): key → value
		seen := make(map[int]bool)
		seen[nums[i]] = true
		imbalance := 0

		for j := i + 1; j < n; j++ {
			v := nums[j]
			if seen[v] {
				// Duplicate value doesn't change imbalance
				ans += imbalance
				continue
			}

			seenPrev := seen[v-1]
			seenNext := seen[v+1]

			if seenPrev && seenNext {
				imbalance-- // v bridges the gap
			} else if !seenPrev && !seenNext {
				imbalance++ // v creates a new isolated point
			}
			// else: v fills one side, no change

			seen[v] = true
			ans += imbalance
		}
	}

	return ans
}
```
