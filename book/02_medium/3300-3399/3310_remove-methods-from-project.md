# 3310 — Remove Methods From Project

## Deskripsi

**Soal:** [3310. Remove Methods From Project](https://leetcode.com/problems/remove-methods-from-project/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + m) Space: O(n + m)  
**Kompleksitas Ruang:** O(n + m)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

## Solusi Go

```go
package main

// LeetCode #3310: Remove Methods From Project
// https://leetcode.com/problems/remove-methods-from-project/
// Difficulty: Medium
// Time: O(n + m) Space: O(n + m)

import "fmt"

func main() {
	fmt.Println(remainingMethods(4, 1, [][]int{{1, 2}, {0, 1}, {2, 3}})) // [0]
	fmt.Println(remainingMethods(5, 0, [][]int{{1, 2}, {0, 2}, {0, 1}, {3, 4}})) // [3 4]
	fmt.Println(remainingMethods(3, 2, [][]int{{0, 1}, {1, 2}, {2, 0}})) // []
}

func remainingMethods(n int, k int, invocations [][]int) []int {
  // Membuat slice 2D untuk DP/tabel
	adj := make([][]int, n)
	for _, inv := range invocations {
		a, b := inv[0], inv[1]
		adj[a] = append(adj[a], b)
	}

	// DFS to find all suspicious methods
  // Membuat slice untuk menyimpan hasil
	suspicious := make([]bool, n)
	var dfs func(u int)
	dfs = func(u int) {
		if suspicious[u] {
			return
		}
		suspicious[u] = true
		for _, v := range adj[u] {
			dfs(v)
		}
	}
	dfs(k)

	// Check if any non-suspicious method calls a suspicious one
	for _, inv := range invocations {
		a, b := inv[0], inv[1]
		if !suspicious[a] && suspicious[b] {
			// Cannot remove - return all methods
  // Membuat slice untuk menyimpan hasil
			res := make([]int, n)
			for i := 0; i < n; i++ {
				res[i] = i
			}
			return res
		}
	}

	// Return non-suspicious methods
	var res []int
	for i := 0; i < n; i++ {
		if !suspicious[i] {
			res = append(res, i)
		}
	}
	return res
}
```
