# 0653 — Two Sum Iv Input Is A Bst

## Deskripsi

**Soal:** [0653. Two Sum Iv Input Is A Bst](https://leetcode.com/problems/two-sum-iv-input-is-a-bst/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #653: Two Sum IV - Input is a BST
// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
// Difficulty: Easy

import "fmt"

// TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example: [5,3,6,2,4,null,7], k=9 => true
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   6,
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(findTarget(root, 9))  // true
	fmt.Println(findTarget(root, 28)) // false
}

// findTarget returns true if there exist two elements in the BST that sum to k.
// Time: O(n). Space: O(n).
func findTarget(root *TreeNode, k int) bool {
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[int]bool)
	return find(root, k, seen)
}

func find(node *TreeNode, k int, seen map[int]bool) bool {
	if node == nil {
		return false
	}
	if seen[k-node.Val] {
		return true
	}
	seen[node.Val] = true
	return find(node.Left, k, seen) || find(node.Right, k, seen)
}
```
