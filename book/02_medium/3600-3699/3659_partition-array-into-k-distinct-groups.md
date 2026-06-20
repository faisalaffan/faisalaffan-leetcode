# 3659 — Partition Array Into K Distinct Groups

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func partitionArrayIntoKDistinctGroups(nums []int, k int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(max(nums))


## 💻 Solusi Go

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
  // Alokasi slice
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
