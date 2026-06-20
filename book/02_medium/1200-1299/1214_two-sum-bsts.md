# 1214 — Two Sum Bsts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n + m)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1214: Two Sum BSTs
// https://leetcode.com/problems/two-sum-bsts/
// Difficulty: Medium [Paid]

// Given two BSTs and a target, return true if there exists
// a node from each tree whose values sum to target.

// Time: O(n + m)
// Space: O(n + m)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
  // Membuat map (HashMap) — pencarian O(1)
	vals := make(map[int]bool)

	var collect func(node *TreeNode)
	collect = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals[node.Val] = true
		collect(node.Left)
		collect(node.Right)
	}
	collect(root1)

	var find func(node *TreeNode) bool
	find = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if vals[target-node.Val] {
			return true
		}
		return find(node.Left) || find(node.Right)
	}
	return find(root2)
}

func main() {
	// Tree1: [2,1,4], Tree2: [1,0,3], target=5
	r1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4}}
	r2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 3}}
	fmt.Printf("%t (expected: true)\n", twoSumBSTs(r1, r2, 5))

	// target=10 -> false
	fmt.Printf("%t (expected: false)\n", twoSumBSTs(r1, r2, 10))
}
```
