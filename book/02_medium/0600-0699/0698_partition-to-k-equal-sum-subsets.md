# 0698 — Partition To K Equal Sum Subsets

## Deskripsi

**Soal:** [0698. Partition To K Equal Sum Subsets](https://leetcode.com/problems/partition-to-k-equal-sum-subsets/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(k^(n-k) * n!) roughly  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #698: Partition to K Equal Sum Subsets
// https://leetcode.com/problems/partition-to-k-equal-sum-subsets/
// Difficulty: Medium
// Time: O(k^(n-k) * n!) roughly
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	fmt.Println(canPartitionKSubsets([]int{1, 2, 3, 4}, 3))
}

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%k != 0 {
		return false
	}

	target := sum / k
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

  // Membuat slice untuk menyimpan hasil
	used := make([]bool, len(nums))

	var backtrack func(start int, currentSum int, subsetsFormed int) bool
	backtrack = func(start int, currentSum int, subsetsFormed int) bool {
		if subsetsFormed == k {
			return true
		}
		if currentSum == target {
			return backtrack(0, 0, subsetsFormed+1)
		}

		for i := start; i < len(nums); i++ {
			if used[i] || currentSum+nums[i] > target {
				continue
			}
			// Pruning: skip duplicate values
			if i > 0 && !used[i-1] && nums[i] == nums[i-1] {
				continue
			}

			used[i] = true
			if backtrack(i+1, currentSum+nums[i], subsetsFormed) {
				return true
			}
			used[i] = false

			if currentSum == 0 {
				return false
			}
		}
		return false
	}

	return backtrack(0, 0, 0)
}
```
