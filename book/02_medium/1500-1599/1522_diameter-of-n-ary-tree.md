# 1522 — Diameter Of N Ary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func Diameter(root *Node) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS

**Waktu:** O(N), Space: O(H) where H = height  |  **Ruang:** O(H) where H = height

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1522: Diameter of N-Ary Tree
// https://leetcode.com/problems/diameter-of-n-ary-tree/
// Difficulty: Medium [Paid]

import "fmt"

// Node is an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

func main() {
	// Tree: 1 -> [2, 3, 4]; 3 -> [5, 6]
	child5 := &Node{Val: 5}
	child6 := &Node{Val: 6}
	child2 := &Node{Val: 2}
	child3 := &Node{Val: 3, Children: []*Node{child5, child6}}
	child4 := &Node{Val: 4}
	root := &Node{Val: 1, Children: []*Node{child2, child3, child4}}
	fmt.Println(Diameter(root))

	// Single node
	fmt.Println(Diameter(&Node{Val: 1}))

	// Linear chain: 1 -> 2 -> 3
	n3 := &Node{Val: 3}
	n2 := &Node{Val: 2, Children: []*Node{n3}}
	n1 := &Node{Val: 1, Children: []*Node{n2}}
	fmt.Println(Diameter(n1))
}

func Diameter(root *Node) int {
	// Time: O(N), Space: O(H) where H = height
	maxDiameter := 0

	var dfs func(node *Node) int
	dfs = func(node *Node) int {
		if node == nil {
			return 0
		}
		// Track top two deepest paths from children
		firstMax, secondMax := 0, 0

		for _, child := range node.Children {
			depth := dfs(child)
			if depth > firstMax {
				secondMax = firstMax
				firstMax = depth
			} else if depth > secondMax {
				secondMax = depth
			}
		}

		// Diameter through this node = sum of two deepest child paths
		if firstMax+secondMax > maxDiameter {
			maxDiameter = firstMax + secondMax
		}

		// Return max depth from this node
		return firstMax + 1
	}

	dfs(root)
	return maxDiameter
}
```
