# 0160 — Intersection Of Two Linked Lists

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func GetIntersectionNode(headA, headB *ListNode) *ListNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n+m)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #160: Intersection of Two Linked Lists
// https://leetcode.com/problems/intersection-of-two-linked-lists/
// Difficulty: Easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n+m) | Space: O(1)
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

func main() {
	common := &ListNode{8, &ListNode{4, &ListNode{5, nil}}}
	a := &ListNode{4, &ListNode{1, common}}
	b := &ListNode{5, &ListNode{6, &ListNode{1, common}}}
	fmt.Println(GetIntersectionNode(a, b).Val)
}
```
