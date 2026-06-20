# 1382 — Balance A Binary Search Tree

## Deskripsi

**Soal:** [1382. Balance A Binary Search Tree](https://leetcode.com/problems/balance-a-binary-search-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) where n = number of nodes  
**Kompleksitas Ruang:** O(n) for sorted values array

**Algoritma:** Binary Search (pencarian biner)

## Solusi Go

```go
package main

// LeetCode #1382: Balance a Binary Search Tree
// https://leetcode.com/problems/balance-a-binary-search-tree/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{Val: 4},
			},
		},
	}
	result := balanceBST(root)
	fmt.Println(result.Val) // 2

	// Test case 2 - empty
	fmt.Println(balanceBST(nil)) // nil

	// Test case 3
	root2 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	result2 := balanceBST(root2)
	fmt.Println(result2.Val) // 2
}

// Time: O(n) where n = number of nodes
// Space: O(n) for sorted values array
func balanceBST(root *TreeNode) *TreeNode {
	// Inorder traversal to get sorted values
	values := []int{}
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		values = append(values, node.Val)
		inorder(node.Right)
	}
	inorder(root)

	// Build balanced BST from sorted array
	var build func(int, int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := left + (right-left)/2
		node := &TreeNode{Val: values[mid]}
		node.Left = build(left, mid-1)
		node.Right = build(mid+1, right)
		return node
	}

	return build(0, len(values)-1)
}
```
