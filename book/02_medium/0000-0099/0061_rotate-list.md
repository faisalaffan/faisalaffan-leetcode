# 0061 — Rotate List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func rotateRight(head *ListNode, k int) *ListNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #61: Rotate List
// https://leetcode.com/problems/rotate-list/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// Find length and tail
	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	k = k % length
	if k == 0 {
		return head
	}

	// Find new head (length - k)th node
	curr := head
	for i := 0; i < length-k-1; i++ {
		curr = curr.Next
	}

	newHead := curr.Next
	curr.Next = nil
	tail.Next = head

	return newHead
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
	// Test case 1: [1,2,3,4,5], k=2 -> [4,5,1,2,3]
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result := rotateRight(head, 2)
	printList(result)

	// Test case 2: [0,1,2], k=4 -> [2,0,1]
	head = &ListNode{0, &ListNode{1, &ListNode{2, nil}}}
	result = rotateRight(head, 4)
	printList(result)

	// Test case 3: [], k=0 -> []
	result = rotateRight(nil, 0)
	printList(result)
}

// Time: O(n) | Space: O(1)
```
