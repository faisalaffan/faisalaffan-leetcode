# 0765 — Couples Holding Hands

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minSwapsCouples(row []int) int
```

> **💡 Hint:** Greedy

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #765: Couples Holding Hands
// https://leetcode.com/problems/couples-holding-hands/
// Difficulty: Hard
//
// N couples (2N people) sit in 2N seats. Each couple is numbered
// (0,1), (2,3), (4,5), ... Minimum swaps to make every couple sit
// together (seats 2i and 2i+1 are adjacent).
//
// Approach: Greedy
// For each even position i, check if row[i] and row[i+1] are a couple.
// If not, find the partner of row[i] and swap it with row[i+1].

import "fmt"

func main() {
	fmt.Println(minSwapsCouples([]int{0, 2, 1, 3}))    // 1
	fmt.Println(minSwapsCouples([]int{3, 2, 0, 1}))    // 0
	fmt.Println(minSwapsCouples([]int{0, 1, 2, 3}))    // 0
	fmt.Println(minSwapsCouples([]int{3, 0, 1, 2}))    // 1
	fmt.Println(minSwapsCouples([]int{5, 4, 2, 6, 3, 1, 0, 7})) // 2
}

func minSwapsCouples(row []int) int {
	n := len(row)
	// pos[person] = index in row
  // Alokasi slice integer
	pos := make([]int, n)
	for i, p := range row {
		pos[p] = i
	}

	swaps := 0
	for i := 0; i < n; i += 2 {
		p1 := row[i]
		// partner of p1: if p1 is even, partner is p1+1; if odd, partner is p1-1
		partner := p1 ^ 1 // XOR trick: (even -> +1, odd -> -1)
		if row[i+1] == partner {
			continue // already sitting together
		}
		// Find where partner is, swap it with row[i+1]
		j := pos[partner]
		// swap row[i+1] with row[j]
		row[i+1], row[j] = row[j], row[i+1]
		// update positions
		pos[row[j]] = j
		pos[row[i+1]] = i + 1
		swaps++
	}
	return swaps
}
```
