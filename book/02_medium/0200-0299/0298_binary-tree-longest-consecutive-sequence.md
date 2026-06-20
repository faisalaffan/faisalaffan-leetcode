# 0298 — Binary Tree Longest Consecutive Sequence

## Deskripsi

**Soal:** [0298. Binary Tree Longest Consecutive Sequence](https://leetcode.com/problems/binary-tree-longest-consecutive-sequence/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(h)  
**Kompleksitas Ruang:** O(h)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func longestConsecutive(root *TreeNode) int`

## Solusi Go

```go
package main

// LeetCode #298: Binary Tree Longest Consecutive Sequence
// https://leetcode.com/problems/binary-tree-longest-consecutive-sequence/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestConsecutive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxLen := 1
	var dfs func(node *TreeNode, parentVal int, length int)
	dfs = func(node *TreeNode, parentVal int, length int) {
		if node == nil {
			return
		}

		if node.Val == parentVal+1 {
			length++
		} else {
			length = 1
		}

		if length > maxLen {
			maxLen = length
		}

		dfs(node.Left, node.Val, length)
		dfs(node.Right, node.Val, length)
	}

	dfs(root, root.Val, 1)
	return maxLen
}

func main() {
	root := &TreeNode{1, nil, &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, &TreeNode{5, nil, nil}}}}
	fmt.Println(longestConsecutive(root))

	root2 := &TreeNode{2, nil, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, nil}}
	fmt.Println(longestConsecutive(root2))

	fmt.Println(longestConsecutive(nil))
}
```
