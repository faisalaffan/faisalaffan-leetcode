# 0872 — Leaf Similar Trees

## Deskripsi

**Soal:** [0872. Leaf Similar Trees](https://leetcode.com/problems/leaf-similar-trees/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n + m). Space: O(n + m).  
**Kompleksitas Ruang:** O(n + m).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #872: Leaf-Similar Trees
// https://leetcode.com/problems/leaf-similar-trees/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// [3,5,1,6,2,9,8,null,null,7,4]
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 8},
		},
	}
	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 8}},
		},
	}
	fmt.Println(leafSimilar(root1, root2)) // true

	root3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	root4 := &TreeNode{Val: 2, Left: &TreeNode{Val: 2}}
	fmt.Println(leafSimilar(root3, root4)) // false
}

// leafSimilar checks if two trees have the same leaf value sequence.
// Time: O(n + m). Space: O(n + m).
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
  // Membuat slice untuk menyimpan hasil
	leaves1 := make([]int, 0)
	collectLeaves(root1, &leaves1)
  // Membuat slice untuk menyimpan hasil
	leaves2 := make([]int, 0)
	collectLeaves(root2, &leaves2)
	if len(leaves1) != len(leaves2) {
		return false
	}
	for i, v := range leaves1 {
		if v != leaves2[i] {
			return false
		}
	}
	return true
}

func collectLeaves(node *TreeNode, leaves *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*leaves = append(*leaves, node.Val)
		return
	}
	collectLeaves(node.Left, leaves)
	collectLeaves(node.Right, leaves)
}
```
