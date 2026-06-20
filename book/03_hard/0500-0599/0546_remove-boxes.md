# 0546 — Remove Boxes

## Deskripsi

**Soal:** [0546. Remove Boxes](https://leetcode.com/problems/remove-boxes/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #546: Remove Boxes
// https://leetcode.com/problems/remove-boxes/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println(removeBoxes([]int{1, 3, 2, 2, 2, 3, 4, 3, 1})) // Expected: 23
}

func removeBoxes(boxes []int) int {
	n := len(boxes)
	// dp[l][r][k] = max points for boxes[l..r] with k extra boxes of same color as boxes[l] attached to the left
  // Membuat slice 2D untuk DP/tabel
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
		}
	}
	return dfs(boxes, 0, n-1, 0, dp)
}

func dfs(boxes []int, l, r, k int, dp [][][]int) int {
	if l > r {
		return 0
	}
	if dp[l][r][k] > 0 {
		return dp[l][r][k]
	}

	// compress consecutive same colors
	origL := l
	origK := k
	for l+1 <= r && boxes[l+1] == boxes[l] {
		l++
		k++
	}

	// option 1: remove boxes[l..l] plus the k attached ones
	res := (k+1)*(k+1) + dfs(boxes, l+1, r, 0, dp)

	// option 2: merge with later same-color boxes
	for m := l + 1; m <= r; m++ {
		if boxes[m] == boxes[l] {
			res = max(res, dfs(boxes, l+1, m-1, 0, dp)+dfs(boxes, m, r, k+1, dp))
		}
	}

	dp[origL][r][origK] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
