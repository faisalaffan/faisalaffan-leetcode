# 0450 — Delete Node In A Bst

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func deleteNode(root *TreeNode, key int) *TreeNode`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(h)  |  **Ruang:** O(h)


## 💻 Solusi Go

```go
package main

// LeetCode #450: Delete Node in a BST
// https://leetcode.com/problems/delete-node-in-a-bst/
// Difficulty: Medium
// Time: O(h) | Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		// Node to delete found
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// Find inorder successor (leftmost in right subtree)
		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		root.Val = successor.Val
		root.Right = deleteNode(root.Right, successor.Val)
	}
	return root
}

func inorderPrint(node *TreeNode) {
	if node == nil {
		return
	}
	inorderPrint(node.Left)
	fmt.Print(node.Val, " ")
	inorderPrint(node.Right)
}

func main() {
	// Test case 1: [5,3,6,2,4,null,7], key=3
	root1 := &TreeNode{Val: 5}
	root1.Left = &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}
	root1.Right = &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}
	fmt.Print("Test 1 before: ")
	inorderPrint(root1)
	fmt.Println()
	root1 = deleteNode(root1, 3)
	fmt.Print("Test 1 after: ")
	inorderPrint(root1)
	fmt.Println()
	// Expected inorder: 2 4 5 6 7

	// Test case 2: Key not found
	root2 := &TreeNode{Val: 1}
	root2 = deleteNode(root2, 2)
	fmt.Print("Test 2: ")
	inorderPrint(root2)
	fmt.Println()
	// Expected: 1

	// Test case 3: Delete leaf
	root3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}}
	root3 = deleteNode(root3, 2)
	fmt.Print("Test 3: ")
	inorderPrint(root3)
	fmt.Println()
	// Expected: 0 1
}
```
