# 0369 — Plus One Linked List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func plusOne(head *ListNode) *ListNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #369: Plus One Linked List
// https://leetcode.com/problems/plus-one-linked-list/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func plusOne(head *ListNode) *ListNode {
	// Sentinel node
	sentinel := &ListNode{Next: head}
	notNine := sentinel

	// Find rightmost node that is not 9
	for node := head; node != nil; node = node.Next {
		if node.Val != 9 {
			notNine = node
		}
	}

	// Increment rightmost non-9 node
	notNine.Val++
	// Set all following 9s to 0
	for node := notNine.Next; node != nil; node = node.Next {
		node.Val = 0
	}

	if sentinel.Val == 1 {
		return sentinel
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
	// Test case 1: 1->2->3
	head1 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	fmt.Print("Test 1: ")
	printList(plusOne(head1))
	// Expected: 1->2->4

	// Test case 2: 9->9->9
	head2 := &ListNode{9, &ListNode{9, &ListNode{9, nil}}}
	fmt.Print("Test 2: ")
	printList(plusOne(head2))
	// Expected: 1->0->0->0

	// Test case 3: 9
	head3 := &ListNode{9, nil}
	fmt.Print("Test 3: ")
	printList(plusOne(head3))
	// Expected: 1->0
}
```
