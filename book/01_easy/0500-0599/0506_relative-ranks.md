# 0506 — Relative Ranks

## Deskripsi

**Soal:** [0506. Relative Ranks](https://leetcode.com/problems/relative-ranks/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func RelativeRanks(score []int) []string`

## Solusi Go

```go
package main

// LeetCode #506: Relative Ranks
// https://leetcode.com/problems/relative-ranks/
// Difficulty: Easy

import (
	"fmt"
	"sort"
	"strconv"
)

// Time: O(n log n), Space: O(n)
func RelativeRanks(score []int) []string {
  // Membuat slice untuk menyimpan hasil
	sorted := make([]int, len(score))
	copy(sorted, score)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
  // Membuat map untuk pencarian O(1): key → value
	rank := make(map[int]string)
	for i, s := range sorted {
		switch i {
		case 0:
			rank[s] = "Gold Medal"
		case 1:
			rank[s] = "Silver Medal"
		case 2:
			rank[s] = "Bronze Medal"
		default:
			rank[s] = strconv.Itoa(i + 1)
		}
	}
  // Membuat slice untuk menyimpan hasil
	result := make([]string, len(score))
	for i, s := range score {
		result[i] = rank[s]
	}
	return result
}

func main() {
	fmt.Println(RelativeRanks([]int{5, 4, 3, 2, 1}))
	fmt.Println(RelativeRanks([]int{10, 3, 8, 9, 4}))
}
```
