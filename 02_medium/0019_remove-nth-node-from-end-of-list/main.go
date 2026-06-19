package main

// LeetCode #19: Remove Nth Node From End of List
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
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
	// Test case 1: remove 2nd from end of [1,2,3,4,5] -> [1,2,3,5]
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result := removeNthFromEnd(head, 2)
	printList(result)

	// Test case 2: remove 1st from end of [1] -> []
	head = &ListNode{1, nil}
	result = removeNthFromEnd(head, 1)
	printList(result)

	// Test case 3: remove 1st from end of [1,2] -> [1]
	head = &ListNode{1, &ListNode{2, nil}}
	result = removeNthFromEnd(head, 1)
	printList(result)
}

// Time: O(n) | Space: O(1)
