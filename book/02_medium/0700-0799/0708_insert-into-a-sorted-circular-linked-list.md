# 0708 — Insert Into A Sorted Circular Linked List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func insert(head *CNode, insertVal int) *CNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #708: Insert into a Sorted Circular Linked List
// https://leetcode.com/problems/insert-into-a-sorted-circular-linked-list/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	head := &CNode{Val: 3}
	head.Next = &CNode{Val: 4}
	head.Next.Next = &CNode{Val: 1}
	head.Next.Next.Next = head

	result := insert(head, 2)
	fmt.Println(result.Val)
}

type CNode struct {
	Val  int
	Next *CNode
}

func insert(head *CNode, insertVal int) *CNode {
	newNode := &CNode{Val: insertVal}
	if head == nil {
		newNode.Next = newNode
		return newNode
	}

	curr := head
	for curr.Next != head {
		if curr.Val <= insertVal && insertVal <= curr.Next.Val {
			break
		}
		if curr.Val > curr.Next.Val {
			if insertVal >= curr.Val || insertVal <= curr.Next.Val {
				break
			}
		}
		curr = curr.Next
	}

	newNode.Next = curr.Next
	curr.Next = newNode
	return head
}
```
