# 1490 — Clone N Ary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func cloneTree(root *Node) *Node
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1490: Clone N-ary Tree
// https://leetcode.com/problems/clone-n-ary-tree/
// Difficulty: Medium (listed here as Hard)
//
// Clone an N-ary tree. Each node has a value and a list of children.

import "fmt"

// Node represents an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

// cloneTree creates a deep copy of the N-ary tree.
func cloneTree(root *Node) *Node {
	if root == nil {
		return nil
	}

	// Recursively clone children
	clonedChildren := make([]*Node, len(root.Children))
	for i, child := range root.Children {
		clonedChildren[i] = cloneTree(child)
	}

	return &Node{
		Val:      root.Val,
		Children: clonedChildren,
	}
}

// Helper function to compare two trees for testing
func equalTrees(a, b *Node) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	if len(a.Children) != len(b.Children) {
		return false
	}
  // Range loop: iterasi dengan indeks + nilai
	for i := range a.Children {
		if !equalTrees(a.Children[i], b.Children[i]) {
			return false
		}
	}
	return true
}

// Helper to print tree (preorder)
func printTree(root *Node, indent int) {
	if root == nil {
		return
	}
	for i := 0; i < indent; i++ {
		fmt.Print("  ")
	}
	fmt.Printf("Node(%d)\n", root.Val)
	for _, child := range root.Children {
		printTree(child, indent+1)
	}
}

func main() {
	// Test case 1: single node
	root1 := &Node{Val: 1}
	clone1 := cloneTree(root1)
	fmt.Printf("Test 1 - Clone of single node: equal=%v (expected true)\n", equalTrees(root1, clone1))
	fmt.Printf("  Original pointer: %p, Clone pointer: %p (different: %v)\n",
		root1, clone1, root1 != clone1)

	// Test case 2: tree with children
	root2 := &Node{
		Val: 1,
		Children: []*Node{
			{Val: 2, Children: []*Node{
				{Val: 5},
				{Val: 6},
			}},
			{Val: 3},
			{Val: 4, Children: []*Node{
				{Val: 7},
				{Val: 8},
				{Val: 9},
			}},
		},
	}
	clone2 := cloneTree(root2)
	fmt.Printf("\nTest 2 - Clone of complex tree: equal=%v (expected true)\n", equalTrees(root2, clone2))

	// Modify original and ensure clone is unchanged
	root2.Children[0].Val = 99
	fmt.Printf("  After modifying original: equal=%v (expected false)\n", equalTrees(root2, clone2))

	fmt.Println("\n  Original tree:")
	printTree(clone2, 0)
	fmt.Println("  (clone preserved after original modification)")

	// Test case 3: nil tree
	clone3 := cloneTree(nil)
	fmt.Printf("\nTest 3 - Clone of nil: %v (expected nil)\n", clone3)

	// Test case 4: deep structural verification
	root4 := &Node{
		Val: 10,
		Children: []*Node{
			{Val: 20},
		},
	}
	clone4 := cloneTree(root4)
	fmt.Printf("\nTest 4 - Structural equality: equal=%v (expected true)\n", equalTrees(root4, clone4))
	// Verify they are truly independent
	root4.Children[0].Val = 30
	fmt.Printf("  After modifying child: equal=%v (expected false)\n", equalTrees(root4, clone4))
	fmt.Printf("  Clone child val: %d (expected 20)\n", clone4.Children[0].Val)
}
```
