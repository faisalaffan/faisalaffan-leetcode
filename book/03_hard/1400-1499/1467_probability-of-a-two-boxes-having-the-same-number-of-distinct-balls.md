# 1467 — Probability Of A Two Boxes Having The Same Number Of Distinct Balls

## Deskripsi

**Soal:** [1467. Probability Of A Two Boxes Having The Same Number Of Distinct Balls](https://leetcode.com/problems/probability-of-a-two-boxes-having-the-same-number-of-distinct-balls/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), DFS (Depth-First Search / pencarian kedalaman)

> **Ide Kunci:** DP with combinatorics. Use DFS to distribute balls of each

## Solusi Go

```go
package main

// LeetCode #1467: Probability of a Two Boxes Having The Same Number of Distinct Balls
// https://leetcode.com/problems/probability-of-a-two-boxes-having-the-same-number-of-distinct-balls/
// Difficulty: Hard
//
// Given 2n balls of k distinct colors, each with a given count, randomly
// distribute all balls into two boxes (each gets n balls). Return the
// probability that both boxes have the same number of distinct colors.
//
// Approach: DP with combinatorics. Use DFS to distribute balls of each
// color between the two boxes, counting valid distributions.

import (
	"fmt"
)

func main() {
	// Example 1
	fmt.Println(getProbability([]int{1, 1}))
	// Example 2
	fmt.Println(getProbability([]int{2}))
	// Example 3
	fmt.Println(getProbability([]int{1, 2, 3}))
}

func getProbability(balls []int) float64 {
	n, mx := 0, 0
	for _, x := range balls {
		n += x
		mx = max(mx, x)
	}
	n >>= 1
	m := max(mx, n<<1)
  // Membuat slice 2D untuk DP/tabel
	c := make([][]int, m+1)
  // Iterasi seluruh elemen
	for i := range c {
		c[i] = make([]int, m+1)
	}
	for i := 0; i <= m; i++ {
		c[i][0] = 1
		for j := 1; j <= i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
	k := len(balls)
  // Membuat slice 2D untuk DP/tabel
	f := make([][][]int, k)
  // Iterasi seluruh elemen
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, k<<1|1)
			for h := range f[i][j] {
				f[i][j][h] = -1
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, diff int) int {
		if i >= k {
			if j == 0 && diff == k {
				return 1
			}
			return 0
		}
		if j < 0 {
			return 0
		}
		if f[i][j][diff] != -1 {
			return f[i][j][diff]
		}
		ans := 0
		for x := 0; x <= balls[i]; x++ {
			y := 1
			if x != balls[i] {
				if x == 0 {
					y = -1
				} else {
					y = 0
				}
			}
			ans += dfs(i+1, j-x, diff+y) * c[balls[i]][x]
		}
		f[i][j][diff] = ans
		return ans
	}
	return float64(dfs(0, n, k)) / float64(c[n<<1][n])
}
```
