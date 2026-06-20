# 1265 — Print Immutable Linked List In Reverse

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func printLinkedListInReverse(head *ImmutableListNode) 
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n) for recursive stack

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1265: Print Immutable Linked List in Reverse
// https://leetcode.com/problems/print-immutable-linked-list-in-reverse/
// Difficulty: Medium [Paid]

// Print linked list in reverse using recursion (or stack).
// Immutable means we can't modify the list.

// Time: O(n)
// Space: O(n) for recursive stack

type ImmutableListNode struct {
	val  int
	next *ImmutableListNode
}

func (node *ImmutableListNode) getValue() int {
	return node.val
}

func (node *ImmutableListNode) getNext() *ImmutableListNode {
	return node.next
}

func printLinkedListInReverse(head *ImmutableListNode) {
	if head == nil {
		return
	}
	printLinkedListInReverse(head.getNext())
	fmt.Printf("%d ", head.getValue())
}

func main() {
	head := &ImmutableListNode{1, &ImmutableListNode{2, &ImmutableListNode{3, nil}}}
	fmt.Printf("Reversed: ")
	printLinkedListInReverse(head)
	fmt.Println()
}
```
