# 0671 — Second Minimum Node In A Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func findSecondMinimumValue(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #671: Second Minimum Node In a Binary Tree
// https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/
// Difficulty: Easy

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
	// Example: [2,2,5,null,null,5,7] => 5
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(findSecondMinimumValue(root)) // 5

	root2 := &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(findSecondMinimumValue(root2)) // -1
}

// findSecondMinimumValue finds the second minimum value in a special binary tree
// where each node's value is the minimum of its children.
// Time: O(n). Space: O(n).
func findSecondMinimumValue(root *TreeNode) int {
	result := math.MaxInt64
	dfs(root, root.Val, &result)
	if result == math.MaxInt64 {
		return -1
	}
	return result
}

func dfs(node *TreeNode, rootVal int, second *int) {
	if node == nil {
		return
	}
	if node.Val > rootVal && node.Val < *second {
		*second = node.Val
	}
	dfs(node.Left, rootVal, second)
	dfs(node.Right, rootVal, second)
}
```
