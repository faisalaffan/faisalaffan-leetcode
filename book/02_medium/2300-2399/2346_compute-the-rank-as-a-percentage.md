# 2346 — Compute The Rank As A Percentage

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func rankAsPercentage(ranks []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2346: Compute the Rank as a Percentage
// https://leetcode.com/problems/compute-the-rank-as-a-percentage/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func rankAsPercentage(ranks []int) []int {
	n := len(ranks)
	// This is a database-style problem. Simplified:
	// For each student, compute (rank-1)*100/(n-1)
	// But since it's "compute the rank as a percentage" database problem:
	// The percentage = (R-1)*100 / (N-1)
  // Alokasi slice integer
	result := make([]int, n)
	for i, r := range ranks {
		result[i] = (r - 1) * 100 / (n - 1)
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(rankAsPercentage([]int{1, 2, 3, 4}))
	// Expected: [0, 33, 66, 100]

	// Test case 2
	fmt.Println(rankAsPercentage([]int{1, 2}))
	// Expected: [0, 100]
}
```
