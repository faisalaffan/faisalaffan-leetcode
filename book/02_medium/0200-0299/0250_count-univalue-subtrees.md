# 0250 — Count Univalue Subtrees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func countUnivalSubtrees(root *TreeNode) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n), Space: O(h)  |  **Ruang:** O(h)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #250: Count Univalue Subtrees
// https://leetcode.com/problems/count-univalue-subtrees/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countUnivalSubtrees(root *TreeNode) int {
	count := 0
	isUnival(root, &count)
	return count
}

func isUnival(node *TreeNode, count *int) bool {
	if node == nil {
		return true
	}

	leftUni := isUnival(node.Left, count)
	rightUni := isUnival(node.Right, count)

	if !leftUni || !rightUni {
		return false
	}

	if node.Left != nil && node.Left.Val != node.Val {
		return false
	}
	if node.Right != nil && node.Right.Val != node.Val {
		return false
	}

	*count++
	return true
}

func main() {
	root := &TreeNode{5, &TreeNode{1, &TreeNode{5, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{5, nil, &TreeNode{5, nil, nil}}}
	fmt.Println(countUnivalSubtrees(root))

	root2 := &TreeNode{5, &TreeNode{5, &TreeNode{5, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{5, nil, nil}}
	fmt.Println(countUnivalSubtrees(root2))

	fmt.Println(countUnivalSubtrees(nil))
}
```
