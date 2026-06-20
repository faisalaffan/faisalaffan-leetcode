# 3879 — Maximum Distinct Path Sum In A Binary Tree

## Deskripsi

**Soal:** [3879. Maximum Distinct Path Sum In A Binary Tree](https://leetcode.com/problems/maximum-distinct-path-sum-in-a-binary-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N^2)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func maxDistinctPathSum(root *TreeNode, visited map[int]bool) int`

> **Ide Kunci:** DFS from each node as start, exploring all paths with distinct values.

## Solusi Go

```go
package main

// LeetCode #3879: Maximum Distinct Path Sum in a Binary Tree
// https://leetcode.com/problems/maximum-distinct-path-sum-in-a-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(N^2) | Space: O(N)
// Approach: DFS from each node as start, exploring all paths with distinct values.

import (
	"fmt"
	"math"
)

// TreeNode definition for binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDistinctPathSum(root *TreeNode, visited map[int]bool) int {
	if root == nil || visited[root.Val] {
		return 0
	}

	visited[root.Val] = true
	defer func() { delete(visited, root.Val) }()

	leftSum := maxDistinctPathSum(root.Left, visited)
	rightSum := maxDistinctPathSum(root.Right, visited)

	return root.Val + max(leftSum, rightSum)
}

func MaximumDistinctPathSumInABinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := math.MinInt32

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
  // Membuat map untuk pencarian O(1): key → value
		visited := make(map[int]bool)
		sum := maxDistinctPathSum(node, visited)
		if sum > ans {
			ans = sum
		}
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return ans
}

func main() {
	// Example 1: root = [2,2,1]
	root1 := &TreeNode{Val: 2,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 1},
	}
	fmt.Println(MaximumDistinctPathSumInABinaryTree(root1)) // Expected: 3

	// Example 2: root = [1,-2,5,null,null,3,5]
	root2 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: -2},
		Right: &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 5},
		},
	}
	fmt.Println(MaximumDistinctPathSumInABinaryTree(root2)) // Expected: 9

	// Example 3: root = [4,6,6,null,null,null,9]
	root3 := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 6},
		Right: &TreeNode{Val: 6,
			Right: &TreeNode{Val: 9},
		},
	}
	fmt.Println(MaximumDistinctPathSumInABinaryTree(root3)) // Expected: 19
}
```
