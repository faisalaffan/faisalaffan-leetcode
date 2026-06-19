package main

// LeetCode #2807: Insert Greatest Common Divisors in Linked List
// https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func InsertGreatestCommonDivisorsInLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	cur := head
	for cur != nil && cur.Next != nil {
		g := gcd(cur.Val, cur.Next.Val)
		node := &ListNode{Val: g, Next: cur.Next}
		cur.Next = node
		cur = node.Next
	}

	return head
}

func printList(head *ListNode) {
	for head != nil {
		if head.Next != nil {
			fmt.Printf("%d -> ", head.Val)
		} else {
			fmt.Println(head.Val)
		}
		head = head.Next
	}
}

func main() {
	head := &ListNode{18, &ListNode{6, &ListNode{10, &ListNode{3, nil}}}}
	printList(InsertGreatestCommonDivisorsInLinkedList(head))

	head2 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	printList(InsertGreatestCommonDivisorsInLinkedList(head2))
}
