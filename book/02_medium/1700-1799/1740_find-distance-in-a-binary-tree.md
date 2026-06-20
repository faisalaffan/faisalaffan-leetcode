# 1740 — Find Distance In A Binary Tree

## Deskripsi

**Soal:** [1740. Find Distance In A Binary Tree](https://leetcode.com/problems/find-distance-in-a-binary-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(h)  
**Kompleksitas Ruang:** O(h)

**Algoritma:** —

**Fungsi Solusi:** `func findDistance(root *TreeNode, p int, q int) int`

## Solusi Go

```go
package main

// LeetCode #1740: Find Distance in a Binary Tree
// https://leetcode.com/problems/find-distance-in-a-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDistance(root *TreeNode, p int, q int) int {
	if p == q {
		return 0
	}
	lca := findLCA(root, p, q)
	return distFrom(lca, p, 0) + distFrom(lca, q, 0)
}

func findLCA(node *TreeNode, p, q int) *TreeNode {
	if node == nil || node.Val == p || node.Val == q {
		return node
	}
	left := findLCA(node.Left, p, q)
	right := findLCA(node.Right, p, q)
	if left != nil && right != nil {
		return node
	}
	if left != nil {
		return left
	}
	return right
}

func distFrom(node *TreeNode, target, dist int) int {
	if node == nil {
		return -1
	}
	if node.Val == target {
		return dist
	}
	if d := distFrom(node.Left, target, dist+1); d != -1 {
		return d
	}
	return distFrom(node.Right, target, dist+1)
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	fmt.Println(findDistance(root, 5, 0)) // Expected: 3
	fmt.Println(findDistance(root, 5, 7)) // Expected: 2
	fmt.Println(findDistance(root, 6, 4)) // Expected: 3
}
```
