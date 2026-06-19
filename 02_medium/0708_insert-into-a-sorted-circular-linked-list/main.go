package main

// LeetCode #708: Insert into a Sorted Circular Linked List
// https://leetcode.com/problems/insert-into-a-sorted-circular-linked-list/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	head := &CNode{Val: 3}
	head.Next = &CNode{Val: 4}
	head.Next.Next = &CNode{Val: 1}
	head.Next.Next.Next = head

	result := insert(head, 2)
	fmt.Println(result.Val)
}

type CNode struct {
	Val  int
	Next *CNode
}

func insert(head *CNode, insertVal int) *CNode {
	newNode := &CNode{Val: insertVal}
	if head == nil {
		newNode.Next = newNode
		return newNode
	}

	curr := head
	for curr.Next != head {
		if curr.Val <= insertVal && insertVal <= curr.Next.Val {
			break
		}
		if curr.Val > curr.Next.Val {
			if insertVal >= curr.Val || insertVal <= curr.Next.Val {
				break
			}
		}
		curr = curr.Next
	}

	newNode.Next = curr.Next
	curr.Next = newNode
	return head
}
