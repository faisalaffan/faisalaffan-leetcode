# 0328 — Odd Even Linked List

## Deskripsi

**Soal:** [0328. Odd Even Linked List](https://leetcode.com/problems/odd-even-linked-list/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func oddEvenList(head *ListNode) *ListNode`

## Solusi Go

```go
package main

// LeetCode #328: Odd Even Linked List
// https://leetcode.com/problems/odd-even-linked-list/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd, even := head, head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
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
	// Test case 1: 1->2->3->4->5
	head1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Print("Test 1: ")
	printList(oddEvenList(head1))
	// Expected: 1->3->5->2->4

	// Test case 2: 2->1->3->5->6->4->7
	head2 := &ListNode{2, &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{6, &ListNode{4, &ListNode{7, nil}}}}}}}
	fmt.Print("Test 2: ")
	printList(oddEvenList(head2))
	// Expected: 2->3->6->7->1->5->4

	// Test case 3: Single node
	head3 := &ListNode{1, nil}
	fmt.Print("Test 3: ")
	printList(oddEvenList(head3))
	// Expected: 1
}
```
