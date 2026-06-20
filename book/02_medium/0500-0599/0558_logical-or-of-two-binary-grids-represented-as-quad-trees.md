# 0558 — Logical Or Of Two Binary Grids Represented As Quad Trees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func Intersect(quadTree1 *QuadTreeNode, quadTree2 *QuadTreeNode) *QuadTreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) where n = number of nodes in smaller tree  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #558: Logical OR of Two Binary Grids Represented as Quad-Trees
// https://leetcode.com/problems/logical-or-of-two-binary-grids-represented-as-quad-trees/
// Difficulty: Medium
// Time: O(n) where n = number of nodes in smaller tree
// Space: O(log n)

import "fmt"

type QuadTreeNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadTreeNode
	TopRight    *QuadTreeNode
	BottomLeft  *QuadTreeNode
	BottomRight *QuadTreeNode
}

func main() {
	// Tree 1: full leaf (true)
	t1 := &QuadTreeNode{Val: true, IsLeaf: true}
	// Tree 2: full leaf (false)
	t2 := &QuadTreeNode{Val: false, IsLeaf: true}
	result := Intersect(t1, t2)
	fmt.Printf("leaf=%v, val=%v\n", result.IsLeaf, result.Val)
}

func Intersect(quadTree1 *QuadTreeNode, quadTree2 *QuadTreeNode) *QuadTreeNode {
	if quadTree1.IsLeaf {
		if quadTree1.Val {
			return &QuadTreeNode{Val: true, IsLeaf: true}
		}
		return quadTree2
	}
	if quadTree2.IsLeaf {
		if quadTree2.Val {
			return &QuadTreeNode{Val: true, IsLeaf: true}
		}
		return quadTree1
	}

	topLeft := Intersect(quadTree1.TopLeft, quadTree2.TopLeft)
	topRight := Intersect(quadTree1.TopRight, quadTree2.TopRight)
	bottomLeft := Intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
	bottomRight := Intersect(quadTree1.BottomRight, quadTree2.BottomRight)

	// If all four children are leaves and have same value, merge
	if topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf &&
		topLeft.Val == topRight.Val && topRight.Val == bottomLeft.Val && bottomLeft.Val == bottomRight.Val {
		return &QuadTreeNode{Val: topLeft.Val, IsLeaf: true}
	}

	return &QuadTreeNode{
		IsLeaf:      false,
		TopLeft:     topLeft,
		TopRight:    topRight,
		BottomLeft:  bottomLeft,
		BottomRight: bottomRight,
	}
}
```
