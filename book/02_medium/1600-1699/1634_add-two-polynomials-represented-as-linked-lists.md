# 1634 — Add Two Polynomials Represented As Linked Lists

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func printPoly(p *PolyNode) 
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N+M), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1634: Add Two Polynomials Represented as Linked Lists
// https://leetcode.com/problems/add-two-polynomials-represented-as-linked-lists/
// Difficulty: Medium [Paid]

import "fmt"

// PolyNode represents a term in a polynomial.
type PolyNode struct {
	Coefficient int
	Power       int
	Next        *PolyNode
}

func main() {
	// Polynomial 1: 5x^3 + 4x^2 + (-1)x^0
	p1 := &PolyNode{Coefficient: 5, Power: 3}
	p1.Next = &PolyNode{Coefficient: 4, Power: 2}
	p1.Next.Next = &PolyNode{Coefficient: -1, Power: 0}

	// Polynomial 2: 3x^2 + 2x^1 + 1x^0
	p2 := &PolyNode{Coefficient: 3, Power: 2}
	p2.Next = &PolyNode{Coefficient: 2, Power: 1}
	p2.Next.Next = &PolyNode{Coefficient: 1, Power: 0}

	result := AddPoly(p1, p2)
	printPoly(result)

	// Polynomial: x^3 + (-1)x^0
	p3 := &PolyNode{Coefficient: 1, Power: 3, Next: &PolyNode{Coefficient: -1, Power: 0}}
	// Polynomial: x^2 + x^1
	p4 := &PolyNode{Coefficient: 1, Power: 2, Next: &PolyNode{Coefficient: 1, Power: 1}}
	result2 := AddPoly(p3, p4)
	printPoly(result2)
}

func printPoly(p *PolyNode) {
	first := true
	for p != nil {
		if p.Coefficient != 0 {
			if !first {
				fmt.Print(" + ")
			}
			fmt.Printf("%dx^%d", p.Coefficient, p.Power)
			first = false
		}
		p = p.Next
	}
	fmt.Println()
}

func AddPoly(poly1 *PolyNode, poly2 *PolyNode) *PolyNode {
	// Time: O(N+M), Space: O(1)
	dummy := &PolyNode{}
	curr := dummy

	for poly1 != nil && poly2 != nil {
		if poly1.Power > poly2.Power {
			if poly1.Coefficient != 0 {
				curr.Next = &PolyNode{Coefficient: poly1.Coefficient, Power: poly1.Power}
				curr = curr.Next
			}
			poly1 = poly1.Next
		} else if poly2.Power > poly1.Power {
			if poly2.Coefficient != 0 {
				curr.Next = &PolyNode{Coefficient: poly2.Coefficient, Power: poly2.Power}
				curr = curr.Next
			}
			poly2 = poly2.Next
		} else {
			sum := poly1.Coefficient + poly2.Coefficient
			if sum != 0 {
				curr.Next = &PolyNode{Coefficient: sum, Power: poly1.Power}
				curr = curr.Next
			}
			poly1 = poly1.Next
			poly2 = poly2.Next
		}
	}

	for poly1 != nil {
		if poly1.Coefficient != 0 {
			curr.Next = &PolyNode{Coefficient: poly1.Coefficient, Power: poly1.Power}
			curr = curr.Next
		}
		poly1 = poly1.Next
	}

	for poly2 != nil {
		if poly2.Coefficient != 0 {
			curr.Next = &PolyNode{Coefficient: poly2.Coefficient, Power: poly2.Power}
			curr = curr.Next
		}
		poly2 = poly2.Next
	}

	return dummy.Next
}
```
