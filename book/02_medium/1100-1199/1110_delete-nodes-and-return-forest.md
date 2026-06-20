# 1110 — Delete Nodes And Return Forest

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func delNodes(root *TreeNode, toDelete []int) []*TreeNode`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DFS

**Waktu:** O(n)  |  **Ruang:** O(n + h) where h is tree height

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1110: Delete Nodes And Return Forest
// https://leetcode.com/problems/delete-nodes-and-return-forest/
// Difficulty: Medium
//
// Approach: DFS post-order. If node should be deleted, add children to forest.
// Time: O(n)
// Space: O(n + h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
			Right: &TreeNode{Val: 5, Left: nil, Right: nil},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
			Right: &TreeNode{Val: 7, Left: nil, Right: nil},
		},
	}
	result := delNodes(root, []int{3, 5})
	for _, tree := range result {
		if tree != nil {
			fmt.Printf("%d ", tree.Val)
		}
	}
	fmt.Println()
}

func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
  // HashMap: O(1) lookup
	deleteSet := make(map[int]bool)
	for _, v := range toDelete {
		deleteSet[v] = true
	}

	result := make([]*TreeNode, 0)
	if !deleteSet[root.Val] {
		result = append(result, root)
	}

	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)

		if deleteSet[node.Val] {
			if node.Left != nil {
				result = append(result, node.Left)
			}
			if node.Right != nil {
				result = append(result, node.Right)
			}
			return nil
		}
		return node
	}

	dfs(root)
	return result
}
```
