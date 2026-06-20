# 1261 — Find Elements In A Contaminated Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func Constructor(root *TreeNode) FindElements`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n) init, O(1) find  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1261: Find Elements in a Contaminated Binary Tree
// https://leetcode.com/problems/find-elements-in-a-contaminated-binary-tree/
// Difficulty: Medium

// Tree is contaminated (all vals = -1). Recover with rules:
// root.val = 0, left.val = 2*parent+1, right.val = 2*parent+2.
// Then support Find(target) operation.

// Time: O(n) init, O(1) find
// Space: O(n)

type FindElements struct {
	vals map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	fe := FindElements{vals: make(map[int]bool)}
	fe.recover(root, 0)
	return fe
}

func (fe *FindElements) recover(node *TreeNode, val int) {
	if node == nil {
		return
	}
	node.Val = val
	fe.vals[val] = true
	fe.recover(node.Left, 2*val+1)
	fe.recover(node.Right, 2*val+2)
}

func (fe *FindElements) Find(target int) bool {
	return fe.vals[target]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [-1,null,-1]
	root := &TreeNode{Val: -1, Right: &TreeNode{Val: -1}}
	fe := Constructor(root)
	fmt.Printf("%t (expected: false)\n", fe.Find(1))
	fmt.Printf("%t (expected: true)\n", fe.Find(2))

	// Tree: [-1,-1,-1,-1,-1]
	root2 := &TreeNode{
		Val: -1,
		Left: &TreeNode{
			Val:   -1,
			Left:  &TreeNode{Val: -1},
			Right: &TreeNode{Val: -1},
		},
		Right: &TreeNode{Val: -1},
	}
	fe2 := Constructor(root2)
	fmt.Printf("%t (expected: true)\n", fe2.Find(1))
	fmt.Printf("%t (expected: true)\n", fe2.Find(3))
	fmt.Printf("%t (expected: false)\n", fe2.Find(5))
}
```
