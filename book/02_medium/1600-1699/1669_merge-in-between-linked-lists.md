# 1669 — Merge In Between Linked Lists

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1669: Merge In Between Linked Lists
// https://leetcode.com/problems/merge-in-between-linked-lists/
// Difficulty: Medium
// Time: O(n + m), Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	dummy := &ListNode{Next: list1}
	prev := dummy

	// Move to node just before position a
	for i := 0; i < a; i++ {
		prev = prev.Next
	}

	// Find end of segment to remove (node at position b)
	end := prev
	for i := a; i <= b; i++ {
		end = end.Next
	}

	// Connect prev node to list2
	prev.Next = list2

	// Find tail of list2
	tail := list2
	for tail != nil && tail.Next != nil {
		tail = tail.Next
	}

	// Connect tail of list2 to node after position b
	if tail != nil {
		tail.Next = end.Next
	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1
	l1 := &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}}
	l2 := &ListNode{100, &ListNode{101, &ListNode{102, nil}}}
	result := mergeInBetween(l1, 3, 4, l2)
	fmt.Print("Test 1: ")
	printList(result) // Expected: 0 1 2 100 101 102 5

	// Test case 2
	l1 = &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}
	l2 = &ListNode{100, nil}
	result = mergeInBetween(l1, 1, 2, l2)
	fmt.Print("Test 2: ")
	printList(result) // Expected: 0 100 3

	// Test case 3
	l1 = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	l2 = &ListNode{10, &ListNode{20, nil}}
	result = mergeInBetween(l1, 0, 1, l2)
	fmt.Print("Test 3: ")
	printList(result) // Expected: 10 20 3
}
```
