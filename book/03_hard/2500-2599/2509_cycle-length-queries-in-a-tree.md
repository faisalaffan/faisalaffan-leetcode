# 2509 — Cycle Length Queries In A Tree

## Deskripsi

**Soal:** [2509. Cycle Length Queries In A Tree](https://leetcode.com/problems/cycle-length-queries-in-a-tree/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2509: Cycle Length Queries in a Tree
// https://leetcode.com/problems/cycle-length-queries-in-a-tree/
// Difficulty: Hard
//
// Complete binary tree with n levels. For each query (a,b), find the length
// of the cycle formed by adding an edge between a and b.
// Cycle length = depth(a) + depth(b) - 2*depth(LCA(a,b)) + 1.
// Node numbering: root=1, left=2*i, right=2*i+1.

import "fmt"

func main() {
	// Example 1: n=3, queries=[[5,3],[4,7],[2,3]] => [4,5,3]
	fmt.Println(cycleLengthQueries(3, [][]int{{5, 3}, {4, 7}, {2, 3}}))
	// Example 2: n=2, queries=[[1,2]] => [2]
	fmt.Println(cycleLengthQueries(2, [][]int{{1, 2}}))
	// Edge: same node
	fmt.Println(cycleLengthQueries(2, [][]int{{1, 1}}))
	// Edge: root with another
	fmt.Println(cycleLengthQueries(3, [][]int{{1, 7}}))
}

func cycleLengthQueries(n int, queries [][]int) []int {
  // Membuat slice untuk menyimpan hasil
	result := make([]int, len(queries))
	for i, q := range queries {
		a, b := q[0], q[1]
		if a == b {
			result[i] = 1 // cycle of a self-loop
			continue
		}
		// Find depths and LCA
		da := depth(a)
		db := depth(b)
		lca := findLCA(a, b)
		dlca := depth(lca)
		// Cycle length = da + db - 2*dlca + 1
		result[i] = da + db - 2*dlca + 1
	}
	return result
}

func depth(x int) int {
	d := 0
	for x > 0 {
		x >>= 1
		d++
	}
	return d
}

func findLCA(a, b int) int {
	// Bring nodes to same depth
	for depth(a) > depth(b) {
		a >>= 1
	}
	for depth(b) > depth(a) {
		b >>= 1
	}
	// Move up together
	for a != b {
		a >>= 1
		b >>= 1
	}
	return a
}
```
