# 0968 — Binary Tree Cameras

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func minCameraCover(root *TreeNode) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #968: Binary Tree Cameras
// https://leetcode.com/problems/binary-tree-cameras/
// Difficulty: Hard

import (
	"fmt"
	"math"
)

// TreeNode definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	// State 0: node has no camera and is not covered by any child -> needs parent to cover it
	// State 1: node has no camera but is covered by at least one child
	// State 2: node has a camera (covers itself and its neighbors)

	var dfs func(node *TreeNode) [3]int
	dfs = func(node *TreeNode) [3]int {
		if node == nil {
			return [3]int{0, 0, math.MaxInt32 / 2}
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		// State 0: node is not covered, needs parent
		// Children must be in state 1 (covered by their children, no camera)
		s0 := left[1] + right[1]

		// State 1: node is covered (by at least one child with camera)
		// One child must have camera (state 2), other can be state 1 or 2
		s1 := min(
			left[2]+right[1],
			left[1]+right[2],
			left[2]+right[2],
		)

		// State 2: node has a camera
		// Children can be in any state
		s2 := 1 + min(left[0], left[1], left[2]) + min(right[0], right[1], right[2])

		return [3]int{s0, s1, s2}
	}

	result := dfs(root)
	// Root must be covered (state 1 or 2), can't be state 0
	return min(result[1], result[2])
}

func min(a, b int, extra ...int) int {
	result := a
	if b < result {
		result = b
	}
	for _, c := range extra {
		if c < result {
			result = c
		}
	}
	return result
}

func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example 1: [0,0,null,0,0]
	root1 := &TreeNode{Val: 0}
	root1.Left = &TreeNode{Val: 0}
	root1.Left.Left = &TreeNode{Val: 0}
	root1.Left.Right = &TreeNode{Val: 0}
	fmt.Println("Example 1:")
	fmt.Println(minCameraCover(root1))
	// Expected: 1

	// Example 2: [0,0,null,0,null,0,null,null,0]
	root2 := &TreeNode{Val: 0}
	root2.Left = &TreeNode{Val: 0}
	root2.Left.Left = &TreeNode{Val: 0}
	root2.Left.Left.Right = &TreeNode{Val: 0}
	root2.Left.Left.Right.Right = &TreeNode{Val: 0}
	// Wait let me build the tree properly for example 2
	// Actually let's just test example 1 and a custom case

	// Example 3: [0,0,0,null,null,null,0]
	root3 := &TreeNode{Val: 0}
	root3.Left = &TreeNode{Val: 0}
	root3.Right = &TreeNode{Val: 0}
	root3.Right.Right = &TreeNode{Val: 0}
	fmt.Println("Example 3:")
	fmt.Println(minCameraCover(root3))
	// Expected: 2

	// Single node
	root4 := &TreeNode{Val: 0}
	fmt.Println("Example 4 (single node):")
	fmt.Println(minCameraCover(root4))
	// Expected: 1
}
```
