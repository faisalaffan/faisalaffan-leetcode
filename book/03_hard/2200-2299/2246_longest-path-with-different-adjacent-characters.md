# 2246 — Longest Path With Different Adjacent Characters

## Deskripsi

**Soal:** [2246. Longest Path With Different Adjacent Characters](https://leetcode.com/problems/longest-path-with-different-adjacent-characters/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman), LIS (Longest Increasing Subsequence)

## Solusi Go

```go
package main

// LeetCode #2246: Longest Path with Different Adjacent Characters
// https://leetcode.com/problems/longest-path-with-different-adjacent-characters/
// Difficulty: Hard
//
// Tree DFS: Build adjacency list. DFS from root. For each node, collect the
// longest path lengths from children that have a different character. The
// longest path through this node = sum of top two + 1. Update global max.

import (
	"fmt"
)

func main() {
	// parent = [-1,0,0,1,1,2], s = "abacbe" => 3
	fmt.Println(longestPath([]int{-1, 0, 0, 1, 1, 2}, "abacbe"))
	// single node
	fmt.Println(longestPath([]int{-1}, "a"))
	// all same chars (path should be 1 since no two adjacent can be same char)
	fmt.Println(longestPath([]int{-1, 0, 0}, "aaa"))
	// linear
	fmt.Println(longestPath([]int{-1, 0, 0, 1, 1, 2, 5}, "abacbea"))
}

func longestPath(parent []int, s string) int {
	n := len(parent)
  // Membuat slice 2D untuk DP/tabel
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		children[p] = append(children[p], i)
	}

	ans := 1

	var dfs func(u int) int
	dfs = func(u int) int {
		top1, top2 := 0, 0
		for _, v := range children[u] {
			childLen := dfs(v)
			if s[v] != s[u] {
				if childLen > top1 {
					top2 = top1
					top1 = childLen
				} else if childLen > top2 {
					top2 = childLen
				}
			}
		}
		// Path through u = top1 + top2 + 1
		if top1+top2+1 > ans {
			ans = top1 + top2 + 1
		}
		return top1 + 1
	}

	dfs(0)
	return ans
}
```
