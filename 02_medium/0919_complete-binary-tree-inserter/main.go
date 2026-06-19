package main

// LeetCode #919: Complete Binary Tree Inserter
// https://leetcode.com/problems/complete-binary-tree-inserter/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	root *TreeNode
	q    []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	q := []*TreeNode{root}
	for {
		node := q[0]
		if node.Left != nil {
			q = append(q, node.Left)
		} else {
			break
		}
		if node.Right != nil {
			q = append(q, node.Right)
		} else {
			break
		}
		q = q[1:]
	}
	return CBTInserter{root, q}
}

func (this *CBTInserter) Insert(val int) int {
	parent := this.q[0]
	node := &TreeNode{Val: val}
	if parent.Left == nil {
		parent.Left = node
	} else {
		parent.Right = node
		this.q = this.q[1:]
	}
	this.q = append(this.q, node)
	return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

func main() {
	root := &TreeNode{Val: 1}
	obj := Constructor(root)
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.Get_root().Val)
	fmt.Println(obj.Insert(3))
	fmt.Println(obj.Insert(4))
	fmt.Println(obj.Get_root().Val)
}
