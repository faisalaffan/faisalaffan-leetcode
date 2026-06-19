package main

// LeetCode #3319: K-th Largest Perfect Subtree Size in Binary Tree
// https://leetcode.com/problems/k-th-largest-perfect-subtree-size-in-binary-tree/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	// Tree: [1,2,3,4,5,6,7]
	root := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7}}}
	fmt.Println(kthLargestPerfectSubtree(root, 1)) // 7
	fmt.Println(kthLargestPerfectSubtree(root, 3)) // 3
	fmt.Println(kthLargestPerfectSubtree(root, 5)) // -1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestPerfectSubtree(root *TreeNode, k int) int {
	var sizes []int
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left < 0 || left != right {
			return -1
		}
		cur := left + right + 1
		sizes = append(sizes, cur)
		return cur
	}
	dfs(root)

	if len(sizes) < k {
		return -1
	}
	sort.Slice(sizes, func(i, j int) bool { return sizes[i] > sizes[j] })
	return sizes[k-1]
}
