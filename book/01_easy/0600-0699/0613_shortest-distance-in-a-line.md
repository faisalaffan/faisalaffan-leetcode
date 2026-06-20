# 0613 — Shortest Distance In A Line

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ShortestDistanceInALine() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #613: Shortest Distance in a Line
// https://leetcode.com/problems/shortest-distance-in-a-line/
// Difficulty: Easy [Paid]

import "fmt"

func ShortestDistanceInALine() string {
	return "SELECT MIN(ABS(p1.x - p2.x)) AS shortest FROM Point p1 JOIN Point p2 ON p1.x != p2.x"
}

func main() {
	fmt.Println(ShortestDistanceInALine())
}
```
