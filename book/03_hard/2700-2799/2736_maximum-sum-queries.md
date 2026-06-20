# 2736 — Maximum Sum Queries

## Deskripsi

**Soal:** [2736. Maximum Sum Queries](https://leetcode.com/problems/maximum-sum-queries/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Fenwick Tree (Binary Indexed Tree)

> **Ide Kunci:** Sort offline + BIT (Fenwick Tree) for prefix max on compressed nums2.

## Solusi Go

```go
package main

// LeetCode #2736: Maximum Sum Queries
// https://leetcode.com/problems/maximum-sum-queries/
// Difficulty: Hard
//
// Approach: Sort offline + BIT (Fenwick Tree) for prefix max on compressed nums2.
// Sort (nums1[i], nums2[i]) pairs by nums1 descending.
// Sort queries by x descending.
// Process queries in order, adding all pairs with nums1 >= x to BIT.
// BIT is indexed by nums2 values (compressed in reverse: larger nums2 -> smaller index),
// so querying prefix up to revComp(y) gives max sum for nums2 >= y.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1: nums1=[4,3,1,2], nums2=[2,4,9,5], queries=[[4,1],[1,3],[2,5]] -> [6,10,7]
	fmt.Println(maximumSumQueries([]int{4, 3, 1, 2}, []int{2, 4, 9, 5}, [][]int{{4, 1}, {1, 3}, {2, 5}}))
	// Example 2: nums1=[2,1], nums2=[3,3], queries=[[1,1]]
	fmt.Println(maximumSumQueries([]int{2, 1}, []int{3, 3}, [][]int{{1, 1}}))
}

func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	n := len(nums1)
	m := len(queries)

	// Pair nums1 and nums2, sorted by nums1 descending
  // Membuat slice untuk menyimpan hasil
	pairs := make([][3]int, n) // [nums1, nums2, nums1+nums2]
	for i := 0; i < n; i++ {
		pairs[i] = [3]int{nums1[i], nums2[i], nums1[i] + nums2[i]}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})

	// Attach original index to queries, sort by x descending
  // Membuat slice untuk menyimpan hasil
	qi := make([][3]int, m) // [x, y, originalIndex]
	for i, q := range queries {
		qi[i] = [3]int{q[0], q[1], i}
	}
	sort.Slice(qi, func(i, j int) bool {
		return qi[i][0] > qi[j][0]
	})

	// Coordinate compress nums2 values (sorted ascending)
  // Membuat slice untuk menyimpan hasil
	allVals := make([]int, 0, n+m)
	allVals = append(allVals, nums2...)
	for _, q := range queries {
		allVals = append(allVals, q[1])
	}
	sort.Ints(allVals)
  // Membuat slice untuk menyimpan hasil
	uniq := make([]int, 0, len(allVals))
	for i, v := range allVals {
		if i == 0 || v != allVals[i-1] {
			uniq = append(uniq, v)
		}
	}
	k := len(uniq)

	// BIT for prefix max.
	// Reverse compression: larger nums2 value -> smaller BIT index.
	// So query(revComp(y)) = max over all nums2 >= y.
  // Membuat slice untuk menyimpan hasil
	bit := make([]int, k+2)
  // Iterasi seluruh elemen
	for i := range bit {
		bit[i] = -1
	}

	update := func(idx, val int) {
		for idx < len(bit) {
			if val > bit[idx] {
				bit[idx] = val
			}
			idx += idx & -idx
		}
	}

	query := func(idx int) int {
		res := -1
		for idx > 0 {
			if bit[idx] > res {
				res = bit[idx]
			}
			idx -= idx & -idx
		}
		return res
	}

	// revComp: for a value v in uniq, return compressed index (1-based, reversed)
	// Largest value -> index 1, smallest -> index k
	revComp := func(v int) int {
		idx := sort.SearchInts(uniq, v)
		return k - idx
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, m)
	ptr := 0

	for _, q := range qi {
		x, y, origIdx := q[0], q[1], q[2]

		// Add all pairs with nums1 >= x
		for ptr < n && pairs[ptr][0] >= x {
			update(revComp(pairs[ptr][1]), pairs[ptr][2])
			ptr++
		}

		// Find first value >= y in uniq
		pos := sort.SearchInts(uniq, y)
		if pos == k {
			ans[origIdx] = -1
		} else {
			ans[origIdx] = query(k - pos)
		}
	}

	return ans
}
```
