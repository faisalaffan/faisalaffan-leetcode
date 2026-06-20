# 2003 — Smallest Missing Genetic Value In Each Subtree

## Deskripsi

**Soal:** [2003. Smallest Missing Genetic Value In Each Subtree](https://leetcode.com/problems/smallest-missing-genetic-value-in-each-subtree/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func smallestMissingValueSubtree(parents []int, nums []int) []int`

> **Ide Kunci:** DFS + set.

## Solusi Go

```go
package main

// LeetCode #2003: Smallest Missing Genetic Value in Each Subtree
// https://leetcode.com/problems/smallest-missing-genetic-value-in-each-subtree/
// Difficulty: Hard
// Approach: DFS + set.
// Only the node with value 1 (and its ancestors) can have missing value > 1.
// Other nodes' answer is always 1.
// Traverse from the 1-node upward, collecting subtree values to find mex.

import "fmt"

func smallestMissingValueSubtree(parents []int, nums []int) []int {
	n := len(parents)
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = 1
	}

	// Find node with value 1
	oneNode := -1
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			oneNode = i
			break
		}
	}
	if oneNode == -1 {
		return ans // all answers are 1
	}

	// Build children adjacency
  // Membuat slice 2D untuk DP/tabel
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parents[i]
		children[p] = append(children[p], i)
	}

	// Reconstruct parent chain from oneNode to root
  // Membuat map untuk pencarian O(1): key → value
	pathSet := make(map[int]bool)
	curr := oneNode
	for curr != -1 {
		pathSet[curr] = true
		curr = parents[curr]
		if curr == -1 {
			break
		}
	}

  // Membuat map untuk pencarian O(1): key → value
	visited := make(map[int]bool)
	mex := 1

	// DFS to collect values in a subtree
	var dfs func(u int)
	dfs = func(u int) {
		visited[nums[u]] = true
		for _, v := range children[u] {
			dfs(v)
		}
	}

	// Process from oneNode upward
	curr = oneNode
	for curr != -1 {
		// DFS all children of curr that are NOT on the path
		for _, v := range children[curr] {
			if !pathSet[v] {
				dfs(v)
			}
		}
		// Add curr's own value
		visited[nums[curr]] = true

		// Update mex
		for visited[mex] {
			mex++
		}
		ans[curr] = mex
		curr = parents[curr]
		if curr == -1 {
			break
		}
	}

	return ans
}

func main() {
	// Example: parents=[-1,0,0,2], nums=[1,2,3,4] -> [5,1,1,1]
	fmt.Println(smallestMissingValueSubtree([]int{-1, 0, 0, 2}, []int{1, 2, 3, 4}))

	// Additional tests
	fmt.Println(smallestMissingValueSubtree([]int{-1, 0, 1, 0}, []int{1, 2, 3, 4}))
	fmt.Println(smallestMissingValueSubtree([]int{-1, 0, 1, 1}, []int{2, 3, 4, 5}))
}
```
