# 1739 — Building Boxes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumBoxes(n int) int
```

> **💡 Hint:** Greedy accumulation. Build a tetrahedral layer structure,

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1739: Building Boxes
// https://leetcode.com/problems/building-boxes/
// Difficulty: Hard
//
// You have n boxes to place on the floor. The boxes must be placed
// such that:
// - Each box is at the corner of a unit cube grid.
// - Boxes can be stacked, but each box above must be supported.
// Find the minimum number of boxes touching the floor.
//
// Approach: Greedy accumulation. Build a tetrahedral layer structure,
// counting boxes placed on the floor minimally.

import "fmt"

func main() {
	// Example 1
	fmt.Println(minimumBoxes(3))
	// Example 2
	fmt.Println(minimumBoxes(4))
	// Example 3
	fmt.Println(minimumBoxes(10))
	// Edge: n = 1
	fmt.Println(minimumBoxes(1))
}

func minimumBoxes(n int) int {
	s, k := 0, 1
	for s+k*(k+1)/2 <= n {
		s += k * (k + 1) / 2
		k++
	}
	k--
	ans := k * (k + 1) / 2
	k = 1
	for s < n {
		ans++
		s += k
		k++
	}
	return ans
}
```
