# 2458 — Height Of Binary Tree After Subtree Removal Queries

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func treeQueries(root *TreeNode, queries []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, DFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2458: Height of Binary Tree After Subtree Removal Queries
// https://leetcode.com/problems/height-of-binary-tree-after-subtree-removal-queries/
// Difficulty: Hard
//
// Pre-process with two DFS passes:
// 1. Compute height of each node (longest path from node to leaf).
// 2. Compute max height of the tree if we remove each node's subtree, using
//    pre-order traversal. For each node, the alternative is max of:
//    - The answer from parent (tree excluding parent's subtree)
//    - depth + 1 + height of sibling
// Time O(N + Q) | Space O(N)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example 1: root = [1,3,4,2,null,6,5,null,null,null,null,null,7], queries=[4]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 5}
	root.Right.Left.Right = &TreeNode{Val: 7}
	fmt.Println(treeQueries(root, []int{4}))

	// Single node
	root2 := &TreeNode{Val: 1}
	fmt.Println(treeQueries(root2, []int{1}))

	// Two nodes
	root3 := &TreeNode{Val: 1}
	root3.Left = &TreeNode{Val: 2}
	fmt.Println(treeQueries(root3, []int{1, 2}))
}

func treeQueries(root *TreeNode, queries []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	height := make(map[*TreeNode]int)
	computeHeight(root, height)

  // Membuat map (HashMap) — pencarian O(1)
	ans := make(map[int]int)
	dfs(root, 0, 0, height, ans)

  // Alokasi slice integer
	res := make([]int, len(queries))
	for i, q := range queries {
		res[i] = ans[q]
	}
	return res
}

func computeHeight(node *TreeNode, h map[*TreeNode]int) int {
	if node == nil {
		return -1
	}
	left := computeHeight(node.Left, h)
	right := computeHeight(node.Right, h)
	val := 1 + max(left, right)
	h[node] = val
	return val
}

func dfs(node *TreeNode, depth int, parentAns int, h map[*TreeNode]int, ans map[int]int) {
	if node == nil {
		return
	}
	ans[node.Val] = parentAns

	if node.Left != nil {
		alt := parentAns
		if node.Right != nil {
			alt = max(alt, depth+1+h[node.Right])
		}
		dfs(node.Left, depth+1, alt, h, ans)
	}

	if node.Right != nil {
		alt := parentAns
		if node.Left != nil {
			alt = max(alt, depth+1+h[node.Left])
		}
		dfs(node.Right, depth+1, alt, h, ans)
	}
}
```
