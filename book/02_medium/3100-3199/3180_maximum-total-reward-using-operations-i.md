# 3180 — Maximum Total Reward Using Operations I

## Deskripsi

**Soal:** [3180. Maximum Total Reward Using Operations I](https://leetcode.com/problems/maximum-total-reward-using-operations-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * maxVal)  
**Kompleksitas Ruang:** O(maxVal)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func maxTotalReward(rewardValues []int) int`

## Solusi Go

```go
package main

// LeetCode #3180: Maximum Total Reward Using Operations I
// https://leetcode.com/problems/maximum-total-reward-using-operations-i/
// Difficulty: Medium
// Time: O(n * maxVal) | Space: O(maxVal)

import (
	"fmt"
	"sort"
)

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	maxVal := rewardValues[len(rewardValues)-1]
	size := 2 * maxVal
  // Membuat slice untuk menyimpan hasil
	dp := make([]bool, size)
	dp[0] = true

	for _, v := range rewardValues {
		for x := size - 1 - v; x >= 0; x-- {
			if dp[x] && v > x {
				dp[x+v] = true
			}
		}
	}

	for x := size - 1; x >= 0; x-- {
		if dp[x] {
			return x
		}
	}
	return 0
}

func main() {
	fmt.Println(maxTotalReward([]int{1, 1, 3, 3})) // Expected: 4
	fmt.Println(maxTotalReward([]int{1, 6, 4, 3, 2})) // Expected: 11
}
```
