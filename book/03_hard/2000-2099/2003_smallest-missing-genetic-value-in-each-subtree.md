# 2003 — Smallest Missing Genetic Value In Each Subtree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func smallestMissingValueSubtree(parents []int, nums []int) []int
```

> **💡 Hint:** DFS + set.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, DFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

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
  // Alokasi slice integer
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
  // Membuat matriks/slice 2D untuk DP
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parents[i]
		children[p] = append(children[p], i)
	}

	// Reconstruct parent chain from oneNode to root
  // Membuat map (HashMap) — pencarian O(1)
	pathSet := make(map[int]bool)
	curr := oneNode
	for curr != -1 {
		pathSet[curr] = true
		curr = parents[curr]
		if curr == -1 {
			break
		}
	}

  // Membuat map (HashMap) — pencarian O(1)
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
