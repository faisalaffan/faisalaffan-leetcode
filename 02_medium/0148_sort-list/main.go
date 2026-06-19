package main

// LeetCode #148: Sort List
// https://leetcode.com/problems/sort-list/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Find middle
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil

	left := sortList(head)
	right := sortList(slow)

	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}

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
	// Test case 1: [4,2,1,3] -> [1,2,3,4]
	head := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	result := sortList(head)
	printList(result)

	// Test case 2: [-1,5,3,4,0] -> [-1,0,3,4,5]
	head = &ListNode{-1, &ListNode{5, &ListNode{3, &ListNode{4, &ListNode{0, nil}}}}}
	result = sortList(head)
	printList(result)

	// Test case 3: [] -> []
	result = sortList(nil)
	printList(result)
}

// Time: O(n log n) | Space: O(log n)
