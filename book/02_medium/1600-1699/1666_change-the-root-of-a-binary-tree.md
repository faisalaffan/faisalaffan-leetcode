# 1666 — Change The Root Of A Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func flipBinaryTree(root *Node, leaf *Node) *Node
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1) where n = path length from leaf to root  
**Kompleksitas Ruang:** O(1) where n = path length from leaf to root

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1666: Change the Root of a Binary Tree
// https://leetcode.com/problems/change-the-root-of-a-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1) where n = path length from leaf to root

import "fmt"

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func flipBinaryTree(root *Node, leaf *Node) *Node {
	if leaf == nil {
		return nil
	}

	cur := leaf
	var newParent *Node

	for cur != nil {
		oldParent := cur.Parent
		cur.Parent = newParent

		// Clear the link if newParent was one of cur's children
		if cur.Left == newParent {
			cur.Left = nil
		}
		if cur.Right == newParent {
			cur.Right = nil
		}

		// Make the old parent a child of current node
		if oldParent != nil {
			if cur.Right == nil {
				cur.Right = oldParent
			} else if cur.Left == nil {
				cur.Left = oldParent
			} else {
				// Both children occupied: move left to right, put oldParent in left
				cur.Right = cur.Left
				cur.Left = oldParent
			}
		}

		newParent = cur
		cur = oldParent
	}

	return leaf
}

func printTree(node *Node, indent string) {
	if node == nil {
		return
	}
	fmt.Printf("%sNode(%d)", indent, node.Val)
	if node.Parent != nil {
		fmt.Printf(" parent=%d", node.Parent.Val)
	} else {
		fmt.Printf(" parent=nil")
	}
	fmt.Println()
	printTree(node.Left, indent+"  L:")
	printTree(node.Right, indent+"  R:")
}

func main() {
	// Test case 1: leaf = 5 (a true leaf)
	// Tree:
	//     1
	//    / \
	//   2   3
	//  / \
	// 4   5
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2, Parent: root}
	root.Right = &Node{Val: 3, Parent: root}
	root.Left.Left = &Node{Val: 4, Parent: root.Left}
	root.Left.Right = &Node{Val: 5, Parent: root.Left}

	newRoot := flipBinaryTree(root, root.Left.Right)
	fmt.Println("Test 1: Flip with leaf=5")
	fmt.Println("New root is leaf:", newRoot.Val == 5)
	fmt.Println("New root parent nil:", newRoot.Parent == nil)
	fmt.Println("5.Right should be 2:", newRoot.Right.Val == 2)
	fmt.Println("2.Parent should be 5:", newRoot.Right.Parent.Val == 5)
	fmt.Println()

	// Test case 2: leaf = root (no change needed)
	root2 := &Node{Val: 10}
	root2.Left = &Node{Val: 20, Parent: root2}
	newRoot2 := flipBinaryTree(root2, root2)
	fmt.Println("Test 2: Leaf is root (no change)")
	fmt.Println("Root unchanged:", newRoot2.Val == 10)
	fmt.Println("Root parent nil:", newRoot2.Parent == nil)
	fmt.Println()

	// Test case 3: Simple chain
	// 1 -> 2 -> 3, leaf = 3
	root3 := &Node{Val: 1}
	root3.Right = &Node{Val: 2, Parent: root3}
	root3.Right.Right = &Node{Val: 3, Parent: root3.Right}

	newRoot3 := flipBinaryTree(root3, root3.Right.Right)
	fmt.Println("Test 3: Chain flip with leaf=3")
	fmt.Println("New root is leaf:", newRoot3.Val == 3)
	fmt.Println("3.Right should be 2:", newRoot3.Right.Val == 2)
	fmt.Println("2.Right should be 1:", newRoot3.Right.Right.Val == 1)
	fmt.Println("1.Parent should be 2:", newRoot3.Right.Right.Parent.Val == 2)
	fmt.Println()
}
```
