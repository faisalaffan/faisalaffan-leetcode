# 2777 — Date Range Generator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DateRangeGenerator(start, end time.Time, step time.Duration) []time.Time
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2777: Date Range Generator
// https://leetcode.com/problems/date-range-generator/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"time"
)

func DateRangeGenerator(start, end time.Time, step time.Duration) []time.Time {
	dates := make([]time.Time, 0)
	for cur := start; !cur.After(end); cur = cur.Add(step) {
		dates = append(dates, cur)
	}
	return dates
}

func main() {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC)
	dates := DateRangeGenerator(start, end, 24*time.Hour)
	for _, d := range dates {
		fmt.Println(d.Format("2006-01-02"))
	}

	fmt.Println("---")

	start2 := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	end2 := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	dates2 := DateRangeGenerator(start2, end2, 24*time.Hour)
	for _, d := range dates2 {
		fmt.Println(d.Format("2006-01-02"))
	}
}
```
