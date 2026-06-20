# 3462 — Maximum Sum With At Most K Elements

## Deskripsi

**Soal:** [3462. Maximum Sum With At Most K Elements](https://leetcode.com/problems/maximum-sum-with-at-most-k-elements/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n*m*log(m) + total*log(k)) Space: O(total)  
**Kompleksitas Ruang:** O(total)

**Algoritma:** —

**Fungsi Solusi:** `func maxSum(grid [][]int, limits []int, k int) int64`

## Solusi Go

```go
package main

// LeetCode #3462: Maximum Sum With at Most K Elements
// https://leetcode.com/problems/maximum-sum-with-at-most-k-elements/
// Difficulty: Medium
// Time: O(n*m*log(m) + total*log(k)) Space: O(total)

import (
	"fmt"
	"sort"
)

func maxSum(grid [][]int, limits []int, k int) int64 {
	var candidates []int
	for i, row := range grid {
		sort.Ints(row)
		lim := limits[i]
		m := len(row)
		for j := m - lim; j < m; j++ {
			candidates = append(candidates, row[j])
		}
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] > candidates[j]
	})
	var sum int64
	for i := 0; i < k && i < len(candidates); i++ {
		sum += int64(candidates[i])
	}
	return sum
}

func main() {
	fmt.Println(maxSum([][]int{{5, 3, 7}, {8, 2, 6}}, []int{2, 2}, 3)) // 21
	fmt.Println(maxSum([][]int{{1, 2}, {3, 4}}, []int{1, 1}, 2)) // 7
}
```
