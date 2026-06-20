# 0113 — Path Sum Ii

## Deskripsi

**Soal:** [0113. Path Sum Ii](https://leetcode.com/problems/path-sum-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func pathSum(root *TreeNode, targetSum int) [][]int`

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
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
