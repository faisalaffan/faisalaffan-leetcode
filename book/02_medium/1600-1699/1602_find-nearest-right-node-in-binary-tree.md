# 1602 — Find Nearest Right Node In Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, BFS

**Kompleksitas Waktu:** O(N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1602: Find Nearest Right Node in Binary Tree
// https://leetcode.com/problems/find-nearest-right-node-in-binary-tree/
// Difficulty: Medium [Paid]

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [1, 2, 3, null, 4, 5, 6]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}
	root.Right = &TreeNode{Val: 3, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}}

	// Find nearest right node of node with value 4
	u := root.Left.Right // node 4
	result := FindNearestRightNode(root, u)
	if result != nil {
		fmt.Println("Nearest right of 4:", result.Val) // should be 5
	} else {
		fmt.Println("Nearest right of 4: nil")
	}

	// Tree: [3, 4, 2, null, null, null, 1]
	root2 := &TreeNode{Val: 3}
	root2.Left = &TreeNode{Val: 4}
	root2.Right = &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}

	result2 := FindNearestRightNode(root2, root2.Left)
	if result2 != nil {
		fmt.Println("Nearest right of 4:", result2.Val)
	} else {
		fmt.Println("Nearest right of 4: nil")
	}
}

func FindNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {
	// Time: O(N), Space: O(N)
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node == u {
				// Return the next node in the queue (right sibling)
				if i+1 < levelSize {
					return queue[0]
				}
				return nil
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return nil
}
```
