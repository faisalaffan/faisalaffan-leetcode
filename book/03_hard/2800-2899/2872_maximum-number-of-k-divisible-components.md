# 2872 — Maximum Number Of K Divisible Components

## Deskripsi

**Soal:** [2872. Maximum Number Of K Divisible Components](https://leetcode.com/problems/maximum-number-of-k-divisible-components/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman), LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int`

## Solusi Go

```go
package main

// LeetCode #2872: Maximum Number of K-Divisible Components
// https://leetcode.com/problems/maximum-number-of-k-divisible-components/
// Difficulty: Hard
//
// DFS on tree. Post-order traversal computes subtree sum modulo k. Whenever
// subtree sum % k == 0, we can cut the edge to the parent (increment count).
// O(N) time, O(N) space.

import "fmt"

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	// Build adjacency list
  // Membuat slice 2D untuk DP/tabel
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	count := 0

	var dfs func(node int, parent int) int
	dfs = func(node int, parent int) int {
		sum := values[node] % k
		for _, neighbor := range adj[node] {
			if neighbor == parent {
				continue
			}
			childSum := dfs(neighbor, node)
			if childSum == 0 {
				count++
			} else {
				sum = (sum + childSum) % k
			}
		}
		return sum
	}

	rootSum := dfs(0, -1)
	if rootSum == 0 {
		count++
	}

	return count
}

func main() {
	// Example: n=5, edges=[[0,2],[1,2],[1,3],[2,4]], values=[1,8,1,4,4], k=6 => 2
	fmt.Println(maxKDivisibleComponents(5,
		[][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}},
		[]int{1, 8, 1, 4, 4}, 6))
	// Single node
	fmt.Println(maxKDivisibleComponents(1, [][]int{}, []int{5}, 5))
	fmt.Println(maxKDivisibleComponents(1, [][]int{}, []int{3}, 5))
	// Two nodes
	fmt.Println(maxKDivisibleComponents(2,
		[][]int{{0, 1}},
		[]int{2, 4}, 6))
	fmt.Println(maxKDivisibleComponents(2,
		[][]int{{0, 1}},
		[]int{1, 1}, 2))
	// Linear chain
	fmt.Println(maxKDivisibleComponents(3,
		[][]int{{0, 1}, {1, 2}},
		[]int{1, 2, 3}, 3))
	// All divisible individually
	fmt.Println(maxKDivisibleComponents(3,
		[][]int{{0, 1}, {1, 2}},
		[]int{3, 6, 9}, 3))
	// All values divisible by k
	fmt.Println(maxKDivisibleComponents(4,
		[][]int{{0, 1}, {1, 2}, {2, 3}},
		[]int{6, 6, 6, 6}, 6))
	// Star: center value not divisible by k, leaves are
	fmt.Println(maxKDivisibleComponents(4,
		[][]int{{0, 1}, {0, 2}, {0, 3}},
		[]int{1, 6, 6, 6}, 6))
}
```
