# 0549 — Binary Tree Longest Consecutive Sequence Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func LongestConsecutive(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h) where h is tree height

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #549: Binary Tree Longest Consecutive Sequence II
// https://leetcode.com/problems/binary-tree-longest-consecutive-sequence-ii/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(LongestConsecutive(root))
}

func LongestConsecutive(root *TreeNode) int {
	maxLen := 0
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		inc, dec := 1, 1

		leftInc, leftDec := dfs(node.Left)
		rightInc, rightDec := dfs(node.Right)

		if node.Left != nil {
			if node.Left.Val == node.Val+1 {
				inc = max(inc, leftInc+1)
			}
			if node.Left.Val == node.Val-1 {
				dec = max(dec, leftDec+1)
			}
		}
		if node.Right != nil {
			if node.Right.Val == node.Val+1 {
				inc = max(inc, rightInc+1)
			}
			if node.Right.Val == node.Val-1 {
				dec = max(dec, rightDec+1)
			}
		}

		maxLen = max(maxLen, inc+dec-1)
		return inc, dec
	}

	dfs(root)
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
