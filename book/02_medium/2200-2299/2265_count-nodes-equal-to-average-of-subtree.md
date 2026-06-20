# 2265 — Count Nodes Equal To Average Of Subtree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func averageOfSubtree(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2265: Count Nodes Equal to Average of Subtree
// https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/
// Difficulty: Medium
// Time: O(n) | Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	count := 0
	dfs(root, &count)
	return count
}

func dfs(node *TreeNode, count *int) (int, int) {
	if node == nil {
		return 0, 0
	}
	leftSum, leftCount := dfs(node.Left, count)
	rightSum, rightCount := dfs(node.Right, count)

	sum := leftSum + rightSum + node.Val
	totalNodes := leftCount + rightCount + 1

	if sum/totalNodes == node.Val {
		*count++
	}
	return sum, totalNodes
}

func main() {
	// Test case 1: [4,8,5,0,1,null,6]
	root1 := &TreeNode{Val: 4}
	root1.Left = &TreeNode{Val: 8}
	root1.Right = &TreeNode{Val: 5}
	root1.Left.Left = &TreeNode{Val: 0}
	root1.Left.Right = &TreeNode{Val: 1}
	root1.Right.Right = &TreeNode{Val: 6}
	fmt.Println(averageOfSubtree(root1))
	// Expected: 5

	// Test case 2: [1]
	root2 := &TreeNode{Val: 1}
	fmt.Println(averageOfSubtree(root2))
	// Expected: 1
}
```
