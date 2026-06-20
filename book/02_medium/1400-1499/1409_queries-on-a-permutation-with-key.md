# 1409 — Queries On A Permutation With Key

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func processQueries(queries []int, m int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n) where m = len(queries), n = m (since P has m elements)  
**Kompleksitas Ruang:** O(n) for the permutation

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1409: Queries on a Permutation With Key
// https://leetcode.com/problems/queries-on-a-permutation-with-key/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(processQueries([]int{3, 1, 2, 1}, 5)) // [2,1,2,1]

	// Test case 2
	fmt.Println(processQueries([]int{4, 1, 2, 2}, 4)) // [3,1,2,0]

	// Test case 3
	fmt.Println(processQueries([]int{7, 5, 5, 8, 3}, 8)) // [6,5,0,7,5]
}

// Time: O(m*n) where m = len(queries), n = m (since P has m elements)
// Space: O(n) for the permutation
func processQueries(queries []int, m int) []int {
	// Build permutation P = [1, 2, ..., m]
  // Alokasi slice integer
	p := make([]int, m)
	for i := 0; i < m; i++ {
		p[i] = i + 1
	}

  // Alokasi slice integer
	result := make([]int, len(queries))

	for idx, q := range queries {
		// Find position of q in P
		pos := 0
		for p[pos] != q {
			pos++
		}
		result[idx] = pos

		// Move q to front by shifting elements before it
		for i := pos; i > 0; i-- {
			p[i] = p[i-1]
		}
		p[0] = q
	}

	return result
}
```
