# 1038 — Binary Search Tree To Greater Sum Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func bstToGst(root *TreeNode) *TreeNode
```

> **💡 Hint:** Reverse inorder (right -> root -> left) accumulating sum

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search, DFS, BFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h) where h is tree height

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1038: Binary Search Tree to Greater Sum Tree
// https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/
// Difficulty: Medium
//
// Approach: Reverse inorder (right -> root -> left) accumulating sum
// Time: O(n)
// Space: O(h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0, Left: nil, Right: nil},
			Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
		},
		Right: &TreeNode{
			Val:   6,
			Left:  &TreeNode{Val: 5, Left: nil, Right: nil},
			Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 8, Left: nil, Right: nil}},
		},
	}
	result := bstToGst(root)
	printTree(result)
	fmt.Println()
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}
	dfs(root)
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			fmt.Print("null ")
			continue
		}
		fmt.Printf("%d ", node.Val)
		queue = append(queue, node.Left, node.Right)
	}
}
```
