# 0103 — Binary Tree Zigzag Level Order Traversal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func zigzagLevelOrder(root *TreeNode) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, BFS

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #103: Binary Tree Zigzag Level Order Traversal
// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	leftToRight := true

	for len(queue) > 0 {
		levelSize := len(queue)
  // Alokasi slice
		level := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			if leftToRight {
				level[i] = node.Val
			} else {
				level[levelSize-1-i] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
		queue = queue[levelSize:]
		leftToRight = !leftToRight
	}

	return result
}

func main() {
	// Test case 1: [3,9,20,null,null,15,7] -> [[3],[20,9],[15,7]]
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(zigzagLevelOrder(root))

	// Test case 2
	fmt.Println(zigzagLevelOrder(nil)) // []

	// Test case 3: [1] -> [[1]]
	fmt.Println(zigzagLevelOrder(&TreeNode{Val: 1})) // [[1]]
}

// Time: O(n) | Space: O(n)
```
