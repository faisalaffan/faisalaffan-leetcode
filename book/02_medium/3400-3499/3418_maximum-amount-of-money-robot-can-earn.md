# 3418 — Maximum Amount Of Money Robot Can Earn

## Deskripsi

**Soal:** [3418. Maximum Amount Of Money Robot Can Earn](https://leetcode.com/problems/maximum-amount-of-money-robot-can-earn/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m*n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func maximumAmount(coins [][]int) int`

## Solusi Go

```go
package main

// LeetCode #3418: Maximum Amount of Money Robot Can Earn
// https://leetcode.com/problems/maximum-amount-of-money-robot-can-earn/
// Difficulty: Medium
// Time: O(m*n) Space: O(n)

import (
	"fmt"
	"math"
)

func maximumAmount(coins [][]int) int {
	n := len(coins[0])
  // Membuat slice untuk menyimpan hasil
	f := make([][3]int, n+1)
	for j := range f {
		f[j] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2}
	}
	f[1] = [3]int{}

	for _, row := range coins {
		for j, x := range row {
			f[j+1][2] = max(f[j][2]+x, f[j+1][2]+x,
				max(f[j][1], f[j+1][1]))
			f[j+1][1] = max(f[j][1]+x, f[j+1][1]+x,
				max(f[j][0], f[j+1][0]))
			f[j+1][0] = max(f[j][0], f[j+1][0]) + x
		}
	}
	return f[n][2]
}

func main() {
	fmt.Println(maximumAmount([][]int{{0, 1, -1}, {1, -2, 3}, {2, -3, 4}})) // 8
	fmt.Println(maximumAmount([][]int{{10, 10, 10}, {10, 10, 10}}))        // 40
}
```
