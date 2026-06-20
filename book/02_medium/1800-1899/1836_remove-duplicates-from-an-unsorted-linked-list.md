# 1836 — Remove Duplicates From An Unsorted Linked List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func deleteDuplicatesUnsorted(head *ListNode) *ListNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1836: Remove Duplicates From an Unsorted Linked List
// https://leetcode.com/problems/remove-duplicates-from-an-unsorted-linked-list/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(n)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
	// Count frequencies
  // Membuat map (HashMap) — pencarian O(1)
	count := make(map[int]int)
	curr := head
	for curr != nil {
		count[curr.Val]++
		curr = curr.Next
	}

	// Remove duplicates
	dummy := &ListNode{Next: head}
	prev := dummy
	curr = head
	for curr != nil {
		if count[curr.Val] > 1 {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return dummy.Next
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
	l1 := makeList([]int{1, 2, 3, 2})
	printList(deleteDuplicatesUnsorted(l1)) // Expected: 1 3

	l2 := makeList([]int{2, 1, 1, 2})
	printList(deleteDuplicatesUnsorted(l2)) // Expected: (empty)

	l3 := makeList([]int{3, 2, 2, 1, 3, 2, 4})
	printList(deleteDuplicatesUnsorted(l3)) // Expected: 1 4
}
```
