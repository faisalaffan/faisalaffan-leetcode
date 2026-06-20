# 0256 — Paint House

## Deskripsi

**Soal:** [0256. Paint House](https://leetcode.com/problems/paint-house/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func minCost(costs [][]int) int`

## Solusi Go

```go
package main

// LeetCode #256: Paint House
// https://leetcode.com/problems/paint-house/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}

	for i := 1; i < len(costs); i++ {
		costs[i][0] += min(costs[i-1][1], costs[i-1][2])
		costs[i][1] += min(costs[i-1][0], costs[i-1][2])
		costs[i][2] += min(costs[i-1][0], costs[i-1][1])
	}

	last := costs[len(costs)-1]
	return min(last[0], min(last[1], last[2]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minCost([][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}))
	fmt.Println(minCost([][]int{{7, 6, 2}}))
	fmt.Println(minCost([][]int{}))
}
```
