# 0113 — Path Sum Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func pathSum(root *TreeNode, targetSum int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #113: Path Sum II
// https://leetcode.com/problems/path-sum-ii/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	var dfs func(node *TreeNode, remaining int, path []int)
	dfs = func(node *TreeNode, remaining int, path []int) {
		if node == nil {
			return
		}
		remaining -= node.Val
		path = append(path, node.Val)

		if node.Left == nil && node.Right == nil && remaining == 0 {
  // Alokasi slice integer
			validPath := make([]int, len(path))
			copy(validPath, path)
			result = append(result, validPath)
		} else {
			dfs(node.Left, remaining, path)
			dfs(node.Right, remaining, path)
		}

		path = path[:len(path)-1]
	}
	dfs(root, targetSum, []int{})
	return result
}

func main() {
	// Test case 1
	root := &TreeNode{Val: 5,
		Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}}}
	fmt.Println(pathSum(root, 22)) // [[5 4 11 2] [5 8 4 5]]

	// Test case 2
	fmt.Println(pathSum(nil, 0)) // []

	// Test case 3: [1,2], target=1 -> []
	root = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	fmt.Println(pathSum(root, 1)) // []
}

// Time: O(n^2) | Space: O(n)
```
