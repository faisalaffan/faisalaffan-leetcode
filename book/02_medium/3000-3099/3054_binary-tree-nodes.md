# 3054 — Binary Tree Nodes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func classifyBinaryTreeNodes(nodes []TreeNode) []NodeClassification`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

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
  // HashMap: O(1) lookup
	nodeSet := make(map[int]bool)
  // HashMap: O(1) lookup
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
  // Custom sort
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
