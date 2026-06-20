# 0652 — Find Duplicate Subtrees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func findDuplicateSubtrees(root *TreeNode) []*TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #652: Find Duplicate Subtrees
// https://leetcode.com/problems/find-duplicate-subtrees/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
)

func main() {
	// Test case: root = [1,2,3,4,null,2,4,null,null,4]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Left.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 4}

	result := findDuplicateSubtrees(root)
	for _, node := range result {
		fmt.Println(node.Val)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[string]int)
	result := make([]*TreeNode, 0)

	var dfs func(node *TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}

		left := dfs(node.Left)
		right := dfs(node.Right)
		key := strconv.Itoa(node.Val) + "," + left + "," + right

		seen[key]++
		if seen[key] == 2 {
			result = append(result, node)
		}

		return key
	}

	dfs(root)
	return result
}
```
