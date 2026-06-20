# 0203 — Remove Linked List Elements

## Deskripsi

**Soal:** [0203. Remove Linked List Elements](https://leetcode.com/problems/remove-linked-list-elements/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func RemoveElements(head *ListNode, val int) *ListNode`

## Solusi Go

```go
package main

// LeetCode #203: Remove Linked List Elements
// https://leetcode.com/problems/remove-linked-list-elements/
// Difficulty: Easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n) | Space: O(1)
func RemoveElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	curr := dummy
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
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
	l1 := &ListNode{1, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}
	printList(RemoveElements(l1, 6))
	l2 := &ListNode{7, &ListNode{7, &ListNode{7, &ListNode{7, nil}}}}
	printList(RemoveElements(l2, 7))
}
```
