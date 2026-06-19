package main

// LeetCode #143: Reorder List
// https://leetcode.com/problems/reorder-list/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Find middle
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse second half
	var prev *ListNode
	curr := slow.Next
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	slow.Next = nil

	// Merge two halves
	first, second := head, prev
	for first != nil && second != nil {
		firstNext := first.Next
		secondNext := second.Next
		first.Next = second
		second.Next = firstNext
		first = firstNext
		second = secondNext
	}
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1: [1,2,3,4] -> [1,4,2,3]
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	reorderList(head)
	printList(head)

	// Test case 2: [1,2,3,4,5] -> [1,5,2,4,3]
	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	reorderList(head)
	printList(head)

	// Test case 3: [1] -> [1]
	head = &ListNode{1, nil}
	reorderList(head)
	printList(head)
}

// Time: O(n) | Space: O(1)
