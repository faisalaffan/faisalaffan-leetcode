# 1644 — Lowest Common Ancestor Of A Binary Tree Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func LowestCommonAncestorII(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N), Space: O(N)  |  **Ruang:** O(N)


## 💻 Solusi Go

```go
package main

// LeetCode #1644: Lowest Common Ancestor of a Binary Tree II
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-ii/
// Difficulty: Medium [Paid]

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [3,5,1,6,2,0,8,null,null,7,4]
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}}
	root.Right = &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}

	// LCA of 5 and 1 is 3
	p := root.Left // 5
	q := root.Right // 1
	lca := LowestCommonAncestorII(root, p, q)
	if lca != nil {
		fmt.Println("LCA of 5 and 1:", lca.Val) // 3
	} else {
		fmt.Println("LCA: nil")
	}

	// LCA of 5 and 4 is 5
	p = root.Left // 5
	q = root.Left.Right.Right // 4
	lca = LowestCommonAncestorII(root, p, q)
	if lca != nil {
		fmt.Println("LCA of 5 and 4:", lca.Val) // 5
	}

	// Node not in tree
	notInTree := &TreeNode{Val: 10}
	lca = LowestCommonAncestorII(root, root.Left, notInTree)
	if lca == nil {
		fmt.Println("LCA of 5 and 10: nil (10 not in tree)")
	}
}

// LowestCommonAncestorII differs from LCA I in that p and q may not exist in the tree.
func LowestCommonAncestorII(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	// Time: O(N), Space: O(N)
	foundP := false
	foundQ := false

	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		if node == p {
			foundP = true
			return node
		}
		if node == q {
			foundQ = true
			return node
		}

		if left != nil && right != nil {
			return node
		}
		if left != nil {
			return left
		}
		return right
	}

	lca := dfs(root)

	if foundP && foundQ {
		return lca
	}
	return nil
}
```
