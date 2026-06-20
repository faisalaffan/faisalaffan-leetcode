# 2807 — Insert Greatest Common Divisors In Linked List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func InsertGreatestCommonDivisorsInLinkedList(head *ListNode) *ListNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** GCD / Matematika

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **GCD / Matematika** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2807: Insert Greatest Common Divisors in Linked List
// https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func InsertGreatestCommonDivisorsInLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	cur := head
	for cur != nil && cur.Next != nil {
		g := gcd(cur.Val, cur.Next.Val)
		node := &ListNode{Val: g, Next: cur.Next}
		cur.Next = node
		cur = node.Next
	}

	return head
}

func printList(head *ListNode) {
	for head != nil {
		if head.Next != nil {
			fmt.Printf("%d -> ", head.Val)
		} else {
			fmt.Println(head.Val)
		}
		head = head.Next
	}
}

func main() {
	head := &ListNode{18, &ListNode{6, &ListNode{10, &ListNode{3, nil}}}}
	printList(InsertGreatestCommonDivisorsInLinkedList(head))

	head2 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	printList(InsertGreatestCommonDivisorsInLinkedList(head2))
}
```
