# 1154 — Day Of The Year

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func dayOfYear(date string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1154: Day of the Year
// https://leetcode.com/problems/day-of-the-year/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(dayOfYear("2019-01-09")) // 9
	fmt.Println(dayOfYear("2019-02-10")) // 41
	fmt.Println(dayOfYear("2000-03-01")) // 61
}

// LeetCode submission: dayOfYear
func dayOfYear(date string) int {
	parts := strings.Split(date, "-")
	year, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	day, _ := strconv.Atoi(parts[2])

	leap := (year%4 == 0 && year%100 != 0) || (year%400 == 0)
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if leap {
		days[1] = 29
	}

	ans := 0
	for i := 0; i < month-1; i++ {
		ans += days[i]
	}
	return ans + day
}
```
