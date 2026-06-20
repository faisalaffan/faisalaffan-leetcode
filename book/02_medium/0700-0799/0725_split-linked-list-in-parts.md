# 0725 — Split Linked List In Parts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func splitListToParts(head *ListNode, k int) []*ListNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #725: Split Linked List in Parts
// https://leetcode.com/problems/split-linked-list-in-parts/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}

	parts := splitListToParts(head, 3)
	for _, p := range parts {
		if p == nil {
			fmt.Println("nil")
		} else {
			fmt.Println(p.Val)
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	length := 0
	for curr := head; curr != nil; curr = curr.Next {
		length++
	}

	partSize := length / k
	extra := length % k

	result := make([]*ListNode, k)
	curr := head

	for i := 0; i < k && curr != nil; i++ {
		result[i] = curr
		size := partSize
		if i < extra {
			size++
		}

		for j := 0; j < size-1; j++ {
			curr = curr.Next
		}

		next := curr.Next
		curr.Next = nil
		curr = next
	}

	return result
}
```
