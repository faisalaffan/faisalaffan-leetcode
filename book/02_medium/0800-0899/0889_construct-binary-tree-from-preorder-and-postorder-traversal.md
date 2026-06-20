# 0889 — Construct Binary Tree From Preorder And Postorder Traversal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func preorder(root *TreeNode) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #889: Construct Binary Tree from Preorder and Postorder Traversal
// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1: pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
	r1 := ConstructBinaryTreeFromPreorderAndPostorderTraversal([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1})
	fmt.Println(preorder(r1))

	// Test case 2: pre = [1], post = [1]
	r2 := ConstructBinaryTreeFromPreorderAndPostorderTraversal([]int{1}, []int{1})
	fmt.Println(preorder(r2))

	// Test case 3: pre = [2,1,3], post = [3,1,2]
	r3 := ConstructBinaryTreeFromPreorderAndPostorderTraversal([]int{2, 1, 3}, []int{3, 1, 2})
	fmt.Println(preorder(r3))
}

func preorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{root.Val}
	res = append(res, preorder(root.Left)...)
	res = append(res, preorder(root.Right)...)
	return res
}

// Time: O(n) | Space: O(n)
func ConstructBinaryTreeFromPreorderAndPostorderTraversal(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}

	// The second element in preorder is the root of the left subtree
	// Find it in postorder to determine left subtree size
	leftSize := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(postorder); i++ {
		if postorder[i] == preorder[1] {
			leftSize = i + 1
			break
		}
	}

	root.Left = ConstructBinaryTreeFromPreorderAndPostorderTraversal(preorder[1:1+leftSize], postorder[:leftSize])
	root.Right = ConstructBinaryTreeFromPreorderAndPostorderTraversal(preorder[1+leftSize:], postorder[leftSize:len(postorder)-1])

	return root
}
```
