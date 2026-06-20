# 0143 — Reorder List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func reorderList(head *ListNode) 
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #143: Reorder List
// https://leetcode.com/problems/reorder-list/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Find middle
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse second half
	var prev *ListNode
	curr := slow.Next
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	slow.Next = nil

	// Merge two halves
	first, second := head, prev
	for first != nil && second != nil {
		firstNext := first.Next
		secondNext := second.Next
		first.Next = second
		second.Next = firstNext
		first = firstNext
		second = secondNext
	}
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
	// Test case 1: [1,2,3,4] -> [1,4,2,3]
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	reorderList(head)
	printList(head)

	// Test case 2: [1,2,3,4,5] -> [1,5,2,4,3]
	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	reorderList(head)
	printList(head)

	// Test case 3: [1] -> [1]
	head = &ListNode{1, nil}
	reorderList(head)
	printList(head)
}

// Time: O(n) | Space: O(1)
```
