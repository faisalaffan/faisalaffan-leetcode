# 2674 — Split A Circular Linked List

## Deskripsi

**Soal:** [2674. Split A Circular Linked List](https://leetcode.com/problems/split-a-circular-linked-list/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Dynamic Programming (DP), LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func splitCircularLinkedList(head *ListNode) []*ListNode`

## Solusi Go

```go
package main

// LeetCode #2674: Split a Circular Linked List
// https://leetcode.com/problems/split-a-circular-linked-list/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitCircularLinkedList(head *ListNode) []*ListNode {
	if head == nil || head.Next == nil {
		return []*ListNode{head, nil}
	}

	// Find the midpoint using slow/fast pointers
	slow, fast := head, head
	for fast.Next != head && fast.Next.Next != head {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// First half: head to slow
	list1 := head

	// Second half: slow.Next to end, make it circular
	list2 := slow.Next

	// Make first list circular
	slow.Next = head

	// Find end of second list and make it circular
	curr := list2
	for curr.Next != head {
		curr = curr.Next
	}
	curr.Next = list2

	return []*ListNode{list1, list2}
}

func printCircular(head *ListNode, count int) {
	curr := head
	for i := 0; i < count && curr != nil; i++ {
		fmt.Print(curr.Val, " ")
		curr = curr.Next
	}
}

func main() {
	// Test case 1: [1,2,3,4]
	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 2}
	head1.Next.Next = &ListNode{Val: 3}
	head1.Next.Next.Next = &ListNode{Val: 4}
	head1.Next.Next.Next.Next = head1

	parts := splitCircularLinkedList(head1)
	fmt.Print("Test 1 list1: ")
	printCircular(parts[0], 2)
	fmt.Println()

	fmt.Print("Test 1 list2: ")
	printCircular(parts[1], 2)
	fmt.Println()

	// Test case 2: single node
	head2 := &ListNode{Val: 1}
	head2.Next = head2
	parts2 := splitCircularLinkedList(head2)
	fmt.Println("Test 2 list1:", parts2[0].Val)
	fmt.Println("Test 2 list2:", parts2[1])
}
```
