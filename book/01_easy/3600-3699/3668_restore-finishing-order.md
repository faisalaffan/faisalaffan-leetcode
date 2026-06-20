# 3668 — Restore Finishing Order

## Deskripsi

**Soal:** [3668. Restore Finishing Order](https://leetcode.com/problems/restore-finishing-order/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3668: Restore Finishing Order
// https://leetcode.com/problems/restore-finishing-order/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(RestoreFinishingOrder([]int{3, 1, 2, 4, 5}, []int{2, 4}))
	fmt.Println(RestoreFinishingOrder([]int{1, 2, 3, 4, 5}, []int{1, 3, 5}))
}

// Time: O(n)
// Space: O(n)
func RestoreFinishingOrder(order []int, friends []int) []int {
  // Membuat map untuk pencarian O(1): key → value
	friendSet := make(map[int]bool)
	for _, f := range friends {
		friendSet[f] = true
	}

  // Membuat slice untuk menyimpan hasil
	res := make([]int, 0, len(friends))
	for _, id := range order {
		if friendSet[id] {
			res = append(res, id)
		}
	}
	return res
}

func init() {
	_ = sort.Ints
}
```
