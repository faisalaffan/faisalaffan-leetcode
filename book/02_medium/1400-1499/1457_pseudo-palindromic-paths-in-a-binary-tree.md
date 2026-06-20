# 1457 — Pseudo Palindromic Paths In A Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func pseudoPalindromicPaths(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, DFS, Stack

**Kompleksitas Waktu:** O(n) where n = number of nodes  
**Kompleksitas Ruang:** O(h) for recursion stack

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1457: Pseudo-Palindromic Paths in a Binary Tree
// https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/
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
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{Val: 1},
		},
	}
	fmt.Println(pseudoPalindromicPaths(root)) // 2

	// Test case 2
	root2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}}},
		Right: &TreeNode{Val: 1},
	}
	fmt.Println(pseudoPalindromicPaths(root2)) // 1

	// Test case 3
	fmt.Println(pseudoPalindromicPaths(&TreeNode{Val: 9})) // 1
}

// Time: O(n) where n = number of nodes
// Space: O(h) for recursion stack
func pseudoPalindromicPaths(root *TreeNode) int {
	count := 0
  // Alokasi slice integer
	freq := make([]int, 10) // node values are 1-9

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		freq[node.Val]++
		if node.Left == nil && node.Right == nil {
			// Check if path is pseudo-palindromic
			oddCount := 0
			for _, f := range freq {
				if f%2 == 1 {
					oddCount++
				}
			}
			if oddCount <= 1 {
				count++
			}
		} else {
			dfs(node.Left)
			dfs(node.Right)
		}
		freq[node.Val]--
	}

	dfs(root)
	return count
}
```
