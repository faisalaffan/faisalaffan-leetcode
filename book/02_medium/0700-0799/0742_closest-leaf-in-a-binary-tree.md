# 0742 — Closest Leaf In A Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func findClosestLeaf(root *BTNode, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, BFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #742: Closest Leaf in a Binary Tree
// https://leetcode.com/problems/closest-leaf-in-a-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	root := &BTNode{Val: 1}
	root.Left = &BTNode{Val: 3}
	root.Right = &BTNode{Val: 2}
	fmt.Println(findClosestLeaf(root, 1))
}

type BTNode struct {
	Val   int
	Left  *BTNode
	Right *BTNode
}

func findClosestLeaf(root *BTNode, k int) int {
  // Membuat map (HashMap) — pencarian O(1)
	graph := make(map[int][]int)
  // Membuat map (HashMap) — pencarian O(1)
	leaves := make(map[int]bool)
  // Membuat map (HashMap) — pencarian O(1)
	visited := make(map[int]bool)

	var buildGraph func(node *BTNode)
	buildGraph = func(node *BTNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			leaves[node.Val] = true
		}
		if node.Left != nil {
			graph[node.Val] = append(graph[node.Val], node.Left.Val)
			graph[node.Left.Val] = append(graph[node.Left.Val], node.Val)
			buildGraph(node.Left)
		}
		if node.Right != nil {
			graph[node.Val] = append(graph[node.Val], node.Right.Val)
			graph[node.Right.Val] = append(graph[node.Right.Val], node.Val)
			buildGraph(node.Right)
		}
	}

	buildGraph(root)

	queue := []int{k}
	visited[k] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if leaves[node] {
			return node
		}
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return -1
}
```
