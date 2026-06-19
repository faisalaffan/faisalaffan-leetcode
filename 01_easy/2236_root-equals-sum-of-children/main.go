package main

// LeetCode #2236: Root Equals Sum of Children
// https://leetcode.com/problems/root-equals-sum-of-children/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{10, &TreeNode{4, nil, nil}, &TreeNode{6, nil, nil}}
	fmt.Println(RootEqualsSumOfChildren(root)) // true

	root2 := &TreeNode{5, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}}
	fmt.Println(RootEqualsSumOfChildren(root2)) // false
}

// Time: O(1), Space: O(1)
func RootEqualsSumOfChildren(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}
