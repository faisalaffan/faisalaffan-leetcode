# 0536 — Construct Binary Tree From String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func Str2tree(s string) *TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #536: Construct Binary Tree from String
// https://leetcode.com/problems/construct-binary-tree-from-string/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := Str2tree("4(2(3)(1))(6(5))")
	printTree(root)
	fmt.Println()
}

func Str2tree(s string) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	return buildTree(s, 0)
}

func buildTree(s string, idx int) *TreeNode {
	if idx >= len(s) || s[idx] == ')' {
		return nil
	}

	// Parse number
	start := idx
	for idx < len(s) && (s[idx] == '-' || (s[idx] >= '0' && s[idx] <= '9')) {
		idx++
	}
	val, _ := strconv.Atoi(s[start:idx])
	node := &TreeNode{Val: val}

	// Parse left child
	if idx < len(s) && s[idx] == '(' {
		node.Left = buildTree(s, idx+1)
		// Skip to matching close paren
		depth := 1
		idx++
		for idx < len(s) && depth > 0 {
			if s[idx] == '(' {
				depth++
			} else if s[idx] == ')' {
				depth--
			}
			idx++
		}
	}

	// Parse right child
	if idx < len(s) && s[idx] == '(' {
		node.Right = buildTree(s, idx+1)
		depth := 1
		idx++
		for idx < len(s) && depth > 0 {
			if s[idx] == '(' {
				depth++
			} else if s[idx] == ')' {
				depth--
			}
			idx++
		}
	}

	return node
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printTree(root.Left)
	printTree(root.Right)
}
```
