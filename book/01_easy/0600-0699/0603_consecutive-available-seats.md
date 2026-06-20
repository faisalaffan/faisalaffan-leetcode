# 0603 — Consecutive Available Seats

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ConsecutiveAvailableSeats() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #603: Consecutive Available Seats
// https://leetcode.com/problems/consecutive-available-seats/
// Difficulty: Easy [Paid]

import "fmt"

func ConsecutiveAvailableSeats() string {
	return "SELECT DISTINCT c1.seat_id FROM Cinema c1 JOIN Cinema c2 ON ABS(c1.seat_id - c2.seat_id) = 1 AND c1.free = 1 AND c2.free = 1 ORDER BY c1.seat_id"
}

func main() {
	fmt.Println(ConsecutiveAvailableSeats())
}
```
