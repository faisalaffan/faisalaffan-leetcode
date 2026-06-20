# 3659 — Partition Array Into K Distinct Groups

## Deskripsi

**Soal:** [3659. Partition Array Into K Distinct Groups](https://leetcode.com/problems/partition-array-into-k-distinct-groups/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(max(nums))

**Algoritma:** —

**Fungsi Solusi:** `func partitionArrayIntoKDistinctGroups(nums []int, k int) bool`

## Solusi Go

```go
package main

// LeetCode #3659: Partition Array Into K-Distinct Groups
// https://leetcode.com/problems/partition-array-into-k-distinct-groups/
// Difficulty: Medium
// Time: O(n) | Space: O(max(nums))

import (
	"fmt"
	"slices"
)

func partitionArrayIntoKDistinctGroups(nums []int, k int) bool {
	n := len(nums)
	if n%k != 0 {
		return false
	}

	maxVal := slices.Max(nums)
  // Membuat slice untuk menyimpan hasil
	cnt := make([]int, maxVal+1)
	for _, x := range nums {
		cnt[x]++
		if cnt[x] > n/k {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(partitionArrayIntoKDistinctGroups([]int{1, 2, 3, 4}, 2))
	fmt.Println(partitionArrayIntoKDistinctGroups([]int{1, 1, 1, 1}, 2))
	fmt.Println(partitionArrayIntoKDistinctGroups([]int{1, 2, 2, 3, 3, 4}, 3))
}
```
