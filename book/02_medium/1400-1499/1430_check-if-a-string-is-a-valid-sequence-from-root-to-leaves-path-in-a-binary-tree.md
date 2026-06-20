# 1430 — Check If A String Is A Valid Sequence From Root To Leaves Path In A Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func isValidSequence(root *TreeNode, arr []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n) where n = depth of the path  
**Kompleksitas Ruang:** O(h) for recursion

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1430: Check If a String Is a Valid Sequence from Root to Leaves Path in a Binary Tree
// https://leetcode.com/problems/check-if-a-string-is-a-valid-sequence-from-root-to-leaves-path-in-a-binary-tree/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{Val: 0, Right: &TreeNode{Val: 1}},
			Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}},
		},
		Right: &TreeNode{
			Val: 0,
			Left: &TreeNode{Val: 0},
		},
	}
	fmt.Println(isValidSequence(root, []int{0, 1, 0, 1})) // true
	fmt.Println(isValidSequence(root, []int{0, 0, 1}))    // false
	fmt.Println(isValidSequence(root, []int{0, 1, 1}))    // false

	// Test case 2 - empty tree
	fmt.Println(isValidSequence(nil, []int{1})) // false
}

// Time: O(n) where n = depth of the path
// Space: O(h) for recursion
func isValidSequence(root *TreeNode, arr []int) bool {
	return dfs(root, arr, 0)
}

func dfs(node *TreeNode, arr []int, idx int) bool {
	if node == nil || idx >= len(arr) {
		return false
	}
	if node.Val != arr[idx] {
		return false
	}
	if idx == len(arr)-1 {
		return node.Left == nil && node.Right == nil // must be a leaf
	}
	return dfs(node.Left, arr, idx+1) || dfs(node.Right, arr, idx+1)
}
```
