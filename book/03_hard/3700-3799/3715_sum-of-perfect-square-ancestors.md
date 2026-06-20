# 3715 — Sum Of Perfect Square Ancestors

## Deskripsi

**Soal:** [3715. Sum Of Perfect Square Ancestors](https://leetcode.com/problems/sum-of-perfect-square-ancestors/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

> **Ide Kunci:** Compute square-free kernel for each value.

## Solusi Go

```go
package main

// LeetCode #3715: Sum of Perfect Square Ancestors
// https://leetcode.com/problems/sum-of-perfect-square-ancestors/
// Difficulty: Hard
//
// For each node i (1-indexed), count ancestors j such that
// nums[i] * nums[j] is a perfect square. Return sum of counts.
//
// Approach: Compute square-free kernel for each value.
// Product of two values is a perfect square iff their kernels are equal.
// DFS tracking kernel frequencies along current path.

import "fmt"

func main() {
	// Example 1
	fmt.Println(sumOfAncestors(3, [][]int{{0, 1}, {1, 2}}, []int{2, 8, 4}))
	// Example 2
	fmt.Println(sumOfAncestors(4, [][]int{{0, 1}, {0, 2}, {1, 3}}, []int{1, 1, 1, 1}))
	// Edge: single node
	fmt.Println(sumOfAncestors(1, [][]int{}, []int{5}))
}

func sumOfAncestors(n int, edges [][]int, nums []int) int64 {
	if n <= 1 {
		return 0
	}

	// Build tree
  // Membuat slice 2D untuk DP/tabel
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Compute square-free kernel for each value
  // Membuat slice untuk menyimpan hasil
	kernel := make([]int, n)
	for i, v := range nums {
		kernel[i] = squareFree(v)
	}

	var result int64
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)

	var dfs func(u, parent int)
	dfs = func(u, parent int) {
		// Count ancestors with same kernel
		result += int64(freq[kernel[u]])

		freq[kernel[u]]++
		for _, v := range adj[u] {
			if v != parent {
				dfs(v, u)
			}
		}
		freq[kernel[u]]--
	}

	dfs(0, -1)
	return result
}

// squareFree removes all perfect square factors from x
func squareFree(x int) int {
	for p := 2; p*p <= x; p++ {
		for x%(p*p) == 0 {
			x /= p * p
		}
	}
	return x
}
```
