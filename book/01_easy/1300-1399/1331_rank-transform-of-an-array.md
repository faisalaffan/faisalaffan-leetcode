# 1331 — Rank Transform Of An Array

## Deskripsi

**Soal:** [1331. Rank Transform Of An Array](https://leetcode.com/problems/rank-transform-of-an-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func arrayRankTransform(arr []int) []int`

## Solusi Go

```go
package main

// LeetCode #1331: Rank Transform of an Array
// https://leetcode.com/problems/rank-transform-of-an-array/
// Difficulty: Easy
//
// LeetCode submission: func arrayRankTransform(arr []int) []int

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(RankTransformOfAnArray([]int{40, 10, 20, 30}))       // [4 1 2 3]
	fmt.Println(RankTransformOfAnArray([]int{100, 100, 100}))        // [1 1 1]
	fmt.Println(RankTransformOfAnArray([]int{37, 12, 28, 9, 100, 56})) // [5 3 4 1 6 2]
}

// Time: O(n log n), Space: O(n)
func RankTransformOfAnArray(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
  // Membuat slice untuk menyimpan hasil
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)

  // Membuat map untuk pencarian O(1): key → value
	rank := make(map[int]int, len(arr))
	cur := 1
	for _, v := range sorted {
		if _, seen := rank[v]; !seen {
			rank[v] = cur
			cur++
		}
	}

  // Membuat slice untuk menyimpan hasil
	res := make([]int, len(arr))
	for i, v := range arr {
		res[i] = rank[v]
	}
	return res
}
```
