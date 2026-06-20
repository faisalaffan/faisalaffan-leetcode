# 1290 — Convert Binary Number In A Linked List To Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func getDecimalValue(head *ListNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1290: Convert Binary Number in a Linked List to Integer
// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 1 -> 0 -> 1 = 5
	head := &ListNode{1, &ListNode{0, &ListNode{1, nil}}}
	fmt.Println(getDecimalValue(head)) // 5

	// 0 -> 0 = 0
	fmt.Println(getDecimalValue(&ListNode{0, &ListNode{0, nil}})) // 0
}

// LeetCode submission: getDecimalValue
func getDecimalValue(head *ListNode) int {
	ans := 0
	for head != nil {
		ans = ans*2 + head.Val
		head = head.Next
	}
	return ans
}
```
