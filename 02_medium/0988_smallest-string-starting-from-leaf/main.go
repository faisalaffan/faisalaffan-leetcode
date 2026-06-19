package main

// LeetCode #988: Smallest String Starting From Leaf
// https://leetcode.com/problems/smallest-string-starting-from-leaf/
// Difficulty: Medium
//
// Approach: DFS backtracking from root to leaf, compare strings lexicographically
// Time: O(n * h) worst-case where h is tree height
// Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example: [0,1,2,3,4,3,4]
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
			Right: &TreeNode{Val: 4, Left: nil, Right: nil},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
			Right: &TreeNode{Val: 4, Left: nil, Right: nil},
		},
	}
	fmt.Println(smallestFromLeaf(root)) // "dba"

	// Example: [25,1,null,0,0,1,null,null,null,0]
	root2 := &TreeNode{
		Val: 25,
		Left: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0, Left: nil, Right: nil},
			Right: &TreeNode{Val: 0, Left: &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 0, Left: nil, Right: nil}}, Right: nil},
		},
		Right: nil,
	}
	fmt.Println(smallestFromLeaf(root2)) // "abz"
}

func smallestFromLeaf(root *TreeNode) string {
	result := ""
	dfs(root, []byte{}, &result)
	return result
}

func dfs(node *TreeNode, buf []byte, result *string) {
	if node == nil {
		return
	}

	buf = append(buf, byte('a'+node.Val))

	if node.Left == nil && node.Right == nil {
		s := reverse(string(buf))
		if *result == "" || s < *result {
			*result = s
		}
		return
	}

	if node.Left != nil {
		dfs(node.Left, buf, result)
	}
	if node.Right != nil {
		dfs(node.Right, buf, result)
	}
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
