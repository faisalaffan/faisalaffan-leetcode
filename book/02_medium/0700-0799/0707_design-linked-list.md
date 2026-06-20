# 0707 — Design Linked List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() MyLinkedList
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) for get/addAtIndex/deleteAtIndex, O(1) for addAtHead/addAtTail  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #707: Design Linked List
// https://leetcode.com/problems/design-linked-list/
// Difficulty: Medium
// Time: O(n) for get/addAtIndex/deleteAtIndex, O(1) for addAtHead/addAtTail
// Space: O(n)

import "fmt"

func main() {
	l := Constructor()
	l.AddAtHead(1)
	l.AddAtTail(3)
	l.AddAtIndex(1, 2)
	fmt.Println(l.Get(1))
	l.DeleteAtIndex(1)
	fmt.Println(l.Get(1))
}

type MyLinkedList struct {
	head *LinkNode
	size int
}

type LinkNode struct {
	val  int
	next *LinkNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	return curr.val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.head = &LinkNode{val, l.head}
	l.size++
}

func (l *MyLinkedList) AddAtTail(val int) {
	if l.head == nil {
		l.AddAtHead(val)
		return
	}
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &LinkNode{val: val}
	l.size++
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.size {
		return
	}
	if index == 0 {
		l.AddAtHead(val)
		return
	}
	curr := l.head
	for i := 0; i < index-1; i++ {
		curr = curr.next
	}
	curr.next = &LinkNode{val, curr.next}
	l.size++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size {
		return
	}
	if index == 0 {
		l.head = l.head.next
		l.size--
		return
	}
	curr := l.head
	for i := 0; i < index-1; i++ {
		curr = curr.next
	}
	curr.next = curr.next.next
	l.size--
}
```
