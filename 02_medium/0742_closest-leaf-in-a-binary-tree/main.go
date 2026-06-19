package main

// LeetCode #742: Closest Leaf in a Binary Tree
// https://leetcode.com/problems/closest-leaf-in-a-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	root := &BTNode{Val: 1}
	root.Left = &BTNode{Val: 3}
	root.Right = &BTNode{Val: 2}
	fmt.Println(findClosestLeaf(root, 1))
}

type BTNode struct {
	Val   int
	Left  *BTNode
	Right *BTNode
}

func findClosestLeaf(root *BTNode, k int) int {
	graph := make(map[int][]int)
	leaves := make(map[int]bool)
	visited := make(map[int]bool)

	var buildGraph func(node *BTNode)
	buildGraph = func(node *BTNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			leaves[node.Val] = true
		}
		if node.Left != nil {
			graph[node.Val] = append(graph[node.Val], node.Left.Val)
			graph[node.Left.Val] = append(graph[node.Left.Val], node.Val)
			buildGraph(node.Left)
		}
		if node.Right != nil {
			graph[node.Val] = append(graph[node.Val], node.Right.Val)
			graph[node.Right.Val] = append(graph[node.Right.Val], node.Val)
			buildGraph(node.Right)
		}
	}

	buildGraph(root)

	queue := []int{k}
	visited[k] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if leaves[node] {
			return node
		}
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return -1
}
