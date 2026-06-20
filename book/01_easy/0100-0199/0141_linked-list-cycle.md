# 0141 — Linked List Cycle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func HasCycle(head *ListNode) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #141: Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/
// Difficulty: Easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n) | Space: O(1)
func HasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	n1 := &ListNode{3, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{0, nil}
	n4 := &ListNode{-4, nil}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2
	fmt.Println(HasCycle(n1))

	single := &ListNode{1, nil}
	fmt.Println(HasCycle(single))
}
```
