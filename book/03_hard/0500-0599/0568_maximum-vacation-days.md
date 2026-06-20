# 0568 — Maximum Vacation Days

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxVacationDays(flights [][]int, days [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #568: Maximum Vacation Days
// https://leetcode.com/problems/maximum-vacation-days/
// Difficulty: Hard

import (
	"fmt"
)

func main() {
	flights := [][]int{
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0},
	}
	days := [][]int{
		{1, 3, 1},
		{6, 0, 3},
		{3, 3, 3},
	}
	fmt.Println(maxVacationDays(flights, days)) // Expected: 12
}

func maxVacationDays(flights [][]int, days [][]int) int {
	n := len(flights)   // cities
	k := len(days[0])   // weeks

	// prev[j] = max vacation days ending at city j for current week
  // Alokasi slice integer
	prev := make([]int, n)
	for j := 0; j < n; j++ {
		// Week 0: can we reach city j?
		if j == 0 || flights[0][j] == 1 {
			prev[j] = days[j][0]
		} else {
			prev[j] = -1
		}
	}

	for w := 1; w < k; w++ {
  // Alokasi slice integer
		cur := make([]int, n)
		for j := 0; j < n; j++ {
			cur[j] = -1
		}
		for j := 0; j < n; j++ {
			for i := 0; i < n; i++ {
				if prev[i] >= 0 && (i == j || flights[i][j] == 1) {
					if prev[i]+days[j][w] > cur[j] {
						cur[j] = prev[i] + days[j][w]
					}
				}
			}
		}
		prev = cur
	}

	ans := 0
	for j := 0; j < n; j++ {
		if prev[j] > ans {
			ans = prev[j]
		}
	}
	return ans
}
```
