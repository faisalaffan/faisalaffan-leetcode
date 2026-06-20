# 0998 — Maximum Binary Tree Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode
```

> **💡 Hint:** Since val is appended to the end of the original array,

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(h) where h is tree height  
**Kompleksitas Ruang:** O(h)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #998: Maximum Binary Tree II
// https://leetcode.com/problems/maximum-binary-tree-ii/
// Difficulty: Medium
//
// Approach: Since val is appended to the end of the original array,
//           if val > root.Val, it becomes the new root (with old root as left child).
//           Otherwise, recurse into the right subtree.
// Time: O(h) where h is tree height
// Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example: root = [4,1,3,null,null,2], val = 5
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{Val: 1, Left: nil, Right: nil},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
			Right: nil,
		},
	}
	result := insertIntoMaxTree(root, 5)
	printTree(result)
	fmt.Println()

	// val < root: root = [5,2,4,null,1], val = 3
	root2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{Val: 1, Left: nil, Right: nil},
		},
		Right: &TreeNode{Val: 4, Left: nil, Right: nil},
	}
	result2 := insertIntoMaxTree(root2, 3)
	printTree(result2)
	fmt.Println()
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val > root.Val {
		return &TreeNode{Val: val, Left: root}
	}
	root.Right = insertIntoMaxTree(root.Right, val)
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
