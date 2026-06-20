# 0558 — Logical Or Of Two Binary Grids Represented As Quad Trees

## Deskripsi

**Soal:** [0558. Logical Or Of Two Binary Grids Represented As Quad Trees](https://leetcode.com/problems/logical-or-of-two-binary-grids-represented-as-quad-trees/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) where n = number of nodes in smaller tree  
**Kompleksitas Ruang:** O(log n)

**Algoritma:** —

## Solusi Go

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
