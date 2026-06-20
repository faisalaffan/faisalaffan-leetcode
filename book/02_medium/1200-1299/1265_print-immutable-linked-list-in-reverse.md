# 1265 — Print Immutable Linked List In Reverse

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan linked list. Tugasmu traversing atau memanipulasi list.

**Cara berpikir:** Traverse dari head. Fast/slow pointer untuk deteksi siklus/cari tengah. Dummy node mempermudah operasi di head.

**Fungsi Solusi:** `func printLinkedListInReverse(head *ImmutableListNode) `

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n) for recursive stack


## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1265: Print Immutable Linked List in Reverse
// https://leetcode.com/problems/print-immutable-linked-list-in-reverse/
// Difficulty: Medium [Paid]

// Print linked list in reverse using recursion (or stack).
// Immutable means we can't modify the list.

// Time: O(n)
// Space: O(n) for recursive stack

type ImmutableListNode struct {
	val  int
	next *ImmutableListNode
}

func (node *ImmutableListNode) getValue() int {
	return node.val
}

func (node *ImmutableListNode) getNext() *ImmutableListNode {
	return node.next
}

func printLinkedListInReverse(head *ImmutableListNode) {
	if head == nil {
		return
	}
	printLinkedListInReverse(head.getNext())
	fmt.Printf("%d ", head.getValue())
}

func main() {
	head := &ImmutableListNode{1, &ImmutableListNode{2, &ImmutableListNode{3, nil}}}
	fmt.Printf("Reversed: ")
	printLinkedListInReverse(head)
	fmt.Println()
}
```
