# 0445 — Add Two Numbers Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan linked list. Tugasmu traversing atau memanipulasi list.

**Cara berpikir:** Traverse dari head. Fast/slow pointer untuk deteksi siklus/cari tengah. Dummy node mempermudah operasi di head.

**Fungsi Solusi:** `func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #445: Add Two Numbers II
// https://leetcode.com/problems/add-two-numbers-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := []int{}
	s2 := []int{}

	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	var head *ListNode
	carry := 0

	for len(s1) > 0 || len(s2) > 0 || carry > 0 {
		sum := carry
		if len(s1) > 0 {
			sum += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			sum += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}

		node := &ListNode{Val: sum % 10, Next: head}
		head = node
		carry = sum / 10
	}

	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1: 7->2->4->3 + 5->6->4 = 7->8->0->7
	l1 := &ListNode{7, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	fmt.Print("Test 1: ")
	printList(addTwoNumbers(l1, l2))
	// Expected: 7 -> 8 -> 0 -> 7

	// Test case 2: 5 + 5 = 1->0
	l3 := &ListNode{5, nil}
	l4 := &ListNode{5, nil}
	fmt.Print("Test 2: ")
	printList(addTwoNumbers(l3, l4))
	// Expected: 1 -> 0

	// Test case 3: 0 + 0 = 0
	l5 := &ListNode{0, nil}
	l6 := &ListNode{0, nil}
	fmt.Print("Test 3: ")
	printList(addTwoNumbers(l5, l6))
	// Expected: 0
}
```
