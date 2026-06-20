# 0671 — Second Minimum Node In A Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func findSecondMinimumValue(root *TreeNode) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #671: Second Minimum Node In a Binary Tree
// https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example: [2,2,5,null,null,5,7] => 5
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(findSecondMinimumValue(root)) // 5

	root2 := &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(findSecondMinimumValue(root2)) // -1
}

// findSecondMinimumValue finds the second minimum value in a special binary tree
// where each node's value is the minimum of its children.
// Time: O(n). Space: O(n).
func findSecondMinimumValue(root *TreeNode) int {
	result := math.MaxInt64
	dfs(root, root.Val, &result)
	if result == math.MaxInt64 {
		return -1
	}
	return result
}

func dfs(node *TreeNode, rootVal int, second *int) {
	if node == nil {
		return
	}
	if node.Val > rootVal && node.Val < *second {
		*second = node.Val
	}
	dfs(node.Left, rootVal, second)
	dfs(node.Right, rootVal, second)
}
```
