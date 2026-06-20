# 0092 — Reverse Linked List Ii

## Deskripsi

**Soal:** [0092. Reverse Linked List Ii](https://leetcode.com/problems/reverse-linked-list-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func reverseBetween(head *ListNode, left int, right int) *ListNode`

## Solusi Go

```go
package main

// LeetCode #92: Reverse Linked List II
// https://leetcode.com/problems/reverse-linked-list-ii/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	// Move to position left
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	curr := prev.Next
	var next *ListNode

	// Reverse between left and right
	for i := 0; i < right-left; i++ {
		next = curr.Next
		curr.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
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
	// Test case 1: [1,2,3,4,5], left=2, right=4 -> [1,4,3,2,5]
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result := reverseBetween(head, 2, 4)
	printList(result)

	// Test case 2: [5], left=1, right=1 -> [5]
	head = &ListNode{5, nil}
	result = reverseBetween(head, 1, 1)
	printList(result)

	// Test case 3: [1,2,3,4,5], left=1, right=5 -> [5,4,3,2,1]
	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result = reverseBetween(head, 1, 5)
	printList(result)
}

// Time: O(n) | Space: O(1)
```
