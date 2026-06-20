# 0366 — Find Leaves Of Binary Tree

## Deskripsi

**Soal:** [0366. Find Leaves Of Binary Tree](https://leetcode.com/problems/find-leaves-of-binary-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func findLeaves(root *TreeNode) [][]int`

## Solusi Go

```go
package main

// LeetCode #366: Find Leaves of Binary Tree
// https://leetcode.com/problems/find-leaves-of-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLeaves(root *TreeNode) [][]int {
	result := [][]int{}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		leftHeight := dfs(node.Left)
		rightHeight := dfs(node.Right)
		height := max(leftHeight, rightHeight) + 1

		if height >= len(result) {
			result = append(result, []int{})
		}
		result[height] = append(result[height], node.Val)
		return height
	}
	dfs(root)
	return result
}

func main() {
	// Test case 1: [1,2,3,4,5]
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
	root1.Right = &TreeNode{Val: 3}
	fmt.Println("Test 1:", findLeaves(root1))
	// Expected: [[4,5,3],[2],[1]]

	// Test case 2: Single node
	root2 := &TreeNode{Val: 1}
	fmt.Println("Test 2:", findLeaves(root2))
	// Expected: [[1]]

	// Test case 3: Nil
	fmt.Println("Test 3:", findLeaves(nil))
	// Expected: []
}
```
