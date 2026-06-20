# 3319 — K Th Largest Perfect Subtree Size In Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func kthLargestPerfectSubtree(root *TreeNode, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3319: K-th Largest Perfect Subtree Size in Binary Tree
// https://leetcode.com/problems/k-th-largest-perfect-subtree-size-in-binary-tree/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	// Tree: [1,2,3,4,5,6,7]
	root := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7}}}
	fmt.Println(kthLargestPerfectSubtree(root, 1)) // 7
	fmt.Println(kthLargestPerfectSubtree(root, 3)) // 3
	fmt.Println(kthLargestPerfectSubtree(root, 5)) // -1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestPerfectSubtree(root *TreeNode, k int) int {
	var sizes []int
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left < 0 || left != right {
			return -1
		}
		cur := left + right + 1
		sizes = append(sizes, cur)
		return cur
	}
	dfs(root)

	if len(sizes) < k {
		return -1
	}
  // Custom sort dengan comparator
	sort.Slice(sizes, func(i, j int) bool { return sizes[i] > sizes[j] })
	return sizes[k-1]
}
```
