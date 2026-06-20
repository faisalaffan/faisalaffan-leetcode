# 1110 — Delete Nodes And Return Forest

## Deskripsi

**Soal:** [1110. Delete Nodes And Return Forest](https://leetcode.com/problems/delete-nodes-and-return-forest/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n + h) where h is tree height

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

> **Ide Kunci:** DFS post-order. If node should be deleted, add children to forest.

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
	deleteSet := make(map[int]bool)
	for _, v := range toDelete {
		deleteSet[v] = true
	}

  // Membuat slice untuk menyimpan hasil
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
