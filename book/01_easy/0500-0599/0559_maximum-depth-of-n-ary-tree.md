# 0559 — Maximum Depth Of N Ary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MaximumDepthOfNAryTree(root *Node) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(h)  |  **Ruang:** O(h)


## 💻 Solusi Go

```go
package main

// LeetCode #559: Maximum Depth of N-ary Tree
// https://leetcode.com/problems/maximum-depth-of-n-ary-tree/
// Difficulty: Easy

import "fmt"

// Node represents an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

// Time: O(n), Space: O(h)
func MaximumDepthOfNAryTree(root *Node) int {
	if root == nil {
		return 0
	}
	maxDepth := 0
	for _, child := range root.Children {
		if depth := MaximumDepthOfNAryTree(child); depth > maxDepth {
			maxDepth = depth
		}
	}
	return maxDepth + 1
}

func main() {
	// Test: [1,null,3,2,4,null,5,6]
	root1 := &Node{
		Val: 1,
		Children: []*Node{
			{Val: 3, Children: []*Node{
				{Val: 5},
				{Val: 6},
			}},
			{Val: 2},
			{Val: 4},
		},
	}
	fmt.Println(MaximumDepthOfNAryTree(root1))

	// Test: single node
	root2 := &Node{Val: 1}
	fmt.Println(MaximumDepthOfNAryTree(root2))
}
```
