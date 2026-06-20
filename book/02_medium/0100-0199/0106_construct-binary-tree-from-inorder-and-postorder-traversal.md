# 0106 — Construct Binary Tree From Inorder And Postorder Traversal

## Deskripsi

**Soal:** [0106. Construct Binary Tree From Inorder And Postorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func buildTree(inorder []int, postorder []int) *TreeNode`

## Solusi Go

```go
package main

// LeetCode #106: Construct Binary Tree from Inorder and Postorder Traversal
// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

  // Membuat map untuk pencarian O(1): key → value
	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}

	var build func(inStart, inEnd, postStart, postEnd int) *TreeNode
	build = func(inStart, inEnd, postStart, postEnd int) *TreeNode {
		if inStart > inEnd || postStart > postEnd {
			return nil
		}

		rootVal := postorder[postEnd]
		root := &TreeNode{Val: rootVal}
		inIdx := inorderMap[rootVal]
		rightSize := inEnd - inIdx

		root.Left = build(inStart, inIdx-1, postStart, postEnd-rightSize-1)
		root.Right = build(inIdx+1, inEnd, postEnd-rightSize, postEnd-1)
		return root
	}

	return build(0, len(inorder)-1, 0, len(postorder)-1)
}

func printInorder(root *TreeNode) {
	if root == nil {
		return
	}
	printInorder(root.Left)
	fmt.Printf("%d ", root.Val)
	printInorder(root.Right)
}

func main() {
	// Test case 1
	root := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	printInorder(root) // 9 3 15 20 7
	fmt.Println()

	// Test case 2
	root = buildTree([]int{-1}, []int{-1})
	printInorder(root) // -1
	fmt.Println()
}

// Time: O(n) | Space: O(n)
```
