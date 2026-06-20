# 1008 — Construct Binary Search Tree From Preorder Traversal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func bstFromPreorder(preorder []int) *TreeNode
```

> **💡 Hint:** Use upper bound recursion. First element is root. Recursively build

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Binary Search, BFS, Bitmask

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h) where h is tree height

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

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
