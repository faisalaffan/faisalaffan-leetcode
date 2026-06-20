# 0603 — Consecutive Available Seats

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ConsecutiveAvailableSeats() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


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
