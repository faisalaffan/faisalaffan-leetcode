# 0142 — Linked List Cycle Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan linked list. Tugasmu traversing atau memanipulasi list.

**Cara berpikir:** Traverse dari head. Fast/slow pointer untuk deteksi siklus/cari tengah. Dummy node mempermudah operasi di head.

**Fungsi Solusi:** `func detectCycle(head *ListNode) *ListNode`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #142: Linked List Cycle II
// https://leetcode.com/problems/linked-list-cycle-ii/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}

	return nil
}

func main() {
	// Test case 1: [3,2,0,-4], pos=1
	n0 := &ListNode{Val: 3}
	n1 := &ListNode{Val: 2}
	n2 := &ListNode{Val: 0}
	n3 := &ListNode{Val: -4}
	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n1 // cycle
	result := detectCycle(n0)
	if result != nil {
		fmt.Println(result.Val) // 2
	}

	// Test case 2: no cycle
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	result = detectCycle(head)
	fmt.Println(result) // nil

	// Test case 3: single node, no cycle
	result = detectCycle(&ListNode{Val: 1})
	fmt.Println(result) // nil
}

// Time: O(n) | Space: O(1)
```
