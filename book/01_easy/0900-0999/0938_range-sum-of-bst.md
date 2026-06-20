# 0938 — Range Sum Of Bst

## Deskripsi

**Soal:** [0938. Range Sum Of Bst](https://leetcode.com/problems/range-sum-of-bst/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #938: Range Sum of BST
// https://leetcode.com/problems/range-sum-of-bst/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 7},
		},
		Right: &TreeNode{
			Val:   15,
			Right: &TreeNode{Val: 18},
		},
	}
	fmt.Println(rangeSumBST(root, 7, 15)) // 32
}

// rangeSumBST returns the sum of all node values in the range [low, high].
// Time: O(n). Space: O(n).
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}
```
