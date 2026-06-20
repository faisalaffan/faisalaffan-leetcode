# 2925 — Maximum Score After Applying Operations On A Tree

## Deskripsi

**Soal:** [2925. Maximum Score After Applying Operations On A Tree](https://leetcode.com/problems/maximum-score-after-applying-operations-on-a-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2925: Maximum Score After Applying Operations on a Tree
// https://leetcode.com/problems/maximum-score-after-applying-operations-on-a-tree/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(maximumScoreAfterOperations([][]int{{0, 1}, {0, 2}, {0, 3}}, []int{1, 2, 3, 4}))
	fmt.Println(maximumScoreAfterOperations([][]int{{0, 1}}, []int{10, 20}))
}

func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
  // Membuat slice 2D untuk DP/tabel
	g := make([][]int, len(values))
	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	var dfs func(int, int) (int64, int64)
	dfs = func(i, fa int) (int64, int64) {
		a, b := int64(0), int64(0)
		leaf := true
		for _, j := range g[i] {
			if j != fa {
				leaf = false
				aa, bb := dfs(j, i)
				a += aa
				b += bb
			}
		}
		if leaf {
			return int64(values[i]), 0
		}
		va := int64(values[i]) + a
		vb := b + int64(values[i])
		if a > vb {
			vb = a
		}
		return va, vb
	}
	_, b := dfs(0, -1)
	return b
}
```
