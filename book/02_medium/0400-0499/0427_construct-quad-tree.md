# 0427 — Construct Quad Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func construct(grid [][]int) *Node
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2 log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #427: Construct Quad Tree
// https://leetcode.com/problems/construct-quad-tree/
// Difficulty: Medium
// Time: O(n^2 log n) | Space: O(log n)

import "fmt"

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	return build(grid, 0, 0, len(grid))
}

func build(grid [][]int, r, c, size int) *Node {
	if size == 1 {
		return &Node{Val: grid[r][c] == 1, IsLeaf: true}
	}

	half := size / 2
	topLeft := build(grid, r, c, half)
	topRight := build(grid, r, c+half, half)
	bottomLeft := build(grid, r+half, c, half)
	bottomRight := build(grid, r+half, c+half, half)

	// Check if all children are leaves with same value
	if topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf &&
		topLeft.Val == topRight.Val && topRight.Val == bottomLeft.Val && bottomLeft.Val == bottomRight.Val {
		return &Node{Val: topLeft.Val, IsLeaf: true}
	}

	return &Node{
		IsLeaf:      false,
		Val:         true,
		TopLeft:     topLeft,
		TopRight:    topRight,
		BottomLeft:  bottomLeft,
		BottomRight: bottomRight,
	}
}

func printQuadTree(node *Node, indent string) {
	if node == nil {
		return
	}
	if node.IsLeaf {
		fmt.Printf("%sLeaf: %v\n", indent, node.Val)
	} else {
		fmt.Printf("%sInternal:\n", indent)
		printQuadTree(node.TopLeft, indent+"  TL: ")
		printQuadTree(node.TopRight, indent+"  TR: ")
		printQuadTree(node.BottomLeft, indent+"  BL: ")
		printQuadTree(node.BottomRight, indent+"  BR: ")
	}
}

func main() {
	// Test case 1
	grid1 := [][]int{
		{0, 1},
		{1, 0},
	}
	fmt.Println("Test 1:")
	n1 := construct(grid1)
	fmt.Println("  IsLeaf:", n1.IsLeaf) // false (quadtree)

	// Test case 2: All same value
	grid2 := [][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}
	n2 := construct(grid2)
	fmt.Println("Test 2: IsLeaf:", n2.IsLeaf, "Val:", n2.Val)
	// Expected: true, 1

	// Test case 3
	grid3 := [][]int{
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
	}
	n3 := construct(grid3)
	fmt.Println("Test 3: IsLeaf:", n3.IsLeaf)
	// Expected: false
}
```
