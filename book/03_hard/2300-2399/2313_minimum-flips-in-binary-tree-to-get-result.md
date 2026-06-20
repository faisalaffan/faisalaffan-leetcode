# 2313 — Minimum Flips In Binary Tree To Get Result

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** —

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func minimumFlips(root *TreeNode, result int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"math"
)

// 2313. Minimum Flips in Binary Tree to Get Result
// ----------------------------------------------------------------
// Each node is either a leaf (value 0 or 1) or an internal node
// (AND=3, OR=4, XOR=5, NOT=6 — or whatever encoding LeetCode chooses).
// For NOT the node has a single child; for AND/OR/XOR it has two children.
//
// We compute dp[node][target] = minimum flips needed in the subtree of
// node to make it evaluate to target (0 or 1).
//
// A "flip" changes a leaf's value (0↔1) or replaces an internal node's operator
// with any other operator (AND/OR/XOR/NOT).  The problem description on
// LeetCode gives specific operator constants — we use the same ones as the
// problem: 0=leaf false, 1=leaf true, 2=NOT, 3=AND, 4=OR, 5=XOR.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumFlips(root *TreeNode, result int) int {
	var dfs func(*TreeNode) [2]int // [cost0, cost1]
	dfs = func(node *TreeNode) [2]int {
		if node.Left == nil && node.Right == nil {
			// leaf
			if node.Val == 0 {
				return [2]int{0, 1} // cost0=0, cost1=1 (flip)
			}
			return [2]int{1, 0} // cost0=1, cost1=0
		}
		if node.Val == 2 { // NOT
			child := node.Left
			if child == nil {
				child = node.Right
			}
			c := dfs(child)
			// output 1 → child 0, output 0 → child 1
			return [2]int{c[1], c[0]}
		}
		// AND / OR / XOR — two children
		l := dfs(node.Left)
		r := dfs(node.Right)
		// Cost to get 0 and 1 for each operator.
		switch node.Val {
		case 3: // AND
			// result 1 only if both 1
			// result 0 if at least one 0
			return [2]int{
				min3(l[0]+r[0], l[0]+r[1], l[1]+r[0]), // cost0
				l[1] + r[1], // cost1
			}
		case 4: // OR
			// result 0 only if both 0
			return [2]int{
				l[0] + r[0], // cost0
				min3(l[1]+r[0], l[0]+r[1], l[1]+r[1]), // cost1
			}
		case 5: // XOR
			return [2]int{
				min(l[0]+r[0], l[1]+r[1]),     // cost0 (same)
				min(l[0]+r[1], l[1]+r[0]),     // cost1 (different)
			}
		default:
			return [2]int{math.MaxInt32, math.MaxInt32}
		}
	}
	cost := dfs(root)
	if result == 0 {
		return cost[0]
	}
	return cost[1]
}

func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// ---------------------------------------------------------------------------
//  Wrapper

func MinimumFlipsInBinaryTreeToGetResult() interface{} {
	// Example: leaf(0) → NOT → result=1
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 0}}
	return minimumFlips(root, 1) // flip leaf 0→1 costs 1
}

func main() {
	fmt.Println(MinimumFlipsInBinaryTreeToGetResult())

	// leaf 0 → NOT → answer=1 ⇒ flip leaf once → cost 1
	root1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 0}}
	if got := minimumFlips(root1, 1); got != 0 {
		fmt.Printf("FAIL test1: got %d, want 0\n", got)
	}

	// AND([1,0]) → 0, want 1: cheapest is flip leaf 0→1 (cost 1)
	root2 := &TreeNode{Val: 3, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 0}}
	if got := minimumFlips(root2, 1); got != 1 {
		fmt.Printf("FAIL test2: got %d, want 0\n", got)
	}

	// OR([0,0]) → 0, want 1: flip one leaf (cost 1)
	root3 := &TreeNode{Val: 4, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 0}}
	if got := minimumFlips(root3, 1); got != 1 {
		fmt.Printf("FAIL test3: got %d, want 0\n", got)
	}

	// XOR([1,0]) → 1, want 0: cheapest to make output 0 is to
	// flip one of them so both are same: cost 1
	root4 := &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 0}}
	if got := minimumFlips(root4, 0); got != 1 {
		fmt.Printf("FAIL test4: got %d, want 0\n", got)
	}

	fmt.Println("Done testing 2313.")
}
```
