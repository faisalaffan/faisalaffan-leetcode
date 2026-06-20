# 3263 — Convert Doubly Linked List To Array I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ConvertDoublyLinkedListToArrayI(head *Node) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3263: Convert Doubly Linked List to Array I
// https://leetcode.com/problems/convert-doubly-linked-list-to-array-i/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	// 1 <-> 2 <-> 3
	head := &Node{Val: 1}
	head.Next = &Node{Val: 2, Prev: head}
	head.Next.Next = &Node{Val: 3, Prev: head.Next}
	fmt.Println(ConvertDoublyLinkedListToArrayI(head))
}

// Node represents a doubly-linked list node.
type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

// ConvertDoublyLinkedListToArrayI converts a doubly linked list to an integer array.
// Time: O(n). Space: O(n).
func ConvertDoublyLinkedListToArrayI(head *Node) []int {
	result := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		result = append(result, cur.Val)
	}
	return result
}
```
