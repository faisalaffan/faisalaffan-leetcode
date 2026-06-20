# 2074 — Reverse Nodes In Even Length Groups

## Deskripsi

**Soal:** [2074. Reverse Nodes In Even Length Groups](https://leetcode.com/problems/reverse-nodes-in-even-length-groups/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func reverseEvenLengthGroups(head *ListNode) *ListNode`

## Solusi Go

```go
package main

// LeetCode #2074: Reverse Nodes in Even Length Groups
// https://leetcode.com/problems/reverse-nodes-in-even-length-groups/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	prev := dummy
	groupSize := 1

	for prev.Next != nil {
		// Count nodes in this group
		count := 0
		curr := prev.Next
		for curr != nil && count < groupSize {
			curr = curr.Next
			count++
		}

		if count%2 == 0 {
			// Reverse this group
			first := prev.Next
			curr = first.Next
			for i := 1; i < count; i++ {
				next := curr.Next
				curr.Next = prev.Next
				prev.Next = curr
				first.Next = next
				curr = next
			}
			prev = first
		} else {
			// Move prev to last node of this group
			for i := 0; i < count; i++ {
				prev = prev.Next
			}
		}
		groupSize++
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
	// Test case 1
	head1 := &ListNode{5, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{9, &ListNode{1, &ListNode{7, &ListNode{3, &ListNode{8, &ListNode{4, nil}}}}}}}}}}
	fmt.Print("Test 1: ")
	printList(reverseEvenLengthGroups(head1))
	// Expected: [5,2,6,3,9,1,4,8,3,7]

	// Test case 2
	head2 := &ListNode{1, &ListNode{1, &ListNode{0, &ListNode{6, &ListNode{5, nil}}}}}
	fmt.Print("Test 2: ")
	printList(reverseEvenLengthGroups(head2))
	// Expected: [1,1,0,6,5]

	// Test case 3
	head3 := &ListNode{1, nil}
	fmt.Print("Test 3: ")
	printList(reverseEvenLengthGroups(head3))
	// Expected: [1]
}
```
