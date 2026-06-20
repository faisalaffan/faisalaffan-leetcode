# 1448 — Count Good Nodes In Binary Tree

## Deskripsi

**Soal:** [1448. Count Good Nodes In Binary Tree](https://leetcode.com/problems/count-good-nodes-in-binary-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) where n = number of nodes  
**Kompleksitas Ruang:** O(h) where h = tree height (recursion stack)

**Algoritma:** Stack (tumpukan LIFO)

## Solusi Go

```go
package main

// LeetCode #1448: Count Good Nodes in Binary Tree
// https://leetcode.com/problems/count-good-nodes-in-binary-tree/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 5}},
	}
	fmt.Println(goodNodes(root)) // 4

	// Test case 2
	root2 := &TreeNode{Val: 3, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 2}}}
	fmt.Println(goodNodes(root2)) // 3

	// Test case 3
	fmt.Println(goodNodes(&TreeNode{Val: 1})) // 1
}

// Time: O(n) where n = number of nodes
// Space: O(h) where h = tree height (recursion stack)
func goodNodes(root *TreeNode) int {
	return countGood(root, root.Val)
}

func countGood(node *TreeNode, maxSoFar int) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Val >= maxSoFar {
		count = 1
		maxSoFar = node.Val
	}

	count += countGood(node.Left, maxSoFar)
	count += countGood(node.Right, maxSoFar)

	return count
}
```
