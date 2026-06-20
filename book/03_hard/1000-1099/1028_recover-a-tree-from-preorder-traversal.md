# 1028 — Recover A Tree From Preorder Traversal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func preorderSerialize(root *TreeNode) string
```

> **💡 Hint:** Stack-based parsing.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1028: Recover a Tree From Preorder Traversal
// https://leetcode.com/problems/recover-a-tree-from-preorder-traversal/
// Difficulty: Hard
//
// Approach: Stack-based parsing.
//   Use a stack of nodes. Parse (depth, value) from the traversal string.
//   Pop from stack until stack depth == current depth, then attach as a child.

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// "1-2--3--4-5--6--7"
	root := recoverFromPreorder("1-2--3--4-5--6--7")
	fmt.Println(preorderSerialize(root))

	root2 := recoverFromPreorder("1-2--3---4-5--6---7")
	fmt.Println(preorderSerialize(root2))

	root3 := recoverFromPreorder("3")
	fmt.Println(preorderSerialize(root3))
}

func preorderSerialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	parts := []string{strconv.Itoa(root.Val)}
	if root.Left != nil || root.Right != nil {
		parts = append(parts, preorderSerialize(root.Left))
		parts = append(parts, preorderSerialize(root.Right))
	}
	return strings.Join(parts, ",")
}

func recoverFromPreorder(traversal string) *TreeNode {
	var stack []*TreeNode
	i := 0
	n := len(traversal)

	for i < n {
		// Count dashes to get depth
		depth := 0
		for i < n && traversal[i] == '-' {
			depth++
			i++
		}

		// Parse the number
		start := i
		for i < n && traversal[i] >= '0' && traversal[i] <= '9' {
			i++
		}
		val, _ := strconv.Atoi(traversal[start:i])

		node := &TreeNode{Val: val}

		// Pop stack until depth matches
		for len(stack) > depth {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			parent := stack[len(stack)-1]
			if parent.Left == nil {
				parent.Left = node
			} else {
				parent.Right = node
			}
		}

		stack = append(stack, node)
	}

	if len(stack) == 0 {
		return nil
	}
	return stack[0]
}
```
