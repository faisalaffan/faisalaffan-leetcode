# 1612 — Check If Two Expression Trees Are Equivalent

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckEquivalence(root1 *Node, root2 *Node) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1612: Check If Two Expression Trees are Equivalent
// https://leetcode.com/problems/check-if-two-expression-trees-are-equivalent/
// Difficulty: Medium [Paid]

import "fmt"

// Node is an expression tree node.
// Leaf nodes contain lowercase letters.
// Non-leaf nodes contain '+' operator.
type Node struct {
	Val    byte
	Left   *Node
	Right  *Node
}

func main() {
	// Expression tree 1: (a + (b + c))
	root1 := &Node{Val: '+'}
	root1.Left = &Node{Val: 'a'}
	root1.Right = &Node{Val: '+', Left: &Node{Val: 'b'}, Right: &Node{Val: 'c'}}

	// Expression tree 2: ((b + a) + c)
	root2 := &Node{Val: '+'}
	root2.Left = &Node{Val: '+', Left: &Node{Val: 'b'}, Right: &Node{Val: 'a'}}
	root2.Right = &Node{Val: 'c'}

	fmt.Println(CheckEquivalence(root1, root2)) // true

	// Different trees
	root3 := &Node{Val: '+'}
	root3.Left = &Node{Val: 'a'}
	root3.Right = &Node{Val: 'b'}

	fmt.Println(CheckEquivalence(root1, root3)) // false
}

func CheckEquivalence(root1 *Node, root2 *Node) bool {
	// Time: O(N), Space: O(1)
	// Expression trees only contain '+', so equivalence = same multiset of leaf values
  // Membuat map (HashMap) — pencarian O(1)
	count1 := make(map[byte]int)
  // Membuat map (HashMap) — pencarian O(1)
	count2 := make(map[byte]int)

	countLeaves(root1, count1)
	countLeaves(root2, count2)

	if len(count1) != len(count2) {
		return false
	}
	for k, v := range count1 {
		if count2[k] != v {
			return false
		}
	}
	return true
}

func countLeaves(node *Node, count map[byte]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		count[node.Val]++
		return
	}
	countLeaves(node.Left, count)
	countLeaves(node.Right, count)
}
```
