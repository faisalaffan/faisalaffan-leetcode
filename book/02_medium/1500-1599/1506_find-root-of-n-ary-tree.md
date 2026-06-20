# 1506 — Find Root Of N Ary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindRoot(tree []*Node) *Node`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1506: Find Root of N-Ary Tree
// https://leetcode.com/problems/find-root-of-n-ary-tree/
// Difficulty: Medium [Paid]

import "fmt"

// Node is an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

func main() {
	// Build tree: root = 1 -> [2, 3, 4]; 3 -> [5, 6]
	child5 := &Node{Val: 5}
	child6 := &Node{Val: 6}
	child2 := &Node{Val: 2}
	child3 := &Node{Val: 3, Children: []*Node{child5, child6}}
	child4 := &Node{Val: 4}
	root := &Node{Val: 1, Children: []*Node{child2, child3, child4}}

	// All nodes in random order (without knowing root)
	allNodes := []*Node{child2, child5, child3, child4, root, child6}

	found := FindRoot(allNodes)
	fmt.Println("Found root val:", found.Val)
}

func FindRoot(tree []*Node) *Node {
	// Time: O(N), Space: O(1)
	// The root is the only node that is never a child.
	// XOR all node values + all child values. Root value remains.
	var xorSum int
	for _, node := range tree {
		xorSum ^= node.Val
		for _, child := range node.Children {
			xorSum ^= child.Val
		}
	}

	// Find node with matching value
	for _, node := range tree {
		if node.Val == xorSum {
			return node
		}
	}
	return nil
}
```
