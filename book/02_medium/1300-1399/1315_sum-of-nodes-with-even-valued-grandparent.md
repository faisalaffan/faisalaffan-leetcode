# 1315 — Sum Of Nodes With Even Valued Grandparent

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func sumEvenGrandparent(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Stack

**Kompleksitas Waktu:** O(n) where n is number of nodes  
**Kompleksitas Ruang:** O(h) where h is tree height (recursion stack)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1315: Sum of Nodes with Even-Valued Grandparent
// https://leetcode.com/problems/sum-of-nodes-with-even-valued-grandparent/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 9}},
			Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4}},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}},
		},
	}
	fmt.Println(sumEvenGrandparent(root)) // 18

	// Test case 2: single node
	fmt.Println(sumEvenGrandparent(&TreeNode{Val: 1})) // 0

	// Test case 3
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 5}},
	}
	fmt.Println(sumEvenGrandparent(root2)) // 0 (no even-valued grandparent)
}

// Time: O(n) where n is number of nodes
// Space: O(h) where h is tree height (recursion stack)
func sumEvenGrandparent(root *TreeNode) int {
	return dfs(root, 1, 1) // parent and grandparent start as odd (1)
}

func dfs(node *TreeNode, parent, grandparent int) int {
	if node == nil {
		return 0
	}

	sum := 0
	if grandparent%2 == 0 {
		sum += node.Val
	}

	sum += dfs(node.Left, node.Val, parent)
	sum += dfs(node.Right, node.Val, parent)

	return sum
}
```
