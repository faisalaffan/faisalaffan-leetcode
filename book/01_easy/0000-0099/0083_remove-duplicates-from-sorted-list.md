# 0083 — Remove Duplicates From Sorted List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan linked list. Tugasmu traversing atau memanipulasi list.

**Cara berpikir:** Traverse dari head. Fast/slow pointer untuk deteksi siklus/cari tengah. Dummy node mempermudah operasi di head.

**Fungsi Solusi:** `func DeleteDuplicates(head *ListNode) *ListNode`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #83: Remove Duplicates from Sorted List
// https://leetcode.com/problems/remove-duplicates-from-sorted-list/
// Difficulty: Easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n) | Space: O(1)
func DeleteDuplicates(head *ListNode) *ListNode {
	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	l1 := &ListNode{1, &ListNode{1, &ListNode{2, nil}}}
	printList(DeleteDuplicates(l1))
	l2 := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
	printList(DeleteDuplicates(l2))
}
```
