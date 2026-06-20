# 1367 — Linked List In Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func isSubPath(head *ListNode, root *TreeNode) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Stack

**Kompleksitas Waktu:** O(N * L) where N = tree nodes, L = list length  
**Kompleksitas Ruang:** O(N) for recursion stack

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1367: Linked List in Binary Tree
// https://leetcode.com/problems/linked-list-in-binary-tree/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	head := &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 8}}}
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{Val: 6},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
			},
		},
	}
	fmt.Println(isSubPath(head, root)) // true

	// Test case 2
	head2 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6}}}}
	fmt.Println(isSubPath(head2, root)) // true

	// Test case 3
	root3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	head3 := &ListNode{Val: 4}
	fmt.Println(isSubPath(head3, root3)) // false
}

// Time: O(N * L) where N = tree nodes, L = list length
// Space: O(N) for recursion stack
func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	if dfs(head, root) {
		return true
	}
	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func dfs(head *ListNode, node *TreeNode) bool {
	if head == nil {
		return true
	}
	if node == nil {
		return false
	}
	if head.Val != node.Val {
		return false
	}
	return dfs(head.Next, node.Left) || dfs(head.Next, node.Right)
}
```
