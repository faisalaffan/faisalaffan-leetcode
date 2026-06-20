# 1022 — Sum Of Root To Leaf Binary Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func sumRootToLeaf(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h) where h is tree height

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1022: Sum of Root To Leaf Binary Numbers
// https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/
// Difficulty: Easy
// Time: O(n) | Space: O(h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//     1
	//    / \
	//   0   1
	//  / \ / \
	// 0  1 0  1
	root := &TreeNode{1,
		&TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}},
		&TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}},
	}
	fmt.Println(sumRootToLeaf(root)) // 22

	// Single node
	fmt.Println(sumRootToLeaf(&TreeNode{1, nil, nil})) // 1
}

// LeetCode submission: sumRootToLeaf
func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, val int) int {
	if node == nil {
		return 0
	}
	val = val*2 + node.Val
	if node.Left == nil && node.Right == nil {
		return val
	}
	return dfs(node.Left, val) + dfs(node.Right, val)
}
```
