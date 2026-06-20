# 2440 — Create Components With Same Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func componentValue(nums []int, edges [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2440: Create Components With Same Value
// https://leetcode.com/problems/create-components-with-same-value/
// Difficulty: Hard
//
// DFS to compute subtree sums. If the sum of the whole tree is S, and we want
// to split into k components each with value S/k, then we need S to be divisible
// by k and each component's subtree sum must be a multiple of S/k. We try
// all possible divisors of S. For each target = S / k, run DFS that returns
// the cumulative subtree sum, resetting to 0 when target is reached.
// Time O(N * divisors(S)) | Space O(N)

import "fmt"

func main() {
	// Example 1
	fmt.Println(componentValue([]int{6, 2, 2, 2, 6},
		[][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}))
	// Example 2: single node
	fmt.Println(componentValue([]int{2}, [][]int{}))
	// Example 3
	fmt.Println(componentValue([]int{1, 2, 3, 4, 5},
		[][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}))
}

func componentValue(nums []int, edges [][]int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	total := 0
	for _, v := range nums {
		total += v
	}

  // Matriks 2D
	adj := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	best := 1

	// Try divisors of total as possible number of components
	// We try from largest to smallest, so the first valid gives max components
	for k := n; k >= 2; k-- {
		if total%k != 0 {
			continue
		}
		target := total / k
		if canSplit(nums, adj, target) {
			best = k
			break
		}
	}

	return best - 1 // operations = components - 1
}

func canSplit(nums []int, adj [][]int, target int) bool {
	n := len(nums)
	visited := make([]bool, n)
	var dfs func(u int) int
	dfs = func(u int) int {
		visited[u] = true
		sum := nums[u]
		for _, v := range adj[u] {
			if !visited[v] {
				sub := dfs(v)
				if sub == -1 {
					return -1
				}
				sum += sub
			}
		}
		if sum > target {
			return -1
		}
		if sum == target {
			return 0
		}
		return sum
	}
	return dfs(0) == 0
}
```
