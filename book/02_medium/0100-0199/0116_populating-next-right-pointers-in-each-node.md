# 0116 — Populating Next Right Pointers In Each Node

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func connect(root *Node) *Node
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #116: Populating Next Right Pointers in Each Node
// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
// Difficulty: Medium

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	leftmost := root
	for leftmost.Left != nil {
		head := leftmost
		for head != nil {
			head.Left.Next = head.Right
			if head.Next != nil {
				head.Right.Next = head.Next.Left
			}
			head = head.Next
		}
		leftmost = leftmost.Left
	}

	return root
}

func printLevels(root *Node) {
	curr := root
	for curr != nil {
		head := curr
		for head != nil {
			nextStr := "null"
			if head.Next != nil {
				nextStr = fmt.Sprintf("%d", head.Next.Val)
			}
			fmt.Printf("%d->%s ", head.Val, nextStr)
			head = head.Next
		}
		fmt.Println()
		curr = curr.Left
	}
}

func main() {
	// Test case 1: [1,2,3,4,5,6,7]
	root := &Node{Val: 1,
		Left:  &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}},
		Right: &Node{Val: 3, Left: &Node{Val: 6}, Right: &Node{Val: 7}}}
	connect(root)
	printLevels(root)

	// Test case 2
	root = nil
	connect(root)
	fmt.Println("nil")
}

// Time: O(n) | Space: O(1)
```
