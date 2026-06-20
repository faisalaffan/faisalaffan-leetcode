# 0653 — Two Sum Iv Input Is A Bst

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func findTarget(root *TreeNode, k int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #653: Two Sum IV - Input is a BST
// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
// Difficulty: Easy

import "fmt"

// TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example: [5,3,6,2,4,null,7], k=9 => true
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   6,
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(findTarget(root, 9))  // true
	fmt.Println(findTarget(root, 28)) // false
}

// findTarget returns true if there exist two elements in the BST that sum to k.
// Time: O(n). Space: O(n).
func findTarget(root *TreeNode, k int) bool {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]bool)
	return find(root, k, seen)
}

func find(node *TreeNode, k int, seen map[int]bool) bool {
	if node == nil {
		return false
	}
	if seen[k-node.Val] {
		return true
	}
	seen[node.Val] = true
	return find(node.Left, k, seen) || find(node.Right, k, seen)
}
```
