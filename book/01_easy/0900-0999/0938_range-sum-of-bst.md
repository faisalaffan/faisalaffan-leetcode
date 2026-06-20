# 0938 — Range Sum Of Bst

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func rangeSumBST(root *TreeNode, low int, high int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).


## 💻 Solusi Go

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
