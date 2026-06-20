# 2792 — Count Nodes That Are Great Enough

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func countGreatEnoughNodes(root *TreeNode, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2792: Count Nodes That Are Great Enough
// https://leetcode.com/problems/count-nodes-that-are-great-enough/
// Difficulty: Hard [Paid]
//
// A node is "great enough" if its subtree contains >= k nodes AND its value
// is greater than at least k values in its subtree. Post-order DFS returns
// the k smallest values (k <= 10) from each subtree, merged and propagated up.
// O(N*k) time, O(k*h) space.

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countGreatEnoughNodes(root *TreeNode, k int) int {
	if k == 0 || root == nil {
		result := 0
		var count func(*TreeNode)
		count = func(n *TreeNode) {
			if n == nil {
				return
			}
			result++
			count(n.Left)
			count(n.Right)
		}
		count(root)
		return result
	}

	result := 0
	postOrder(root, k, &result)
	return result
}

// postOrder returns up to k smallest values in the subtree, sorted ascending.
func postOrder(node *TreeNode, k int, result *int) []int {
	if node == nil {
		return []int{}
	}

	left := postOrder(node.Left, k, result)
	right := postOrder(node.Right, k, result)

	merged := mergeSorted(left, right)
	if len(merged) > k {
		merged = merged[:k]
	}

	if len(merged) >= k && merged[k-1] < node.Val {
		*result++
	}

	// Insert node.Val in sorted order, keep at most k
	pos := sort.Search(len(merged), func(i int) bool { return merged[i] >= node.Val })
	merged = append(merged, 0)
	copy(merged[pos+1:], merged[pos:])
	merged[pos] = node.Val
	if len(merged) > k {
		merged = merged[:k]
	}

	return merged
}

func mergeSorted(a, b []int) []int {
  // Alokasi slice
	res := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	res = append(res, a[i:]...)
	res = append(res, b[j:]...)
	return res
}

func main() {
	// Example 1: [7,6,5,4,3,2,1], k=2 => 3
	root1 := &TreeNode{Val: 7,
		Left:  &TreeNode{Val: 6, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}},
	}
	fmt.Println(countGreatEnoughNodes(root1, 2))

	// Example 2: [1,2,3], k=1 => 0
	fmt.Println(countGreatEnoughNodes(
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, 1))

	// Example 3: [3,2,2], k=2 => 1
	fmt.Println(countGreatEnoughNodes(
		&TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}, 2))

	// Single node
	fmt.Println(countGreatEnoughNodes(&TreeNode{Val: 5}, 1))

	// k=0 (all nodes qualify)
	fmt.Println(countGreatEnoughNodes(
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, 0))

	// k larger than subtree
	fmt.Println(countGreatEnoughNodes(&TreeNode{Val: 5}, 10))

	// Linear right-skewed tree
	root := &TreeNode{Val: 1}
	curr := root
	for i := 2; i <= 10; i++ {
		curr.Right = &TreeNode{Val: i}
		curr = curr.Right
	}
	fmt.Println(countGreatEnoughNodes(root, 3))
}
```
