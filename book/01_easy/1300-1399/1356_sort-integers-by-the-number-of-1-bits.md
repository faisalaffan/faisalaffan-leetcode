# 1356 — Sort Integers By The Number Of 1 Bits

## Deskripsi

**Soal:** [1356. Sort Integers By The Number Of 1 Bits](https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func sortByBits(arr []int) []int`

## Solusi Go

```go
package main

// LeetCode #1356: Sort Integers by The Number of 1 Bits
// https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/
// Difficulty: Easy
//
// LeetCode submission: func sortByBits(arr []int) []int

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortIntegersByTheNumberOfOneBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})) // [0 1 2 4 8 3 5 6 7]
	fmt.Println(SortIntegersByTheNumberOfOneBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1})) // [1 2 4 8 16 32 64 128 256 512 1024]
}

// Time: O(n log n), Space: O(1)
func SortIntegersByTheNumberOfOneBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		bi, bj := popcount(arr[i]), popcount(arr[j])
		if bi != bj {
			return bi < bj
		}
		return arr[i] < arr[j]
	})
	return arr
}

func popcount(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}
```
