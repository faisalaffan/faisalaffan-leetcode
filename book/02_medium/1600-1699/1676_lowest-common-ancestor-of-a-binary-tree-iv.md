# 1676 — Lowest Common Ancestor Of A Binary Tree Iv

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func lowestCommonAncestor(root *TreeNode, nodes []*TreeNode) *TreeNode`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DFS

**Waktu:** O(n), Space: O(h) where h is height  |  **Ruang:** O(h) where h is height

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1676: Lowest Common Ancestor of a Binary Tree IV
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-iv/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(h) where h is height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, nodes []*TreeNode) *TreeNode {
  // HashMap: O(1) lookup
	nodeSet := make(map[*TreeNode]bool)
	for _, n := range nodes {
		nodeSet[n] = true
	}
	return dfs(root, nodeSet)
}

func dfs(node *TreeNode, nodeSet map[*TreeNode]bool) *TreeNode {
	if node == nil {
		return nil
	}
	if nodeSet[node] {
		return node
	}
	left := dfs(node.Left, nodeSet)
	right := dfs(node.Right, nodeSet)
	if left != nil && right != nil {
		return node
	}
	if left != nil {
		return left
	}
	return right
}

func main() {
	// Test case 1
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	result := lowestCommonAncestor(root, []*TreeNode{root.Left, root.Right})
	fmt.Println("Test 1 (LCA of 5 and 1):", result.Val) // Expected: 3

	// Test case 2: LCA of 6, 7, 4 → 5
	result = lowestCommonAncestor(root, []*TreeNode{root.Left.Left, root.Left.Right.Left, root.Left.Right.Right})
	fmt.Println("Test 2 (LCA of 6, 7, 4):", result.Val) // Expected: 5

	// Test case 3: Single node
	result = lowestCommonAncestor(root, []*TreeNode{root.Left.Right})
	fmt.Println("Test 3 (single node):", result.Val) // Expected: 2
}
```
