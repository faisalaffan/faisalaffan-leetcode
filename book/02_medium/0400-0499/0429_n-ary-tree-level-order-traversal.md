# 0429 — N Ary Tree Level Order Traversal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func levelOrder(root *Node) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #429: N-ary Tree Level Order Traversal
// https://leetcode.com/problems/n-ary-tree-level-order-traversal/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		levelSize := len(queue)
  // Alokasi slice integer
		level := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			queue = append(queue, node.Children...)
		}
		result = append(result, level)
	}
	return result
}

func main() {
	// Test case 1: [1,null,3,2,4,null,5,6]
	root1 := &Node{Val: 1}
	root1.Children = []*Node{
		{Val: 3, Children: []*Node{{Val: 5}, {Val: 6}}},
		{Val: 2},
		{Val: 4},
	}
	fmt.Println("Test 1:", levelOrder(root1))
	// Expected: [[1],[3,2,4],[5,6]]

	// Test case 2: Single node
	root2 := &Node{Val: 1}
	fmt.Println("Test 2:", levelOrder(root2))
	// Expected: [[1]]

	// Test case 3: Nil
	fmt.Println("Test 3:", levelOrder(nil))
	// Expected: []
}
```
