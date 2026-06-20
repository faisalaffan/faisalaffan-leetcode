# 2833 — Furthest Point From Origin

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FurthestPointFromOrigin(moves string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2833: Furthest Point From Origin
// https://leetcode.com/problems/furthest-point-from-origin/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(FurthestPointFromOrigin("L_RL__R"))
	fmt.Println(FurthestPointFromOrigin("_R__LL_"))
}

func FurthestPointFromOrigin(moves string) int {
	countL, countR, countUnderscore := 0, 0, 0
	for _, c := range moves {
		switch c {
		case 'L':
			countL++
		case 'R':
			countR++
		case '_':
			countUnderscore++
		}
	}
	diff := countL - countR
	if diff < 0 {
		diff = -diff
	}
	return diff + countUnderscore
}
```
