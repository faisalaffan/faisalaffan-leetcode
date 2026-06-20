# 0199 — Binary Tree Right Side View

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func rightSideView(root *TreeNode) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, BFS

**Kompleksitas Waktu:** O(n), Space: O(h) for queue  
**Kompleksitas Ruang:** O(h) for queue

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #199: Binary Tree Right Side View
// https://leetcode.com/problems/binary-tree-right-side-view/
// Difficulty: Medium
// Time: O(n), Space: O(h) for queue

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if i == levelSize-1 {
				result = append(result, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, &TreeNode{4, nil, nil}}}
	fmt.Println(rightSideView(root))

	root2 := &TreeNode{1, nil, &TreeNode{3, nil, nil}}
	fmt.Println(rightSideView(root2))

	fmt.Println(rightSideView(nil))
}
```
