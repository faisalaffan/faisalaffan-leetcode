# 0206 — Reverse Linked List

## Deskripsi

**Soal:** [0206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func ReverseList(head *ListNode) *ListNode`

## Solusi Go

```go
package main

// LeetCode #206: Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/
// Difficulty: Easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n) | Space: O(1)
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	printList(ReverseList(l1))
	l2 := &ListNode{1, &ListNode{2, nil}}
	printList(ReverseList(l2))
}
```
