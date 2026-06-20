# 1766 — Tree Of Coprimes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func getCoprimes(nums []int, edges [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS, Stack

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1766: Tree of Coprimes
// https://leetcode.com/problems/tree-of-coprimes/
// Difficulty: Hard
//
// Approach: DFS with depth tracking per value (nums[i] ≤ 50).
// For each node, find the nearest ancestor with a coprime value.
// Since values are ≤ 50, we can precompute all coprime pairs.
// During DFS, maintain for each value (1-50) a stack of (node, depth)
// for ancestors. For each node, check all coprime values and find
// the deepest ancestor.

import (
	"fmt"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func getCoprimes(nums []int, edges [][]int) []int {
	n := len(nums)
	// Build adjacency
  // Matriks 2D
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Precompute coprime pairs for values 1..50
  // Matriks 2D
	coprime := make([][]int, 51)
	for v := 1; v <= 50; v++ {
		for u := 1; u <= 50; u++ {
			if gcd(u, v) == 1 {
				coprime[v] = append(coprime[v], u)
			}
		}
	}

	// For each value (1..50), maintain a stack of (node, depth)
	// Use arrays of [][2]int (idx 0=node, idx 1=depth)
  // Matriks 2D
	valStack := make([][][2]int, 51)

  // Alokasi slice
	ans := make([]int, n)
  // Range loop
	for i := range ans {
		ans[i] = -1
	}

	visited := make([]bool, n)

	var dfs func(u, depth int)
	dfs = func(u, depth int) {
		visited[u] = true
		val := nums[u]

		// Find best ancestor with coprime value
		bestDepth := -1
		bestNode := -1
		for _, cv := range coprime[val] {
			stack := valStack[cv]
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				if top[1] > bestDepth {
					bestDepth = top[1]
					bestNode = top[0]
				}
			}
		}
		ans[u] = bestNode

		// Push current node
		valStack[val] = append(valStack[val], [2]int{u, depth})

		// DFS children
		for _, v := range adj[u] {
			if !visited[v] {
				dfs(v, depth+1)
			}
		}

		// Pop
		valStack[val] = valStack[val][:len(valStack[val])-1]
	}

	dfs(0, 0)
	return ans
}

func main() {
	// Example test case
	nums := []int{2, 3, 3, 2}
	edges := [][]int{{0, 1}, {1, 2}, {1, 3}}
	fmt.Println("nums=[2,3,3,2],edges=[[0,1],[1,2],[1,3]] →", getCoprimes(nums, edges)) // Expected: [-1,0,0,1]

	// Additional tests
	nums2 := []int{5, 6, 10, 2, 3}
	edges2 := [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 4}}
	fmt.Println("nums=[5,6,10,2,3],edges=[[0,1],[0,2],[1,3],[2,4]] →", getCoprimes(nums2, edges2))

	nums3 := []int{1, 1, 1, 1}
	edges3 := [][]int{{0, 1}, {1, 2}, {2, 3}}
	fmt.Println("nums=[1,1,1,1],edges=[[0,1],[1,2],[2,3]] →", getCoprimes(nums3, edges3)) // Expected: [-1,0,1,2]
}
```
