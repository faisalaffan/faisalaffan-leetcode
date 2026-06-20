# 0725 — Split Linked List In Parts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan linked list. Tugasmu traversing atau memanipulasi list.

**Cara berpikir:** Traverse dari head. Fast/slow pointer untuk deteksi siklus/cari tengah. Dummy node mempermudah operasi di head.

**Fungsi Solusi:** `func splitListToParts(head *ListNode, k int) []*ListNode`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #725: Split Linked List in Parts
// https://leetcode.com/problems/split-linked-list-in-parts/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}

	parts := splitListToParts(head, 3)
	for _, p := range parts {
		if p == nil {
			fmt.Println("nil")
		} else {
			fmt.Println(p.Val)
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	length := 0
	for curr := head; curr != nil; curr = curr.Next {
		length++
	}

	partSize := length / k
	extra := length % k

	result := make([]*ListNode, k)
	curr := head

	for i := 0; i < k && curr != nil; i++ {
		result[i] = curr
		size := partSize
		if i < extra {
			size++
		}

		for j := 0; j < size-1; j++ {
			curr = curr.Next
		}

		next := curr.Next
		curr.Next = nil
		curr = next
	}

	return result
}
```
