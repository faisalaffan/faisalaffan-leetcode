# 0333 — Largest Bst Subtree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func largestBSTSubtree(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Bitmask

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h)

> **Untuk fresh graduate:** Kuasai dulu teknik **Bitmask** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #333: Largest BST Subtree
// https://leetcode.com/problems/largest-bst-subtree/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type bstInfo struct {
	min, max int
	size     int
	isBST    bool
}

func largestBSTSubtree(root *TreeNode) int {
	maxSize := 0
	postOrder(root, &maxSize)
	return maxSize
}

func postOrder(node *TreeNode, maxSize *int) bstInfo {
	if node == nil {
		return bstInfo{1<<31 - 1, -1 << 31, 0, true}
	}

	left := postOrder(node.Left, maxSize)
	right := postOrder(node.Right, maxSize)

	info := bstInfo{}
	if left.isBST && right.isBST && node.Val > left.max && node.Val < right.min {
		info.min = min(node.Val, left.min)
		if left.size == 0 {
			info.min = node.Val
		}
		info.max = max(node.Val, right.max)
		if right.size == 0 {
			info.max = node.Val
		}
		info.size = left.size + right.size + 1
		info.isBST = true
		if info.size > *maxSize {
			*maxSize = info.size
		}
	} else {
		info.isBST = false
		info.size = max(left.size, right.size)
	}
	return info
}

func main() {
	// Test case 1: [10,5,15,1,8,null,7]
	root1 := &TreeNode{Val: 10}
	root1.Left = &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 8}}
	root1.Right = &TreeNode{Val: 15, Right: &TreeNode{Val: 7}}
	fmt.Println("Test 1:", largestBSTSubtree(root1))
	// Expected: 3

	// Test case 2: Single node
	root2 := &TreeNode{Val: 1}
	fmt.Println("Test 2:", largestBSTSubtree(root2))
	// Expected: 1

	// Test case 3: Nil
	fmt.Println("Test 3:", largestBSTSubtree(nil))
	// Expected: 0
}
```
