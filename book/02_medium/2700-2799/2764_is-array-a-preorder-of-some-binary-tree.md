# 2764 — Is Array A Preorder Of Some Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func IsArrayAPreorderOfSomeBinaryTree(nodes [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2764: Is Array a Preorder of Some Binary Tree
// https://leetcode.com/problems/is-array-a-preorder-of-some-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func IsArrayAPreorderOfSomeBinaryTree(nodes [][]int) bool {
	// nodes[i] = [id, parentId]
  // Membuat map (HashMap) — pencarian O(1)
	children := make(map[int][]int)
	for _, node := range nodes {
		id, parent := node[0], node[1]
		children[parent] = append(children[parent], id)
	}

	// Simulate DFS preorder
	stack := []int{-1} // root parent
	idx := 0
	n := len(nodes)

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		childList := children[cur]
		for i := len(childList) - 1; i >= 0; i-- {
			if idx >= n || childList[i] != nodes[idx][0] {
				return false
			}
			stack = append(stack, childList[i])
			idx++
		}
	}

	return idx == n
}

func main() {
	fmt.Println(IsArrayAPreorderOfSomeBinaryTree([][]int{{0, -1}, {1, 0}, {2, 0}}))
	fmt.Println(IsArrayAPreorderOfSomeBinaryTree([][]int{{0, -1}, {1, 0}, {3, 2}, {2, 1}}))
}
```
