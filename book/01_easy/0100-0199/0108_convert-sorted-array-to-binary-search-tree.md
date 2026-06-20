# 0108 — Convert Sorted Array To Binary Search Tree

## Deskripsi

**Soal:** [0108. Convert Sorted Array To Binary Search Tree](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(log n) (recursion stack)

**Algoritma:** Binary Search (pencarian biner), Stack (tumpukan LIFO)

**Fungsi Solusi:** `func SortedArrayToBST(nums []int) *TreeNode`

## Solusi Go

```go
package main

// LeetCode #108: Convert Sorted Array to Binary Search Tree
// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(log n) (recursion stack)
func SortedArrayToBST(nums []int) *TreeNode {
	var build func(int, int) *TreeNode
	build = func(lo, hi int) *TreeNode {
		if lo > hi {
			return nil
		}
		mid := lo + (hi-lo)/2
		return &TreeNode{
			Val:   nums[mid],
			Left:  build(lo, mid-1),
			Right: build(mid+1, hi),
		}
	}
	return build(0, len(nums)-1)
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Print(root.Val, " ")
	inorder(root.Right)
}

func main() {
	inorder(SortedArrayToBST([]int{-10, -3, 0, 5, 9}))
	fmt.Println()
	inorder(SortedArrayToBST([]int{1, 3}))
	fmt.Println()
}
```
