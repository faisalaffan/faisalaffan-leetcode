# 3063 — Linked List Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LinkedListFrequency(head *ListNode) map[int]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3063: Linked List Frequency
// https://leetcode.com/problems/linked-list-frequency/
// Difficulty: Easy [Paid]

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// LeetCode name: frequencies
	head := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{3, nil}}}}}}
	fmt.Println(LinkedListFrequency(head)) // map[1:1 2:2 3:3]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: frequencies
func LinkedListFrequency(head *ListNode) map[int]int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	curr := head
	for curr != nil {
		freq[curr.Val]++
		curr = curr.Next
	}
	return freq
}
```
