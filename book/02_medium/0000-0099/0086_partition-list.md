# 0086 — Partition List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func partition(head *ListNode, x int) *ListNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #86: Partition List
// https://leetcode.com/problems/partition-list/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	lessHead := &ListNode{}
	greaterHead := &ListNode{}
	less, greater := lessHead, greaterHead

	for head != nil {
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			greater.Next = head
			greater = greater.Next
		}
		head = head.Next
	}

	greater.Next = nil
	less.Next = greaterHead.Next

	return lessHead.Next
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
	// Test case 1: [1,4,3,2,5,2], x=3 -> [1,2,2,4,3,5]
	head := &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}
	result := partition(head, 3)
	printList(result)

	// Test case 2: [2,1], x=2 -> [1,2]
	head = &ListNode{2, &ListNode{1, nil}}
	result = partition(head, 2)
	printList(result)
}

// Time: O(n) | Space: O(1)
```
