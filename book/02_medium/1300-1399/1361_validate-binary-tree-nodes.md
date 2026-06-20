# 1361 — Validate Binary Tree Nodes

## Deskripsi

**Soal:** [1361. Validate Binary Tree Nodes](https://leetcode.com/problems/validate-binary-tree-nodes/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) where n = number of nodes  
**Kompleksitas Ruang:** O(n) for in-degree and visited arrays

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman), BFS (Breadth-First Search / pencarian lebar)

## Solusi Go

```go
package main

// LeetCode #1361: Validate Binary Tree Nodes
// https://leetcode.com/problems/validate-binary-tree-nodes/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(validateBinaryTreeNodes(4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1})) // true

	// Test case 2
	fmt.Println(validateBinaryTreeNodes(4, []int{1, -1, 3, -1}, []int{2, 3, -1, -1})) // false

	// Test case 3
	fmt.Println(validateBinaryTreeNodes(2, []int{1, 0}, []int{-1, -1})) // false

	// Test case 4
	fmt.Println(validateBinaryTreeNodes(6, []int{1, -1, -1, 4, -1, -1}, []int{2, -1, -1, 5, -1, -1})) // false
}

// Time: O(n) where n = number of nodes
// Space: O(n) for in-degree and visited arrays
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	// Track in-degree of each node (how many parents)
  // Membuat slice untuk menyimpan hasil
	inDegree := make([]int, n)
	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			inDegree[leftChild[i]]++
			if inDegree[leftChild[i]] > 1 {
				return false
			}
		}
		if rightChild[i] != -1 {
			inDegree[rightChild[i]]++
			if inDegree[rightChild[i]] > 1 {
				return false
			}
		}
	}

	// Find root (node with in-degree 0)
	root := -1
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			if root != -1 {
				return false // more than one root
			}
			root = i
		}
	}
	if root == -1 {
		return false // cycle (no root)
	}

	// BFS/DFS from root to verify all nodes reachable
  // Membuat slice untuk menyimpan hasil
	visited := make([]bool, n)
	var dfs func(int)
	dfs = func(node int) {
		if node == -1 || visited[node] {
			return
		}
		visited[node] = true
		dfs(leftChild[node])
		dfs(rightChild[node])
	}
	dfs(root)

	// All nodes must be visited
	for _, v := range visited {
		if !v {
			return false
		}
	}

	return true
}
```
