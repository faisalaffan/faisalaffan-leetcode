package main

// LeetCode #2181: Merge Nodes in Between Zeros
// https://leetcode.com/problems/merge-nodes-in-between-zeros/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	sum := 0

	for head != nil {
		if head.Val == 0 && sum > 0 {
			curr.Next = &ListNode{Val: sum}
			curr = curr.Next
			sum = 0
		}
		sum += head.Val
		head = head.Next
	}

	return dummy.Next
}

func main() {
	// Test case 1: Example [0,3,1,0,4,5,2,0]
	head1 := &ListNode{0, &ListNode{3, &ListNode{1, &ListNode{0, &ListNode{4, &ListNode{5, &ListNode{2, &ListNode{0, nil}}}}}}}}
	for n := mergeNodes(head1); n != nil; n = n.Next {
		fmt.Printf("%d ", n.Val)
	}
	fmt.Println()
	// Expected: 4 11

	// Test case 2: [0,1,0,3,0,2,2,0]
	head2 := &ListNode{0, &ListNode{1, &ListNode{0, &ListNode{3, &ListNode{0, &ListNode{2, &ListNode{2, &ListNode{0, nil}}}}}}}}
	for n := mergeNodes(head2); n != nil; n = n.Next {
		fmt.Printf("%d ", n.Val)
	}
	fmt.Println()
	// Expected: 1 3 4
}
