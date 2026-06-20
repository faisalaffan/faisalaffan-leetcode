# 0865 — Smallest Subtree With All The Deepest Nodes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func SmallestSubtreeWithAllTheDeepestNodes(root *TreeNode) *TreeNode`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(h)


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
