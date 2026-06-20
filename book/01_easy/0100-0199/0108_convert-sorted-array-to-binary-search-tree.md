# 0108 — Convert Sorted Array To Binary Search Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func SortedArrayToBST(nums []int) *TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search, Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(log n) (recursion stack)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #108: Convert Sorted Array to Binary Search Tree
// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(log n) (recursion stack)
func SortedArrayToBST(nums []int) *TreeNode {
	var build func(int, int) *TreeNode
	build = func(lo, hi int) *TreeNode {
		if lo > hi {
			return nil
		}
		mid := lo + (hi-lo)/2
		return &TreeNode{
			Val:   nums[mid],
			Left:  build(lo, mid-1),
			Right: build(mid+1, hi),
		}
	}
	return build(0, len(nums)-1)
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Print(root.Val, " ")
	inorder(root.Right)
}

func main() {
	inorder(SortedArrayToBST([]int{-10, -3, 0, 5, 9}))
	fmt.Println()
	inorder(SortedArrayToBST([]int{1, 3}))
	fmt.Println()
}
```
