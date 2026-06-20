# 3854 — Minimum Operations To Make Array Parity Alternating

## Deskripsi

**Soal:** [3854. Minimum Operations To Make Array Parity Alternating](https://leetcode.com/problems/minimum-operations-to-make-array-parity-alternating/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** Sliding Window (jendela geser)

**Fungsi Solusi:** `func MinimumOperationsToMakeArrayParityAlternating(nums []int) []int`

> **Ide Kunci:** Compute min operations for two patterns (even-start, odd-start).

## Solusi Go

```go
package main

// LeetCode #3854: Minimum Operations to Make Array Parity Alternating
// https://leetcode.com/problems/minimum-operations-to-make-array-parity-alternating/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Compute min operations for two patterns (even-start, odd-start).
// Then compute min max-min range using sliding window on candidate values.

import (
	"fmt"
	"math"
	"sort"
)

func MinimumOperationsToMakeArrayParityAlternating(nums []int) []int {
	n := len(nums)

	// Helper to compute for a given starting parity
	solve := func(startParity int) (ops int, candidates [][2]int) {
		curParity := startParity
		for i := 0; i < n; i++ {
			if nums[i]%2 != curParity {
				ops++
				candidates = append(candidates, [2]int{nums[i] - 1, i})
				candidates = append(candidates, [2]int{nums[i] + 1, i})
			} else {
				candidates = append(candidates, [2]int{nums[i], i})
			}
			curParity ^= 1
		}
		return
	}

	ops1, cand1 := solve(0) // even at index 0
	ops2, cand2 := solve(1) // odd at index 0

	// Choose pattern with fewer ops
	var ops int
	var candidates [][2]int
	if ops1 < ops2 || (ops1 == ops2) {
		ops = ops1
		candidates = cand1
	} else {
		ops = ops2
		candidates = cand2
	}

	// Sliding window to find min range
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i][0] < candidates[j][0]
	})

	minRange := math.MaxInt32
  // Membuat slice untuk menyimpan hasil
	cnt := make([]int, n)
	empty := n
	left := 0
	for right := 0; right < len(candidates); right++ {
		idx := candidates[right][1]
		if cnt[idx] == 0 {
			empty--
		}
		cnt[idx]++
		for empty == 0 {
			r := candidates[right][0] - candidates[left][0]
			if r < minRange {
				minRange = r
			}
			if cnt[candidates[left][1]] == 1 {
				empty++
			}
			cnt[candidates[left][1]]--
			left++
		}
	}

	return []int{ops, minRange}
}

func main() {
	// Example 1
	fmt.Println(MinimumOperationsToMakeArrayParityAlternating([]int{-2, -3, 1, 4})) // Expected: [2 6]

	// Example 2
	fmt.Println(MinimumOperationsToMakeArrayParityAlternating([]int{0, 2, -2})) // Expected: [1 3]

	// Example 3
	fmt.Println(MinimumOperationsToMakeArrayParityAlternating([]int{7})) // Expected: [0 0]
}
```
