# 0654 — Maximum Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func constructMaximumBinaryTree(nums []int) *TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n^2) worst case, O(n log n) average  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #654: Maximum Binary Tree
// https://leetcode.com/problems/maximum-binary-tree/
// Difficulty: Medium
// Time: O(n^2) worst case, O(n log n) average
// Space: O(n)

import "fmt"

func main() {
	root := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	fmt.Println(treeToSlice(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return nil
	}

	maxIdx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}

	root := &TreeNode{Val: nums[maxIdx]}
	root.Left = constructMaximumBinaryTree(nums[:maxIdx])
	root.Right = constructMaximumBinaryTree(nums[maxIdx+1:])

	return root
}

func treeToSlice(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{root.Val}
	result = append(result, treeToSlice(root.Left)...)
	result = append(result, treeToSlice(root.Right)...)
	return result
}
```
