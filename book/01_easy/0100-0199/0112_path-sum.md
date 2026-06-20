# 0112 — Path Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func HasPathSum(root *TreeNode, targetSum int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(h)


## 💻 Solusi Go

```go
package main

// LeetCode #112: Path Sum
// https://leetcode.com/problems/path-sum/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(h)
func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}
	return HasPathSum(root.Left, targetSum) || HasPathSum(root.Right, targetSum)
}

func main() {
	root := &TreeNode{5, &TreeNode{4, &TreeNode{11, &TreeNode{7, nil, nil}, &TreeNode{2, nil, nil}}, nil}, &TreeNode{8, &TreeNode{13, nil, nil}, &TreeNode{4, nil, &TreeNode{1, nil, nil}}}}
	fmt.Println(HasPathSum(root, 22))
	fmt.Println(HasPathSum(nil, 0))
}
```
