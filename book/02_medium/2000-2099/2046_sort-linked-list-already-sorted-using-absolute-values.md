# 2046 — Sort Linked List Already Sorted Using Absolute Values

## Deskripsi

**Soal:** [2046. Sort Linked List Already Sorted Using Absolute Values](https://leetcode.com/problems/sort-linked-list-already-sorted-using-absolute-values/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func sortLinkedList(head *ListNode) *ListNode`

## Solusi Go

```go
package main

// LeetCode #2046: Sort Linked List Already Sorted Using Absolute Values
// https://leetcode.com/problems/sort-linked-list-already-sorted-using-absolute-values/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Separate into negative and non-negative lists
	var negHead, negTail *ListNode
	var posHead, posTail *ListNode

	curr := head
	for curr != nil {
		if curr.Val < 0 {
			// For absolute-sorted list, negatives are in reverse order
			// Prepend to maintain reverse
			next := curr.Next
			if negHead == nil {
				negHead = curr
				negTail = curr
				curr.Next = nil
			} else {
				curr.Next = negHead
				negHead = curr
			}
			curr = next
		} else {
			if posHead == nil {
				posHead = curr
				posTail = curr
			} else {
				posTail.Next = curr
				posTail = curr
			}
			curr = curr.Next
		}
	}

	if posTail != nil {
		posTail.Next = nil
	}

	// Combine: negatives (reversed order) followed by positives
	if negHead != nil {
		if posHead != nil {
			negTail.Next = posHead
		}
		return negHead
	}
	return posHead
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1
	head1 := &ListNode{1, &ListNode{-2, &ListNode{-3, &ListNode{4, &ListNode{-5, nil}}}}}
	fmt.Print("Test 1: ")
	printList(sortLinkedList(head1))
	// Expected: -5 -3 -2 1 4

	// Test case 2
	head2 := &ListNode{-1, &ListNode{-2, &ListNode{-3, nil}}}
	fmt.Print("Test 2: ")
	printList(sortLinkedList(head2))
	// Expected: -3 -2 -1

	// Test case 3
	head3 := &ListNode{0, &ListNode{1, &ListNode{2, nil}}}
	fmt.Print("Test 3: ")
	printList(sortLinkedList(head3))
	// Expected: 0 1 2
}
```
