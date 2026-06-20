# 2790 — Maximum Number Of Groups With Increasing Length

## Deskripsi

**Soal:** [2790. Maximum Number Of Groups With Increasing Length](https://leetcode.com/problems/maximum-number-of-groups-with-increasing-length/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Greedy (pemilihan optimal lokal)

**Fungsi Solusi:** `func maxIncreasingGroups(usageLimit []int) int`

> **Ide Kunci:** Sort usageLimit ascending. Maintain a running total of available

## Solusi Go

```go
package main

// LeetCode #2790: Maximum Number of Groups With Increasing Length
// https://leetcode.com/problems/maximum-number-of-groups-with-increasing-length/
// Difficulty: Hard
//
// Approach: Sort usageLimit ascending. Maintain a running total of available
// element uses. When total >= triangular number needed for (groups+1) groups,
// increment group count. Greedy is optimal: we can always rearrange elements
// into the required group sizes.

import (
	"fmt"
	"sort"
)

func maxIncreasingGroups(usageLimit []int) int {
	sort.Ints(usageLimit)
	total := 0
	groups := 0

	for _, limit := range usageLimit {
		total += limit
		// Need sum(1..(groups+1)) = (groups+1)*(groups+2)/2 total element-uses
		// to form groups+1 groups of sizes 1, 2, ..., groups+1
		needed := (groups + 1) * (groups + 2) / 2
		if total >= needed {
			groups++
		}
	}

	return groups
}

func main() {
	// Example 1: [1,2,5] -> 2
	fmt.Println(maxIncreasingGroups([]int{1, 2, 5}))
	// Example 2: [2,1,2] -> 2
	fmt.Println(maxIncreasingGroups([]int{2, 1, 2}))
	// Example 3: [1,1] -> 1
	fmt.Println(maxIncreasingGroups([]int{1, 1}))
	// All equal limits
	fmt.Println(maxIncreasingGroups([]int{2, 2, 2}))
	// Single element
	fmt.Println(maxIncreasingGroups([]int{5}))
	// Large limits
	fmt.Println(maxIncreasingGroups([]int{1, 1, 1, 1, 1}))
}
```
