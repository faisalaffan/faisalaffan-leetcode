package main

// LeetCode #237: Delete Node in a Linked List
// https://leetcode.com/problems/delete-node-in-a-linked-list/
// Difficulty: Medium
// Time: O(1), Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	head := &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{9, nil}}}}
	deleteNode(head.Next)
	printList(head)

	head2 := &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{9, nil}}}}
	deleteNode(head2.Next.Next)
	printList(head2)

	head3 := &ListNode{1, &ListNode{2, nil}}
	deleteNode(head3)
	printList(head3)
}
