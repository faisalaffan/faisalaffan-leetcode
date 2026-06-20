# 0147 — Insertion Sort List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func insertionSortList(head *ListNode) *ListNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #147: Insertion Sort List
// https://leetcode.com/problems/insertion-sort-list/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{}

	for head != nil {
		prev := dummy
		for prev.Next != nil && prev.Next.Val < head.Val {
			prev = prev.Next
		}
		next := head.Next
		head.Next = prev.Next
		prev.Next = head
		head = next
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
	// Test case 1: [4,2,1,3] -> [1,2,3,4]
	head := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	result := insertionSortList(head)
	printList(result)

	// Test case 2: [-1,5,3,4,0] -> [-1,0,3,4,5]
	head = &ListNode{-1, &ListNode{5, &ListNode{3, &ListNode{4, &ListNode{0, nil}}}}}
	result = insertionSortList(head)
	printList(result)

	// Test case 3: [] -> []
	result = insertionSortList(nil)
	printList(result)
}

// Time: O(n^2) | Space: O(1)
```
