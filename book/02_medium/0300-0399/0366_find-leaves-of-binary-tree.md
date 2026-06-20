# 0366 — Find Leaves Of Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func findLeaves(root *TreeNode) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #366: Find Leaves of Binary Tree
// https://leetcode.com/problems/find-leaves-of-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLeaves(root *TreeNode) [][]int {
	result := [][]int{}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		leftHeight := dfs(node.Left)
		rightHeight := dfs(node.Right)
		height := max(leftHeight, rightHeight) + 1

		if height >= len(result) {
			result = append(result, []int{})
		}
		result[height] = append(result[height], node.Val)
		return height
	}
	dfs(root)
	return result
}

func main() {
	// Test case 1: [1,2,3,4,5]
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
	root1.Right = &TreeNode{Val: 3}
	fmt.Println("Test 1:", findLeaves(root1))
	// Expected: [[4,5,3],[2],[1]]

	// Test case 2: Single node
	root2 := &TreeNode{Val: 1}
	fmt.Println("Test 2:", findLeaves(root2))
	// Expected: [[1]]

	// Test case 3: Nil
	fmt.Println("Test 3:", findLeaves(nil))
	// Expected: []
}
```
