# 0082 — Remove Duplicates From Sorted List Ii

## Deskripsi

**Soal:** [0082. Remove Duplicates From Sorted List Ii](https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func deleteDuplicates(head *ListNode) *ListNode`

## Solusi Go

```go
package main

// LeetCode #82: Remove Duplicates from Sorted List II
// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			prev.Next = head.Next
		} else {
			prev = head
		}
		head = head.Next
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
	// Test case 1: [1,2,3,3,4,4,5] -> [1,2,5]
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}
	result := deleteDuplicates(head)
	printList(result)

	// Test case 2: [1,1,1,2,3] -> [2,3]
	head = &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}}
	result = deleteDuplicates(head)
	printList(result)

	// Test case 3: [1,1] -> []
	head = &ListNode{1, &ListNode{1, nil}}
	result = deleteDuplicates(head)
	printList(result)
}

// Time: O(n) | Space: O(1)
```
