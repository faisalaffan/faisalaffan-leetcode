# 2097 — Valid Arrangement Of Pairs

## Deskripsi

**Soal:** [2097. Valid Arrangement Of Pairs](https://leetcode.com/problems/valid-arrangement-of-pairs/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman), LIS (Longest Increasing Subsequence)

## Solusi Go

```go
package main

// LeetCode #2097: Valid Arrangement of Pairs
// https://leetcode.com/problems/valid-arrangement-of-pairs/
// Difficulty: Hard
//
// Eulerian Path (Hierholzer's Algorithm). Build adjacency list, compute
// in/out degrees. Start at node with out > in, or any node with edges.
// Use iterative DFS to reconstruct the path, then reverse.

import "fmt"

func main() {
	fmt.Println(validArrangement([][]int{{5, 1}, {4, 5}, {11, 9}, {9, 4}}))
	fmt.Println(validArrangement([][]int{{1, 2}, {2, 3}, {3, 4}}))
	fmt.Println(validArrangement([][]int{{1, 2}, {2, 1}}))
}

func validArrangement(pairs [][]int) [][]int {
  // Membuat map untuk pencarian O(1): key → value
	adj := make(map[int][]int)
  // Membuat map untuk pencarian O(1): key → value
	inDeg := make(map[int]int)
  // Membuat map untuk pencarian O(1): key → value
	outDeg := make(map[int]int)

	for _, p := range pairs {
		u, v := p[0], p[1]
		adj[u] = append(adj[u], v)
		outDeg[u]++
		inDeg[v]++
	}

	// Find start node: node with out > in, or any node with edges
	start := pairs[0][0]
	for node := range adj {
		if outDeg[node] > inDeg[node] {
			start = node
			break
		}
	}

	// Hierholzer's algorithm (iterative)
	stack := []int{start}
	path := []int{}

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		if len(adj[cur]) > 0 {
			next := adj[cur][0]
			adj[cur] = adj[cur][1:]
			stack = append(stack, next)
		} else {
			path = append(path, cur)
			stack = stack[:len(stack)-1]
		}
	}

	// Reverse path to get correct order
  // Membuat slice 2D untuk DP/tabel
	result := make([][]int, len(path)-1)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(path)-1; i++ {
		result[i] = []int{path[len(path)-1-i], path[len(path)-2-i]}
	}
	return result
}
```
