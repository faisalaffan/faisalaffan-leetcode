# 3468 — Find The Number Of Copy Arrays

## Deskripsi

**Soal:** [3468. Find The Number Of Copy Arrays](https://leetcode.com/problems/find-the-number-of-copy-arrays/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func countArrays(original []int, bounds [][]int) int`

## Solusi Go

```go
package main

// LeetCode #3468: Find the Number of Copy Arrays
// https://leetcode.com/problems/find-the-number-of-copy-arrays/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import (
	"fmt"
	"math"
)

func countArrays(original []int, bounds [][]int) int {
	n := len(original)
	lo := math.MinInt64
	hi := math.MaxInt64
	for i := 0; i < n; i++ {
		diff := original[i] - original[0]
		l := bounds[i][0] - diff
		r := bounds[i][1] - diff
		if l > lo {
			lo = l
		}
		if r < hi {
			hi = r
		}
	}
	if hi < lo {
		return 0
	}
	return hi - lo + 1
}

func main() {
	fmt.Println(countArrays([]int{1, 2, 3, 4}, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}})) // 2
	fmt.Println(countArrays([]int{1, 2, 3}, [][]int{{1, 3}, {2, 4}, {3, 5}})) // 3
}
```
