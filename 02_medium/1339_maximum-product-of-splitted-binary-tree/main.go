package main

// LeetCode #1339: Maximum Product of Splitted Binary Tree
// https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/
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
		Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}},
	}
	fmt.Println(maxProduct(root)) // 110

	// Test case 2
	root2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}}
	fmt.Println(maxProduct(root2)) // 90

	// Test case 3
	root3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}}
	fmt.Println(maxProduct(root3)) // 1
}

const mod = 1_000_000_007

// Time: O(n) where n is number of nodes
// Space: O(h) for recursion stack
func maxProduct(root *TreeNode) int {
	var totalSum int
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := node.Val + dfs(node.Left) + dfs(node.Right)
		return sum
	}
	totalSum = dfs(root)

	maxProd := 0
	var findMax func(*TreeNode) int
	findMax = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := findMax(node.Left)
		rightSum := findMax(node.Right)
		subtreeSum := node.Val + leftSum + rightSum

		if subtreeSum != totalSum {
			prod := subtreeSum * (totalSum - subtreeSum)
			if prod > maxProd {
				maxProd = prod
			}
		}

		return subtreeSum
	}
	findMax(root)

	return maxProd % mod
}
