# 0234 — Palindrome Linked List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func IsPalindrome(head *ListNode) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #234: Palindrome Linked List
// https://leetcode.com/problems/palindrome-linked-list/
// Difficulty: Easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n) | Space: O(1)
func IsPalindrome(head *ListNode) bool {
	reverse := func(head *ListNode) *ListNode {
		var prev *ListNode
		for head != nil {
			next := head.Next
			head.Next = prev
			prev = head
			head = next
		}
		return prev
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := reverse(slow)
	first := head
	for second != nil {
		if first.Val != second.Val {
			return false
		}
		first = first.Next
		second = second.Next
	}
	return true
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	fmt.Println(IsPalindrome(l1))
	l2 := &ListNode{1, &ListNode{2, nil}}
	fmt.Println(IsPalindrome(l2))
}
```
