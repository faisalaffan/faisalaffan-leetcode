# 3294 — Convert Doubly Linked List To Array Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func toArray(node *DListNode) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) Space: O(n) for output  
**Kompleksitas Ruang:** O(n) for output

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3294: Convert Doubly Linked List to Array II
// https://leetcode.com/problems/convert-doubly-linked-list-to-array-ii/
// Difficulty: Medium
// Time: O(n) Space: O(n) for output

import "fmt"

func main() {
	// Build list: 1 <-> 2 <-> 3 <-> 4 <-> 5
	nodes := make([]*DListNode, 5)
	for i := 0; i < 5; i++ {
		nodes[i] = &DListNode{Val: i + 1}
	}
	for i := 0; i < 5; i++ {
		if i > 0 {
			nodes[i].Prev = nodes[i-1]
		}
		if i < 4 {
			nodes[i].Next = nodes[i+1]
		}
	}
	fmt.Println(toArray(nodes[2])) // [1 2 3 4 5]

	// Single node
	single := &DListNode{Val: 42}
	fmt.Println(toArray(single)) // [42]

	// Two nodes
	twoA := &DListNode{Val: 10}
	twoB := &DListNode{Val: 20}
	twoA.Next = twoB
	twoB.Prev = twoA
	fmt.Println(toArray(twoB)) // [10 20]
}

type DListNode struct {
	Val  int
	Next *DListNode
	Prev *DListNode
}

func toArray(node *DListNode) []int {
	if node == nil {
		return []int{}
	}

	// Find head
	head := node
	for head.Prev != nil {
		head = head.Prev
	}

	// Collect values
	var res []int
	for cur := head; cur != nil; cur = cur.Next {
		res = append(res, cur.Val)
	}
	return res
}
```
