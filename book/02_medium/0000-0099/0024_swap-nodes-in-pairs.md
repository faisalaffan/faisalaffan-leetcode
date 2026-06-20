# 0024 — Swap Nodes In Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan linked list. Tugasmu traversing atau memanipulasi list.

**Cara berpikir:** Traverse dari head. Fast/slow pointer untuk deteksi siklus/cari tengah. Dummy node mempermudah operasi di head.

**Fungsi Solusi:** `func swapPairs(head *ListNode) *ListNode`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #24: Swap Nodes in Pairs
// https://leetcode.com/problems/swap-nodes-in-pairs/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil && head.Next != nil {
		first := head
		second := head.Next

		prev.Next = second
		first.Next = second.Next
		second.Next = first

		prev = first
		head = first.Next
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
	// Test case 1: [1,2,3,4] -> [2,1,4,3]
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	result := swapPairs(head)
	printList(result)

	// Test case 2: [] -> []
	result = swapPairs(nil)
	printList(result)

	// Test case 3: [1] -> [1]
	head = &ListNode{1, nil}
	result = swapPairs(head)
	printList(result)
}

// Time: O(n) | Space: O(1)
```
