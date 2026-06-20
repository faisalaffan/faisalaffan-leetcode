# 1457 — Pseudo Palindromic Paths In A Binary Tree

## Deskripsi

**Soal:** [1457. Pseudo Palindromic Paths In A Binary Tree](https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) where n = number of nodes  
**Kompleksitas Ruang:** O(h) for recursion stack

**Algoritma:** Stack (tumpukan LIFO)

## Solusi Go

```go
package main

// LeetCode #1457: Pseudo-Palindromic Paths in a Binary Tree
// https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{Val: 1},
		},
	}
	fmt.Println(pseudoPalindromicPaths(root)) // 2

	// Test case 2
	root2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}}},
		Right: &TreeNode{Val: 1},
	}
	fmt.Println(pseudoPalindromicPaths(root2)) // 1

	// Test case 3
	fmt.Println(pseudoPalindromicPaths(&TreeNode{Val: 9})) // 1
}

// Time: O(n) where n = number of nodes
// Space: O(h) for recursion stack
func pseudoPalindromicPaths(root *TreeNode) int {
	count := 0
  // Membuat slice untuk menyimpan hasil
	freq := make([]int, 10) // node values are 1-9

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		freq[node.Val]++
		if node.Left == nil && node.Right == nil {
			// Check if path is pseudo-palindromic
			oddCount := 0
			for _, f := range freq {
				if f%2 == 1 {
					oddCount++
				}
			}
			if oddCount <= 1 {
				count++
			}
		} else {
			dfs(node.Left)
			dfs(node.Right)
		}
		freq[node.Val]--
	}

	dfs(root)
	return count
}
```
