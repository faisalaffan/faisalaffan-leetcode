# 1200 — Minimum Absolute Difference

## Deskripsi

**Soal:** [1200. Minimum Absolute Difference](https://leetcode.com/problems/minimum-absolute-difference/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1200: Minimum Absolute Difference
// https://leetcode.com/problems/minimum-absolute-difference/
// Difficulty: Easy
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))       // [[1,2],[2,3],[3,4]]
	fmt.Println(minimumAbsDifference([]int{1, 3, 6, 10, 15}))  // [[1,3]]
}

// LeetCode submission: minimumAbsDifference
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	minDiff := 1 << 31
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}
	var ans [][]int
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == minDiff {
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
	}
	return ans
}
```
