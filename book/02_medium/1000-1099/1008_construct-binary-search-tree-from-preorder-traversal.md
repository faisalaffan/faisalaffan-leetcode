# 1008 — Construct Binary Search Tree From Preorder Traversal

## Deskripsi

**Soal:** [1008. Construct Binary Search Tree From Preorder Traversal](https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h) where h is tree height

**Algoritma:** Binary Search (pencarian biner)

> **Ide Kunci:** Use upper bound recursion. First element is root. Recursively build

## Solusi Go

```go
package main

// LeetCode #1008: Construct Binary Search Tree from Preorder Traversal
// https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/
// Difficulty: Medium
//
// Approach: Use upper bound recursion. First element is root. Recursively build
//           left subtree with upper bound = root.Val, then right subtree.
// Time: O(n)
// Space: O(h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	result := bstFromPreorder([]int{8, 5, 1, 7, 10, 12})
	printTree(result)
	fmt.Println()

	result2 := bstFromPreorder([]int{1, 3})
	printTree(result2)
	fmt.Println()
}

func bstFromPreorder(preorder []int) *TreeNode {
	idx := 0
	return build(preorder, &idx, 1<<31-1)
}

func build(preorder []int, idx *int, bound int) *TreeNode {
	if *idx >= len(preorder) || preorder[*idx] > bound {
		return nil
	}

	node := &TreeNode{Val: preorder[*idx]}
	*idx++

	node.Left = build(preorder, idx, node.Val)
	node.Right = build(preorder, idx, bound)

	return node
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("[]")
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			fmt.Print("null ")
			continue
		}
		fmt.Printf("%d ", node.Val)
		queue = append(queue, node.Left, node.Right)
	}
}
```
