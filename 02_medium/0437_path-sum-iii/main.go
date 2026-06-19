package main

// LeetCode #437: Path Sum III
// https://leetcode.com/problems/path-sum-iii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	// prefix sum -> count
	prefixSum := map[int]int{0: 1}
	return dfs(root, targetSum, 0, prefixSum)
}

func dfs(node *TreeNode, targetSum, curSum int, prefixSum map[int]int) int {
	if node == nil {
		return 0
	}

	curSum += node.Val
	count := prefixSum[curSum-targetSum]

	prefixSum[curSum]++
	count += dfs(node.Left, targetSum, curSum, prefixSum)
	count += dfs(node.Right, targetSum, curSum, prefixSum)
	prefixSum[curSum]--

	return count
}

func main() {
	// Test case 1
	root1 := &TreeNode{Val: 10}
	root1.Left = &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: -2}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}}
	root1.Right = &TreeNode{Val: -3, Right: &TreeNode{Val: 11}}
	fmt.Println("Test 1:", pathSum(root1, 8))
	// Expected: 3

	// Test case 2: Single node
	root2 := &TreeNode{Val: 1}
	fmt.Println("Test 2:", pathSum(root2, 1))
	// Expected: 1

	// Test case 3: Nil
	fmt.Println("Test 3:", pathSum(nil, 0))
	// Expected: 0
}
