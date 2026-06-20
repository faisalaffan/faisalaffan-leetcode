# 3645 — Maximum Total From Optimal Activation Order

## Deskripsi

**Soal:** [3645. Maximum Total From Optimal Activation Order](https://leetcode.com/problems/maximum-total-from-optimal-activation-order/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func maximumTotalFromOptimalActivationOrder(value []int, limit []int) int64`

## Solusi Go

```go
package main

// LeetCode #3645: Maximum Total from Optimal Activation Order
// https://leetcode.com/problems/maximum-total-from-optimal-activation-order/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maximumTotalFromOptimalActivationOrder(value []int, limit []int) int64 {
	n := len(value)
  // Membuat slice 2D untuk DP/tabel
	groups := make([][]int, n+1)
	for i := 0; i < n; i++ {
		l := limit[i]
		groups[l] = append(groups[l], value[i])
	}

	var total int64 = 0
	for l := 1; l <= n; l++ {
		if len(groups[l]) == 0 {
			continue
		}
		sort.Slice(groups[l], func(i, j int) bool {
			return groups[l][i] > groups[l][j]
		})
		cap := l
		if len(groups[l]) < cap {
			cap = len(groups[l])
		}
		for i := 0; i < cap; i++ {
			total += int64(groups[l][i])
		}
	}
	return total
}

func main() {
	fmt.Println(maximumTotalFromOptimalActivationOrder([]int{3, 5, 2, 4}, []int{2, 1, 3, 2}))
	fmt.Println(maximumTotalFromOptimalActivationOrder([]int{10, 20}, []int{1, 1}))
	fmt.Println(maximumTotalFromOptimalActivationOrder([]int{1, 2, 3, 4, 5}, []int{1, 2, 2, 3, 3}))
}
```
