# 0617 — Merge Two Binary Trees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func MergeTwoBinaryTrees(root1, root2 *TreeNode) *TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n+m), Space: O(h)  
**Kompleksitas Ruang:** O(h)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #617: Merge Two Binary Trees
// https://leetcode.com/problems/merge-two-binary-trees/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n+m), Space: O(h)
func MergeTwoBinaryTrees(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	return &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  MergeTwoBinaryTrees(root1.Left, root2.Left),
		Right: MergeTwoBinaryTrees(root1.Right, root2.Right),
	}
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("null ")
		return
	}
	fmt.Printf("%d ", root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	// Test: root1=[1,3,2,5], root2=[2,1,3,null,4,null,7]
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 2},
	}
	root2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 7},
		},
	}
	printTree(MergeTwoBinaryTrees(root1, root2))
	fmt.Println()

	// Test: root1=[1], root2=[1,2]
	r1 := &TreeNode{Val: 1}
	r2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	printTree(MergeTwoBinaryTrees(r1, r2))
	fmt.Println()
}
```
