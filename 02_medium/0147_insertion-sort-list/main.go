package main

// LeetCode #147: Insertion Sort List
// https://leetcode.com/problems/insertion-sort-list/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{}

	for head != nil {
		prev := dummy
		for prev.Next != nil && prev.Next.Val < head.Val {
			prev = prev.Next
		}
		next := head.Next
		head.Next = prev.Next
		prev.Next = head
		head = next
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
	result := insertionSortList(head)
	printList(result)

	// Test case 2: [-1,5,3,4,0] -> [-1,0,3,4,5]
	head = &ListNode{-1, &ListNode{5, &ListNode{3, &ListNode{4, &ListNode{0, nil}}}}}
	result = insertionSortList(head)
	printList(result)

	// Test case 3: [] -> []
	result = insertionSortList(nil)
	printList(result)
}

// Time: O(n^2) | Space: O(1)
