# 3263 — Convert Doubly Linked List To Array I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ConvertDoublyLinkedListToArrayI(head *Node) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).


## 💻 Solusi Go

```go
package main

// LeetCode #3263: Convert Doubly Linked List to Array I
// https://leetcode.com/problems/convert-doubly-linked-list-to-array-i/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	// 1 <-> 2 <-> 3
	head := &Node{Val: 1}
	head.Next = &Node{Val: 2, Prev: head}
	head.Next.Next = &Node{Val: 3, Prev: head.Next}
	fmt.Println(ConvertDoublyLinkedListToArrayI(head))
}

// Node represents a doubly-linked list node.
type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

// ConvertDoublyLinkedListToArrayI converts a doubly linked list to an integer array.
// Time: O(n). Space: O(n).
func ConvertDoublyLinkedListToArrayI(head *Node) []int {
	result := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		result = append(result, cur.Val)
	}
	return result
}
```
