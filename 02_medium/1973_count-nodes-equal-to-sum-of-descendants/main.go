package main

// LeetCode #1973: Count Nodes Equal to Sum of Descendants
// https://leetcode.com/problems/count-nodes-equal-to-sum-of-descendants/
// Difficulty: Medium [Paid]

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 10,
		Left:  &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{Val: 5}}
	fmt.Println(EqualToDescendants(root))

	root2 := &TreeNode{Val: 2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 1}}
	fmt.Println(EqualToDescendants(root2))
}

// Time: O(n), Space: O(h) where h is tree height
func EqualToDescendants(root *TreeNode) int {
	count := 0
	dfs(root, &count)
	return count
}

func dfs(node *TreeNode, count *int) int64 {
	if node == nil {
		return 0
	}
	leftSum := dfs(node.Left, count)
	rightSum := dfs(node.Right, count)
	if leftSum+rightSum == int64(node.Val) {
		*count++
	}
	return leftSum + rightSum + int64(node.Val)
}
