# 0002 — Add Two Numbers

## Deskripsi

**Soal:** [0002. Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(max(m,n))  
**Kompleksitas Ruang:** O(max(m,n))

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode`

## Solusi Go

```go
package main

// LeetCode #2: Add Two Numbers
// https://leetcode.com/problems/add-two-numbers/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
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
	// Test case 1: l1 = [2,4,3], l2 = [5,6,4] -> [7,0,8]
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := addTwoNumbers(l1, l2)
	printList(result)

	// Test case 2: l1 = [0], l2 = [0] -> [0]
	l1 = &ListNode{0, nil}
	l2 = &ListNode{0, nil}
	result = addTwoNumbers(l1, l2)
	printList(result)

	// Test case 3: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9] -> [8,9,9,9,0,0,0,1]
	l1 = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	l2 = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	result = addTwoNumbers(l1, l2)
	printList(result)
}

// Time: O(max(m,n)) | Space: O(max(m,n))
```
