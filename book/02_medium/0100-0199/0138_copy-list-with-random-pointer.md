# 0138 — Copy List With Random Pointer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func copyRandomList(head *Node) *Node
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #138: Copy List with Random Pointer
// https://leetcode.com/problems/copy-list-with-random-pointer/
// Difficulty: Medium

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Interleave original nodes with copies
	curr := head
	for curr != nil {
		copy := &Node{Val: curr.Val}
		copy.Next = curr.Next
		curr.Next = copy
		curr = copy.Next
	}

	// Set random pointers for copies
	curr = head
	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
		curr = curr.Next.Next
	}

	// Separate original and copy lists
	dummy := &Node{}
	copyCurr := dummy
	curr = head
	for curr != nil {
		copyCurr.Next = curr.Next
		copyCurr = copyCurr.Next
		curr.Next = curr.Next.Next
		curr = curr.Next
	}

	return dummy.Next
}

func printList(head *Node) {
	for head != nil {
		randomStr := "nil"
		if head.Random != nil {
			randomStr = fmt.Sprintf("%d", head.Random.Val)
		}
		fmt.Printf("[%d, %s] ", head.Val, randomStr)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1: [[7,null],[13,0],[11,4],[10,2],[1,0]]
	n0 := &Node{Val: 7}
	n1 := &Node{Val: 13}
	n2 := &Node{Val: 11}
	n3 := &Node{Val: 10}
	n4 := &Node{Val: 1}
	n0.Next, n0.Random = n1, nil
	n1.Next, n1.Random = n2, n0
	n2.Next, n2.Random = n3, n4
	n3.Next, n3.Random = n4, n2
	n4.Next, n4.Random = nil, n0

	copy := copyRandomList(n0)
	printList(copy)

	// Test case 2
	fmt.Println(copyRandomList(nil)) // nil
}

// Time: O(n) | Space: O(1)
```
