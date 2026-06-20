# 3054 — Binary Tree Nodes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func classifyBinaryTreeNodes(nodes []TreeNode) []NodeClassification
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3054: Binary Tree Nodes
// https://leetcode.com/problems/binary-tree-nodes/
// Difficulty: Medium (SQL problem — implemented in Go)
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	N int
	P *int // null for root; using pointer to represent nullable int
}

type NodeClassification struct {
	N    int
	Type string
}

func classifyBinaryTreeNodes(nodes []TreeNode) []NodeClassification {
  // Membuat map (HashMap) — pencarian O(1)
	nodeSet := make(map[int]bool)
  // Membuat map (HashMap) — pencarian O(1)
	parentSet := make(map[int]bool)

	for _, node := range nodes {
		nodeSet[node.N] = true
		if node.P != nil {
			parentSet[*node.P] = true
		}
	}

	var results []NodeClassification
	for _, node := range nodes {
		var nodeType string
		if node.P == nil {
			nodeType = "Root"
		} else if !parentSet[node.N] {
			nodeType = "Leaf"
		} else {
			nodeType = "Inner"
		}
		results = append(results, NodeClassification{N: node.N, Type: nodeType})
	}

	// Order by N ASC
  // Custom sort dengan comparator
	sort.Slice(results, func(i, j int) bool {
		return results[i].N < results[j].N
	})

	return results
}

func intPtr(v int) *int {
	return &v
}

func main() {
	nodes := []TreeNode{
		{N: 1, P: nil},
		{N: 2, P: intPtr(1)},
		{N: 3, P: intPtr(1)},
		{N: 4, P: intPtr(2)},
		{N: 5, P: intPtr(2)},
	}

	fmt.Println("Binary Tree Nodes")
	fmt.Println("================")
	fmt.Printf("%-6s %s\n", "N", "Type")
	fmt.Println("-------------")

	results := classifyBinaryTreeNodes(nodes)
	for _, r := range results {
		fmt.Printf("%-6d %s\n", r.N, r.Type)
	}
}
```
