package main

// LeetCode #2583: Kth Largest Sum in a Binary Tree
// https://leetcode.com/problems/kth-largest-sum-in-a-binary-tree/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type MinHeap []int64

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int64)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	if root == nil {
		return -1
	}

	queue := []*TreeNode{root}
	h := &MinHeap{}
	heap.Init(h)

	for len(queue) > 0 {
		size := len(queue)
		var sum int64
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += int64(node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		heap.Push(h, sum)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	if h.Len() < k {
		return -1
	}
	return (*h)[0]
}

func main() {
	// Test case 1: [5,8,9,2,1,3,7,4,6]
	root1 := &TreeNode{Val: 5}
	root1.Left = &TreeNode{Val: 8}
	root1.Right = &TreeNode{Val: 9}
	root1.Left.Left = &TreeNode{Val: 2}
	root1.Left.Right = &TreeNode{Val: 1}
	root1.Right.Left = &TreeNode{Val: 3}
	root1.Right.Right = &TreeNode{Val: 7}
	root1.Left.Left.Left = &TreeNode{Val: 4}
	root1.Left.Left.Right = &TreeNode{Val: 6}
	fmt.Println("Test 1:", kthLargestLevelSum(root1, 2))

	// Test case 2: single node
	root2 := &TreeNode{Val: 1}
	fmt.Println("Test 2:", kthLargestLevelSum(root2, 1))

	// Test case 3: k larger than levels
	root3 := &TreeNode{Val: 1}
	root3.Left = &TreeNode{Val: 2}
	fmt.Println("Test 3:", kthLargestLevelSum(root3, 3))
}
