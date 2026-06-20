# 0531 — Lonely Pixel I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindLonelyPixel(picture [][]byte) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m + n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #531: Lonely Pixel I
// https://leetcode.com/problems/lonely-pixel-i/
// Difficulty: Medium [Paid]
// Time: O(m * n)
// Space: O(m + n)

import "fmt"

func main() {
	picture := [][]byte{
		{'W', 'W', 'B'},
		{'W', 'B', 'W'},
		{'B', 'W', 'W'},
	}
	fmt.Println(FindLonelyPixel(picture))
}

func FindLonelyPixel(picture [][]byte) int {
	m, n := len(picture), len(picture[0])
  // Alokasi slice integer
	rows := make([]int, m)
  // Alokasi slice integer
	cols := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' {
				rows[i]++
				cols[j]++
			}
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' && rows[i] == 1 && cols[j] == 1 {
				count++
			}
		}
	}

	return count
}
```
