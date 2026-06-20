# 0515 — Find Largest Value In Each Tree Row

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func LargestValues(root *TreeNode) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, BFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #515: Find Largest Value in Each Tree Row
// https://leetcode.com/problems/find-largest-value-in-each-tree-row/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test: root = [1,3,2,5,3,null,9]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 3}}
	root.Right = &TreeNode{Val: 2, Right: &TreeNode{Val: 9}}
	fmt.Println(LargestValues(root))
}

func LargestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		maxVal := math.MinInt32
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Val > maxVal {
				maxVal = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, maxVal)
	}

	return result
}
```
