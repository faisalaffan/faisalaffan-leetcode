# 0865 — Smallest Subtree With All The Deepest Nodes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func SmallestSubtreeWithAllTheDeepestNodes(root *TreeNode) *TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #865: Smallest Subtree with all the Deepest Nodes
// https://leetcode.com/problems/smallest-subtree-with-all-the-deepest-nodes/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1: [3,5,1,6,2,0,8,null,null,7,4] -> [2,7,4]
	root := &TreeNode{3,
		&TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}}},
		&TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{8, nil, nil}},
	}
	r1 := SmallestSubtreeWithAllTheDeepestNodes(root)
	fmt.Println(r1.Val)

	// Test case 2: [1] -> [1]
	root2 := &TreeNode{1, nil, nil}
	r2 := SmallestSubtreeWithAllTheDeepestNodes(root2)
	fmt.Println(r2.Val)

	// Test case 3: [0,1,3,null,2] -> [2]
	root3 := &TreeNode{0,
		&TreeNode{1, nil, &TreeNode{2, nil, nil}},
		&TreeNode{3, nil, nil},
	}
	r3 := SmallestSubtreeWithAllTheDeepestNodes(root3)
	fmt.Println(r3.Val)
}

// Time: O(n) | Space: O(h)
func SmallestSubtreeWithAllTheDeepestNodes(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode) (*TreeNode, int)
	dfs = func(node *TreeNode) (*TreeNode, int) {
		if node == nil {
			return nil, 0
		}
		leftNode, leftDepth := dfs(node.Left)
		rightNode, rightDepth := dfs(node.Right)

		if leftDepth > rightDepth {
			return leftNode, leftDepth + 1
		} else if rightDepth > leftDepth {
			return rightNode, rightDepth + 1
		}
		return node, leftDepth + 1
	}

	node, _ := dfs(root)
	return node
}
```
