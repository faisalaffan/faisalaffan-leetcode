package main

// LeetCode #536: Construct Binary Tree from String
// https://leetcode.com/problems/construct-binary-tree-from-string/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := Str2tree("4(2(3)(1))(6(5))")
	printTree(root)
	fmt.Println()
}

func Str2tree(s string) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	return buildTree(s, 0)
}

func buildTree(s string, idx int) *TreeNode {
	if idx >= len(s) || s[idx] == ')' {
		return nil
	}

	// Parse number
	start := idx
	for idx < len(s) && (s[idx] == '-' || (s[idx] >= '0' && s[idx] <= '9')) {
		idx++
	}
	val, _ := strconv.Atoi(s[start:idx])
	node := &TreeNode{Val: val}

	// Parse left child
	if idx < len(s) && s[idx] == '(' {
		node.Left = buildTree(s, idx+1)
		// Skip to matching close paren
		depth := 1
		idx++
		for idx < len(s) && depth > 0 {
			if s[idx] == '(' {
				depth++
			} else if s[idx] == ')' {
				depth--
			}
			idx++
		}
	}

	// Parse right child
	if idx < len(s) && s[idx] == '(' {
		node.Right = buildTree(s, idx+1)
		depth := 1
		idx++
		for idx < len(s) && depth > 0 {
			if s[idx] == '(' {
				depth++
			} else if s[idx] == ')' {
				depth--
			}
			idx++
		}
	}

	return node
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printTree(root.Left)
	printTree(root.Right)
}
