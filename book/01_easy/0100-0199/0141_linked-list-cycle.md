# 0141 — Linked List Cycle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan linked list. Tugasmu traversing atau memanipulasi list.

**Cara berpikir:** Traverse dari head. Fast/slow pointer untuk deteksi siklus/cari tengah. Dummy node mempermudah operasi di head.

**Fungsi Solusi:** `func HasCycle(head *ListNode) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #141: Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/
// Difficulty: Easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n) | Space: O(1)
func HasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	n1 := &ListNode{3, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{0, nil}
	n4 := &ListNode{-4, nil}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2
	fmt.Println(HasCycle(n1))

	single := &ListNode{1, nil}
	fmt.Println(HasCycle(single))
}
```
