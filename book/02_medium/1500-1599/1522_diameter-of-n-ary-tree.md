# 1522 — Diameter Of N Ary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func Diameter(root *Node) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(N), Space: O(H) where H = height  
**Kompleksitas Ruang:** O(H) where H = height

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
