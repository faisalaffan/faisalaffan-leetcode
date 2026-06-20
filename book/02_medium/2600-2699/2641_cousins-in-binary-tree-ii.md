# 2641 — Cousins In Binary Tree Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func replaceValueInTree(root *TreeNode) *TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2641: Cousins in Binary Tree II
// https://leetcode.com/problems/cousins-in-binary-tree-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	root.Val = 0

	for len(queue) > 0 {
		size := len(queue)
		levelSum := 0

		// First pass: compute total sum of this level (ignore size)
		_ = size
		for _, node := range queue {
			if node.Left != nil {
				levelSum += node.Left.Val
			}
			if node.Right != nil {
				levelSum += node.Right.Val
			}
		}

		// Second pass: update children values
		for _, node := range queue {
			siblingSum := 0
			if node.Left != nil {
				siblingSum += node.Left.Val
			}
			if node.Right != nil {
				siblingSum += node.Right.Val
			}
			if node.Left != nil {
				node.Left.Val = levelSum - siblingSum
			}
			if node.Right != nil {
				node.Right.Val = levelSum - siblingSum
			}
		}

		// Build next level queue
		nextQueue := []*TreeNode{}
		for _, node := range queue {
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		queue = nextQueue
	}
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node.Val, " ")
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	fmt.Println()
}

func main() {
	// Test case 1: [5,4,9,1,10,null,7]
	root1 := &TreeNode{Val: 5}
	root1.Left = &TreeNode{Val: 4}
	root1.Right = &TreeNode{Val: 9}
	root1.Left.Left = &TreeNode{Val: 1}
	root1.Left.Right = &TreeNode{Val: 10}
	root1.Right.Right = &TreeNode{Val: 7}
	replaceValueInTree(root1)
	fmt.Print("Test 1: ")
	printTree(root1)
	// Expected: 0,0,0,7,7,null,11

	// Test case 2: single node
	root2 := &TreeNode{Val: 1}
	replaceValueInTree(root2)
	fmt.Print("Test 2: ")
	printTree(root2)
	// Expected: 0
}
```
