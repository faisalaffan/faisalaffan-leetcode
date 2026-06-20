# 1485 — Clone Binary Tree With Random Pointer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func copyRandomBinaryTree(root *Node) *NodeCopy
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n) where n = number of nodes  
**Kompleksitas Ruang:** O(n) for the map

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1485: Clone Binary Tree With Random Pointer
// https://leetcode.com/problems/clone-binary-tree-with-random-pointer/
// Difficulty: Medium

import "fmt"

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Random *Node
}

type NodeCopy struct {
	Val    int
	Left   *NodeCopy
	Right  *NodeCopy
	Random *NodeCopy
}

func main() {
	// Test case 1
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Random = root.Right
	root.Right.Random = root.Left

	copied := copyRandomBinaryTree(root)
	fmt.Println(copied.Val) // 1
	fmt.Println(copied.Left.Val) // 2
	fmt.Println(copied.Right.Val) // 3
	fmt.Println(copied.Left.Random.Val) // 3
	fmt.Println(copied.Right.Random.Val) // 2

	// Test case 2 - nil
	fmt.Println(copyRandomBinaryTree(nil)) // nil

	// Test case 3 - single node
	root3 := &Node{Val: 42}
	copied3 := copyRandomBinaryTree(root3)
	fmt.Println(copied3.Val) // 42
}

// Time: O(n) where n = number of nodes
// Space: O(n) for the map
func copyRandomBinaryTree(root *Node) *NodeCopy {
	if root == nil {
		return nil
	}

	// Map from original node to copy
  // Membuat map (HashMap) — pencarian O(1)
	nodeMap := make(map[*Node]*NodeCopy)

	var dfs func(*Node) *NodeCopy
	dfs = func(node *Node) *NodeCopy {
		if node == nil {
			return nil
		}
		if copy, ok := nodeMap[node]; ok {
			return copy
		}
		copy := &NodeCopy{Val: node.Val}
		nodeMap[node] = copy
		copy.Left = dfs(node.Left)
		copy.Right = dfs(node.Right)
		return copy
	}

	rootCopy := dfs(root)

	// Set random pointers
	var setRandom func(*Node, *NodeCopy)
	setRandom = func(orig *Node, copy *NodeCopy) {
		if orig == nil || copy == nil {
			return
		}
		if orig.Random != nil {
			copy.Random = nodeMap[orig.Random]
		}
		setRandom(orig.Left, copy.Left)
		setRandom(orig.Right, copy.Right)
	}
	setRandom(root, rootCopy)

	return rootCopy
}
```
