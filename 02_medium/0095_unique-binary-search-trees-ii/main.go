package main

// LeetCode #95: Unique Binary Search Trees II
// https://leetcode.com/problems/unique-binary-search-trees-ii/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return build(1, n)
}

func build(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	result := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftTrees := build(start, i-1)
		rightTrees := build(i+1, end)
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				result = append(result, &TreeNode{Val: i, Left: left, Right: right})
			}
		}
	}
	return result
}

func printTreePreorder(root *TreeNode) {
	if root == nil {
		fmt.Print("null ")
		return
	}
	fmt.Printf("%d ", root.Val)
	printTreePreorder(root.Left)
	printTreePreorder(root.Right)
}

func main() {
	// Test case 1
	trees := generateTrees(3)
	fmt.Println(len(trees)) // 5
	for _, t := range trees {
		printTreePreorder(t)
		fmt.Println()
	}

	// Test case 2
	trees = generateTrees(1)
	fmt.Println(len(trees)) // 1
	for _, t := range trees {
		printTreePreorder(t)
		fmt.Println()
	}
}

// Time: O(4^n / n^(3/2)) | Space: O(4^n / n^(3/2))
