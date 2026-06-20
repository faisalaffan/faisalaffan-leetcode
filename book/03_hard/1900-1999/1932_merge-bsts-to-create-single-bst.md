# 1932 — Merge Bsts To Create Single Bst

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func canMerge(trees []*TreeNode) *TreeNode`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1932: Merge BSTs to Create Single BST
// https://leetcode.com/problems/merge-bsts-to-create-single-bst/
// Difficulty: Hard
// Map leaf values to tree roots. Merge by replacing leaves with matching root subtrees.
// Finally validate BST on the single remaining tree.

import "fmt"
import "math"

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func canMerge(trees []*TreeNode) *TreeNode {
	// Map root values to nodes
  // HashMap: O(1) lookup
	rootMap := make(map[int]*TreeNode)
	for _, t := range trees {
		rootMap[t.Val] = t
	}

	// Count how many times each value appears as a leaf (incoming count)
  // HashMap: O(1) lookup
	incoming := make(map[int]int)
	for _, t := range trees {
		if t.Left != nil {
			incoming[t.Left.Val]++
		}
		if t.Right != nil {
			incoming[t.Right.Val]++
		}
	}

	// Find the ultimate root (value not appearing as any leaf)
	// Also check that each root that appears as a leaf appears exactly once
	var root *TreeNode
	for _, t := range trees {
		cnt := incoming[t.Val]
		if cnt == 0 {
			if root != nil {
				return nil // more than one tree with no incoming
			}
			root = t
		}
	}
	if root == nil {
		return nil
	}

	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		// If node is a leaf (no children), check if its value is a root of another tree
		if node.Left == nil && node.Right == nil {
			if t, ok := rootMap[node.Val]; ok {
				// Don't merge if the target is the root itself (cycle)
				if t == root {
					return false
				}
				// Merge: replace leaf with the subtree
				node.Left = t.Left
				node.Right = t.Right
				delete(rootMap, node.Val)
				return true
			}
		}
		if !dfs(node.Left) || !dfs(node.Right) {
			return false
		}
		return true
	}

	// Iteratively merge (since merging creates new leaves that might match)
	changed := true
	for changed {
		changed = false
		// Collect leaves to merge to avoid modifying during traversal
		var leaves []*TreeNode
		var collectLeaves func(node *TreeNode)
		collectLeaves = func(node *TreeNode) {
			if node == nil {
				return
			}
			if node.Left == nil && node.Right == nil {
				if _, ok := rootMap[node.Val]; ok && node.Val != root.Val {
					leaves = append(leaves, node)
				}
				return
			}
			collectLeaves(node.Left)
			collectLeaves(node.Right)
		}
		collectLeaves(root)

		for _, leaf := range leaves {
			if t, ok := rootMap[leaf.Val]; ok && t != root {
				leaf.Left = t.Left
				leaf.Right = t.Right
				delete(rootMap, leaf.Val)
				changed = true
			}
		}
	}

	// After merging, only the root tree should remain
	if len(rootMap) != 1 {
		return nil
	}
	if _, ok := rootMap[root.Val]; !ok {
		return nil
	}

	// Validate BST
	var isValid func(node *TreeNode, min, max int) bool
	isValid = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}
		return isValid(node.Left, min, node.Val) && isValid(node.Right, node.Val, max)
	}

	if !isValid(root, math.MinInt32, math.MaxInt32) {
		return nil
	}

	return root
}

func main() {
	// Example 1: trees = [[2,1],[3,2,5],[5,4]] → true
	t1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}
	t2 := &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 5}}
	t3 := &TreeNode{Val: 5, Left: &TreeNode{Val: 4}}
	result := canMerge([]*TreeNode{t1, t2, t3})
	fmt.Println(result != nil) // Expected: true

	// Example 2: trees = [[5,3,8],[3,2,6]] → false (6 is larger than 5)
	t4 := &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 8}}
	t5 := &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 6}}
	result2 := canMerge([]*TreeNode{t4, t5})
	fmt.Println(result2 != nil) // Expected: false

	// Simple case: single tree
	t6 := &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}}
	result3 := canMerge([]*TreeNode{t6})
	fmt.Println(result3 != nil) // Expected: true
}
```
