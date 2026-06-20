# 2265 — Count Nodes Equal To Average Of Subtree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func averageOfSubtree(root *TreeNode) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, DFS

**Waktu:** O(n)  |  **Ruang:** O(h)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2265: Count Nodes Equal to Average of Subtree
// https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/
// Difficulty: Medium
// Time: O(n) | Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	count := 0
	dfs(root, &count)
	return count
}

func dfs(node *TreeNode, count *int) (int, int) {
	if node == nil {
		return 0, 0
	}
	leftSum, leftCount := dfs(node.Left, count)
	rightSum, rightCount := dfs(node.Right, count)

	sum := leftSum + rightSum + node.Val
	totalNodes := leftCount + rightCount + 1

	if sum/totalNodes == node.Val {
		*count++
	}
	return sum, totalNodes
}

func main() {
	// Test case 1: [4,8,5,0,1,null,6]
	root1 := &TreeNode{Val: 4}
	root1.Left = &TreeNode{Val: 8}
	root1.Right = &TreeNode{Val: 5}
	root1.Left.Left = &TreeNode{Val: 0}
	root1.Left.Right = &TreeNode{Val: 1}
	root1.Right.Right = &TreeNode{Val: 6}
	fmt.Println(averageOfSubtree(root1))
	// Expected: 5

	// Test case 2: [1]
	root2 := &TreeNode{Val: 1}
	fmt.Println(averageOfSubtree(root2))
	// Expected: 1
}
```
