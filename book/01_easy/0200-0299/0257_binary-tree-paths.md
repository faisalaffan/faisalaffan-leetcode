# 0257 — Binary Tree Paths

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func BinaryTreePaths(root *TreeNode) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #257: Binary Tree Paths
// https://leetcode.com/problems/binary-tree-paths/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(h)
func BinaryTreePaths(root *TreeNode) []string {
	var res []string
	var dfs func(*TreeNode, string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		path += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			res = append(res, path)
			return
		}
		path += "->"
		dfs(node.Left, path)
		dfs(node.Right, path)
	}
	dfs(root, "")
	return res
}

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	fmt.Println(BinaryTreePaths(root))
	fmt.Println(BinaryTreePaths(nil))
}
```
