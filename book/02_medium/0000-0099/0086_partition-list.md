# 0086 — Partition List

## Deskripsi

**Soal:** [0086. Partition List](https://leetcode.com/problems/partition-list/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func partition(head *ListNode, x int) *ListNode`

## Solusi Go

```go
package main

// LeetCode #86: Partition List
// https://leetcode.com/problems/partition-list/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	lessHead := &ListNode{}
	greaterHead := &ListNode{}
	less, greater := lessHead, greaterHead

	for head != nil {
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			greater.Next = head
			greater = greater.Next
		}
		head = head.Next
	}

	greater.Next = nil
	less.Next = greaterHead.Next

	return lessHead.Next
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
	// Test case 1: [1,4,3,2,5,2], x=3 -> [1,2,2,4,3,5]
	head := &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}
	result := partition(head, 3)
	printList(result)

	// Test case 2: [2,1], x=2 -> [1,2]
	head = &ListNode{2, &ListNode{1, nil}}
	result = partition(head, 2)
	printList(result)
}

// Time: O(n) | Space: O(1)
```
