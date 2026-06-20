# 1721 — Swapping Nodes In A Linked List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func swapNodes(head *ListNode, k int) *ListNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1721: Swapping Nodes in a Linked List
// https://leetcode.com/problems/swapping-nodes-in-a-linked-list/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	// First pass: find kth from beginning
	first := head
	for i := 1; i < k; i++ {
		first = first.Next
	}

	// Two pointer approach for kth from end
	slow := head
	fast := first
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// Swap values
	first.Val, slow.Val = slow.Val, first.Val
	return head
}

func makeList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for i := 1; i < len(vals); i++ {
		curr.Next = &ListNode{Val: vals[i]}
		curr = curr.Next
	}
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	l1 := makeList([]int{1, 2, 3, 4, 5})
	printList(swapNodes(l1, 2)) // Expected: 1 4 3 2 5

	l2 := makeList([]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5})
	printList(swapNodes(l2, 5)) // Expected: 7 9 6 6 8 7 3 0 9 5

	l3 := makeList([]int{1, 2})
	printList(swapNodes(l3, 1)) // Expected: 2 1
}
```
