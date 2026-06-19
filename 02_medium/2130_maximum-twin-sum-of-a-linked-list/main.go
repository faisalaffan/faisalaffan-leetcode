package main

// LeetCode #2130: Maximum Twin Sum of a Linked List
// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	// Find middle
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse second half
	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	// Sum pairs
	maxSum := 0
	first, second := head, prev
	for second != nil {
		sum := first.Val + second.Val
		if sum > maxSum {
			maxSum = sum
		}
		first = first.Next
		second = second.Next
	}

	return maxSum
}

func main() {
	// Test case 1
	head1 := &ListNode{5, &ListNode{4, &ListNode{2, &ListNode{1, nil}}}}
	fmt.Println("Test 1:", pairSum(head1))
	// Expected: 6

	// Test case 2
	head2 := &ListNode{4, &ListNode{2, &ListNode{2, &ListNode{3, nil}}}}
	fmt.Println("Test 2:", pairSum(head2))
	// Expected: 7

	// Test case 3
	head3 := &ListNode{1, &ListNode{100000, nil}}
	fmt.Println("Test 3:", pairSum(head3))
	// Expected: 100001
}
