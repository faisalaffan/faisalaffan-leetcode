# 2581 — Count Number Of Possible Root Nodes

## Deskripsi

**Soal:** [2581. Count Number Of Possible Root Nodes](https://leetcode.com/problems/count-number-of-possible-root-nodes/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func rootCount(edges [][]int, guesses [][]int, k int) int`

## Solusi Go

```go
package main

// LeetCode #2581: Count Number of Possible Root Nodes
// https://leetcode.com/problems/count-number-of-possible-root-nodes/
// Difficulty: Hard

import "fmt"

// rootCount uses rerooting DP. Build tree, count correct guesses with root=0,
// then reroot: moving from u to v, subtract (u,v) if guessed, add (v,u) if guessed.
// Count roots where correct >= k.
//
// Complexity: O(n) time, O(n) space
func rootCount(edges [][]int, guesses [][]int, k int) int {
	n := len(edges) + 1

	// Build adjacency
  // Membuat slice 2D untuk DP/tabel
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Build guess set for O(1) lookup
  // Membuat map untuk pencarian O(1): key → value
	guessSet := make(map[[2]int]bool)
	for _, g := range guesses {
		guessSet[[2]int{g[0], g[1]}] = true
	}

	// First DFS from root 0 to count correct guesses
	correct := 0
	var dfs1 func(u, parent int)
	dfs1 = func(u, parent int) {
		for _, v := range adj[u] {
			if v == parent {
				continue
			}
			if guessSet[[2]int{u, v}] {
				correct++
			}
			dfs1(v, u)
		}
	}
	dfs1(0, -1)

	// Rerooting DFS
	result := 0
	if correct >= k {
		result++
	}

	var dfs2 func(u, parent int, cur int)
	dfs2 = func(u, parent int, cur int) {
		for _, v := range adj[u] {
			if v == parent {
				continue
			}
			next := cur
			// Moving root from u to v: lose (u,v), gain (v,u)
			if guessSet[[2]int{u, v}] {
				next--
			}
			if guessSet[[2]int{v, u}] {
				next++
			}
			if next >= k {
				result++
			}
			dfs2(v, u, next)
		}
	}
	dfs2(0, -1, correct)

	return result
}

func main() {
	// Example from LeetCode
	edges1 := [][]int{{0, 1}, {1, 2}, {1, 3}, {4, 2}}
	guesses1 := [][]int{{1, 3}, {0, 1}, {1, 0}, {2, 4}}
	fmt.Println("Test 1: ->", rootCount(edges1, guesses1, 3)) // 3

	// Additional test cases
	edges2 := [][]int{{0, 1}, {1, 2}}
	guesses2 := [][]int{{0, 1}, {1, 2}}
	fmt.Println("Test 2: ->", rootCount(edges2, guesses2, 2)) // 1

	edges3 := [][]int{{0, 1}, {0, 2}}
	guesses3 := [][]int{{0, 1}, {0, 2}}
	fmt.Println("Test 3: ->", rootCount(edges3, guesses3, 2)) // 1

	// Edge cases
	edges4 := [][]int{{0, 1}}
	guesses4 := [][]int{{0, 1}}
	fmt.Println("Test 4: k=0 ->", rootCount(edges4, guesses4, 0)) // 2

	edges5 := [][]int{{0, 1}, {1, 2}}
	guesses5 := [][]int{{0, 1}}
	fmt.Println("Test 5: k=1 ->", rootCount(edges5, guesses5, 1)) // 2
}
```
