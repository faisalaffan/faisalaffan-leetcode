# 3301 — Maximize The Total Height Of Unique Towers

## Deskripsi

**Soal:** [3301. Maximize The Total Height Of Unique Towers](https://leetcode.com/problems/maximize-the-total-height-of-unique-towers/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n) Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3301: Maximize the Total Height of Unique Towers
// https://leetcode.com/problems/maximize-the-total-height-of-unique-towers/
// Difficulty: Medium
// Time: O(n log n) Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumTotalSum([]int{2, 3, 4, 3}))   // 10
	fmt.Println(maximumTotalSum([]int{2, 2, 1}))      // -1
	fmt.Println(maximumTotalSum([]int{5, 4, 3, 2, 1})) // 15
}

func maximumTotalSum(maximumHeight []int) int64 {
	sort.Slice(maximumHeight, func(i, j int) bool {
		return maximumHeight[i] > maximumHeight[j]
	})

	var total int64
	prev := maximumHeight[0]
	total += int64(prev)

	for i := 1; i < len(maximumHeight); i++ {
		if prev <= 1 {
			return -1
		}
		h := maximumHeight[i]
		if h >= prev {
			h = prev - 1
		}
		total += int64(h)
		prev = h
	}
	return total
}
```
