# 1973 — Count Nodes Equal To Sum Of Descendants

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func EqualToDescendants(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, DFS

**Kompleksitas Waktu:** O(n), Space: O(h) where h is tree height  
**Kompleksitas Ruang:** O(h) where h is tree height

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1973: Count Nodes Equal to Sum of Descendants
// https://leetcode.com/problems/count-nodes-equal-to-sum-of-descendants/
// Difficulty: Medium [Paid]

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 10,
		Left:  &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{Val: 5}}
	fmt.Println(EqualToDescendants(root))

	root2 := &TreeNode{Val: 2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 1}}
	fmt.Println(EqualToDescendants(root2))
}

// Time: O(n), Space: O(h) where h is tree height
func EqualToDescendants(root *TreeNode) int {
	count := 0
	dfs(root, &count)
	return count
}

func dfs(node *TreeNode, count *int) int64 {
	if node == nil {
		return 0
	}
	leftSum := dfs(node.Left, count)
	rightSum := dfs(node.Right, count)
	if leftSum+rightSum == int64(node.Val) {
		*count++
	}
	return leftSum + rightSum + int64(node.Val)
}
```
