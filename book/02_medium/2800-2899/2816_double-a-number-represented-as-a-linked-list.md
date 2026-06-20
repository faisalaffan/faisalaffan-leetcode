# 2816 — Double A Number Represented As A Linked List

## Deskripsi

**Soal:** [2816. Double A Number Represented As A Linked List](https://leetcode.com/problems/double-a-number-represented-as-a-linked-list/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func DoubleANumberRepresentedAsALinkedList(head *ListNode) *ListNode`

## Solusi Go

```go
package main

// LeetCode #2816: Double a Number Represented as a Linked List
// https://leetcode.com/problems/double-a-number-represented-as-a-linked-list/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func DoubleANumberRepresentedAsALinkedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// Reverse the list
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	// Double
	carry := 0
	cur = prev
	var newHead *ListNode
	for cur != nil {
		val := cur.Val*2 + carry
		carry = val / 10
		cur.Val = val % 10
		newHead = cur
		cur = cur.Next
	}
	if carry > 0 {
		newHead = &ListNode{Val: carry, Next: newHead}
	}

	// Reverse back
	prev = nil
	cur = newHead
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func printList2(head *ListNode) {
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
	head := &ListNode{1, &ListNode{8, &ListNode{9, nil}}}
	printList2(DoubleANumberRepresentedAsALinkedList(head))

	head2 := &ListNode{9, &ListNode{9, &ListNode{9, nil}}}
	printList2(DoubleANumberRepresentedAsALinkedList(head2))
}
```
