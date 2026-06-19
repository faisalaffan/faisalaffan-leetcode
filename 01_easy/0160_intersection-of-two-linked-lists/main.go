package main

// LeetCode #160: Intersection of Two Linked Lists
// https://leetcode.com/problems/intersection-of-two-linked-lists/
// Difficulty: Easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n+m) | Space: O(1)
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

func main() {
	common := &ListNode{8, &ListNode{4, &ListNode{5, nil}}}
	a := &ListNode{4, &ListNode{1, common}}
	b := &ListNode{5, &ListNode{6, &ListNode{1, common}}}
	fmt.Println(GetIntersectionNode(a, b).Val)
}
