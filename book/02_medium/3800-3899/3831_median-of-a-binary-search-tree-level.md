# 3831 — Median Of A Binary Search Tree Level

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func MedianOfABinarySearchTreeLevel(root *TreeNode, level int) int
```

> **💡 Hint:** DFS to collect all values at target level, sort, compute upper median.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search, DFS

**Kompleksitas Waktu:** O(N log W)  
**Kompleksitas Ruang:** O(W)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3831: Median of a Binary Search Tree Level
// https://leetcode.com/problems/median-of-a-binary-search-tree-level/
// Difficulty: Medium [Paid]
// Time: O(N log W) | Space: O(W)
// Approach: DFS to collect all values at target level, sort, compute upper median.

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MedianOfABinarySearchTreeLevel(root *TreeNode, level int) int {
	values := []int{}

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == level {
			values = append(values, node.Val)
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	dfs(root, 0)

	if len(values) == 0 {
		return -1
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(values)
	return values[len(values)/2] // upper median
}

func main() {
	// Example 1: root = [4,null,5,null,7], level = 2
	root1 := &TreeNode{4, nil, &TreeNode{5, nil, &TreeNode{7, nil, nil}}}
	fmt.Println(MedianOfABinarySearchTreeLevel(root1, 2)) // Expected: 7

	// Example 2: root = [6,3,8], level = 1
	root2 := &TreeNode{6, &TreeNode{3, nil, nil}, &TreeNode{8, nil, nil}}
	fmt.Println(MedianOfABinarySearchTreeLevel(root2, 1)) // Expected: 8

	// Example 3: root = [2,1], level = 2
	root3 := &TreeNode{2, &TreeNode{1, nil, nil}, nil}
	fmt.Println(MedianOfABinarySearchTreeLevel(root3, 2)) // Expected: -1
}
```
