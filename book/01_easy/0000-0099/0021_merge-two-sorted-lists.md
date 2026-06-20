# 0021 — Merge Two Sorted Lists

## Deskripsi

**Soal:** [0021. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n+m)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode`

## Solusi Go

```go
package main

// LeetCode #21: Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/
// Difficulty: Easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n+m) | Space: O(1)
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}
	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	printList(MergeTwoLists(l1, l2))

	l3 := &ListNode{}
	l4 := &ListNode{}
	printList(MergeTwoLists(l3, l4))
}
```
