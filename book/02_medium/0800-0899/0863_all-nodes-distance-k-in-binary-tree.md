# 0863 — All Nodes Distance K In Binary Tree

## Deskripsi

**Soal:** [0863. All Nodes Distance K In Binary Tree](https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #863: All Nodes Distance K in Binary Tree
// https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1: [3,5,1,6,2,0,8,null,null,7,4], target=5, k=2 -> [7,4,1]
	root := &TreeNode{3,
		&TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}}},
		&TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{8, nil, nil}},
	}
	fmt.Println(AllNodesDistanceKInBinaryTree(root, root.Left, 2))

	// Test case 2: target=5, k=1 -> [6,2,3]
	fmt.Println(AllNodesDistanceKInBinaryTree(root, root.Left, 1))

	// Test case 3: k=0
	fmt.Println(AllNodesDistanceKInBinaryTree(root, root.Left, 0))
}

// Time: O(n) | Space: O(n)
func AllNodesDistanceKInBinaryTree(root *TreeNode, target *TreeNode, k int) []int {
  // Membuat map untuk pencarian O(1): key → value
	parent := make(map[*TreeNode]*TreeNode)
	var buildParent func(*TreeNode, *TreeNode)
	buildParent = func(node, par *TreeNode) {
		if node == nil {
			return
		}
		parent[node] = par
		buildParent(node.Left, node)
		buildParent(node.Right, node)
	}
	buildParent(root, nil)

  // Membuat map untuk pencarian O(1): key → value
	visited := make(map[*TreeNode]bool)
	var ans []int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, dist int) {
		if node == nil || visited[node] {
			return
		}
		visited[node] = true
		if dist == k {
			ans = append(ans, node.Val)
			return
		}
		dfs(node.Left, dist+1)
		dfs(node.Right, dist+1)
		dfs(parent[node], dist+1)
	}
	dfs(target, 0)

	return ans
}
```
