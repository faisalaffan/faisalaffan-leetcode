# 2659 — Make Array Empty

## Deskripsi

**Soal:** [2659. Make Array Empty](https://leetcode.com/problems/make-array-empty/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Fenwick Tree (Binary Indexed Tree)

> **Ide Kunci:** BIT (Fenwick Tree) + sorted order.

## Solusi Go

```go
package main

// LeetCode #2659: Make Array Empty
// https://leetcode.com/problems/make-array-empty/
// Difficulty: Hard
//
// Approach: BIT (Fenwick Tree) + sorted order.
// Process elements in increasing value order. Use BIT to track which
// positions remain in the array. The number of "move-to-end" operations
// needed is the count of remaining elements between the current front
// and the next minimum element's position. Total operations = moves + n.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1: [3,4,-1] -> 5
	fmt.Println(countOperationsToMakeArrayEmpty([]int{3, 4, -1}))
	// Example 2: [1,2,4,3] -> 5
	fmt.Println(countOperationsToMakeArrayEmpty([]int{1, 2, 4, 3}))
	// Example 3: [1,2,3] -> 3
	fmt.Println(countOperationsToMakeArrayEmpty([]int{1, 2, 3}))
}

func countOperationsToMakeArrayEmpty(nums []int) int64 {
	n := len(nums)

	type pair struct {
		val, idx int
	}
  // Membuat slice untuk menyimpan hasil
	sorted := make([]pair, n)
	for i, v := range nums {
		sorted[i] = pair{v, i}
	}
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].val != sorted[j].val {
			return sorted[i].val < sorted[j].val
		}
		return sorted[i].idx < sorted[j].idx
	})

	// BIT: 1 = element present, 0 = removed
  // Membuat slice untuk menyimpan hasil
	bit := make([]int, n+1)
	add := func(idx int, v int) {
		for idx++; idx <= n; idx += idx & -idx {
			bit[idx] += v
		}
	}
	sum := func(idx int) int {
		s := 0
		for idx++; idx > 0; idx -= idx & -idx {
			s += bit[idx]
		}
		return s
	}
	rangeSum := func(l, r int) int {
		if l > r {
			return 0
		}
		return sum(r) - sum(l-1)
	}

	for i := 0; i < n; i++ {
		add(i, 1)
	}

	var moves int64
	prev := 0

	for _, p := range sorted {
		pos := p.idx

		if pos >= prev {
			moves += int64(rangeSum(prev, pos-1))
		} else {
			moves += int64(rangeSum(prev, n-1) + rangeSum(0, pos-1))
		}

		add(pos, -1) // remove this element
		prev = pos
	}

	return moves + int64(n)
}
```
