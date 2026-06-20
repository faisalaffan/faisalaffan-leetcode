# 0971 — Flip Binary Tree To Match Preorder Traversal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func flipMatchVoyage(root *TreeNode, voyage []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #971: Flip Binary Tree To Match Preorder Traversal
// https://leetcode.com/problems/flip-binary-tree-to-match-preorder-traversal/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(n)
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	i := 0
	ok := true
  // Alokasi slice integer
	ans := make([]int, 0)

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || !ok {
			return
		}
		if node.Val != voyage[i] {
			ok = false
			return
		}
		i++
		if node.Left != nil && node.Left.Val != voyage[i] {
			ans = append(ans, node.Val)
			dfs(node.Right)
			dfs(node.Left)
		} else {
			dfs(node.Left)
			dfs(node.Right)
		}
	}

	dfs(root)
	if !ok {
		return []int{-1}
	}
	return ans
}

func main() {
	// [1,2] voyage=[2,1]
	root1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	fmt.Println(flipMatchVoyage(root1, []int{2, 1}))

	// [1,2,3] voyage=[1,3,2]
	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(flipMatchVoyage(root2, []int{1, 3, 2}))

	// [1,2,3] voyage=[1,2,3]
	root3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(flipMatchVoyage(root3, []int{1, 2, 3}))
}
```
