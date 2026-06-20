# 0250 — Count Univalue Subtrees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func countUnivalSubtrees(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n), Space: O(h)  
**Kompleksitas Ruang:** O(h)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #250: Count Univalue Subtrees
// https://leetcode.com/problems/count-univalue-subtrees/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countUnivalSubtrees(root *TreeNode) int {
	count := 0
	isUnival(root, &count)
	return count
}

func isUnival(node *TreeNode, count *int) bool {
	if node == nil {
		return true
	}

	leftUni := isUnival(node.Left, count)
	rightUni := isUnival(node.Right, count)

	if !leftUni || !rightUni {
		return false
	}

	if node.Left != nil && node.Left.Val != node.Val {
		return false
	}
	if node.Right != nil && node.Right.Val != node.Val {
		return false
	}

	*count++
	return true
}

func main() {
	root := &TreeNode{5, &TreeNode{1, &TreeNode{5, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{5, nil, &TreeNode{5, nil, nil}}}
	fmt.Println(countUnivalSubtrees(root))

	root2 := &TreeNode{5, &TreeNode{5, &TreeNode{5, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{5, nil, nil}}
	fmt.Println(countUnivalSubtrees(root2))

	fmt.Println(countUnivalSubtrees(nil))
}
```
