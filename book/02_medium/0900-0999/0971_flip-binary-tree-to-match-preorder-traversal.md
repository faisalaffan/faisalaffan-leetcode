# 0971 — Flip Binary Tree To Match Preorder Traversal

## Deskripsi

**Soal:** [0971. Flip Binary Tree To Match Preorder Traversal](https://leetcode.com/problems/flip-binary-tree-to-match-preorder-traversal/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func flipMatchVoyage(root *TreeNode, voyage []int) []int`

## Solusi Go

```go
package main

// LeetCode #971: Flip Binary Tree To Match Preorder Traversal
// https://leetcode.com/problems/flip-binary-tree-to-match-preorder-traversal/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(n)
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	i := 0
	ok := true
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, 0)

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || !ok {
			return
		}
		if node.Val != voyage[i] {
			ok = false
			return
		}
		i++
		if node.Left != nil && node.Left.Val != voyage[i] {
			ans = append(ans, node.Val)
			dfs(node.Right)
			dfs(node.Left)
		} else {
			dfs(node.Left)
			dfs(node.Right)
		}
	}

	dfs(root)
	if !ok {
		return []int{-1}
	}
	return ans
}

func main() {
	// [1,2] voyage=[2,1]
	root1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	fmt.Println(flipMatchVoyage(root1, []int{2, 1}))

	// [1,2,3] voyage=[1,3,2]
	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(flipMatchVoyage(root2, []int{1, 3, 2}))

	// [1,2,3] voyage=[1,2,3]
	root3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(flipMatchVoyage(root3, []int{1, 2, 3}))
}
```
