# 0957 — Prison Cells After N Days

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func prisonAfterNDays(cells []int, n int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #957: Prison Cells After N Days
// https://leetcode.com/problems/prison-cells-after-n-days/
// Difficulty: Medium

import "fmt"

// Time: O(1) | Space: O(1)
func prisonAfterNDays(cells []int, n int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[[8]int]int)
	cycle := false

	for n > 0 {
		key := toArray(cells)
		if day, ok := seen[key]; ok && !cycle {
			n %= day - n
			cycle = true
		}
		seen[key] = n

		if n > 0 {
			n--
			cells = nextDay(cells)
		}
	}

	return cells
}

func toArray(cells []int) [8]int {
	return [8]int{cells[0], cells[1], cells[2], cells[3], cells[4], cells[5], cells[6], cells[7]}
}

func nextDay(cells []int) []int {
  // Alokasi slice integer
	next := make([]int, 8)
	for i := 1; i < 7; i++ {
		if cells[i-1] == cells[i+1] {
			next[i] = 1
		}
	}
	return next
}

func main() {
	fmt.Println(prisonAfterNDays([]int{0, 1, 0, 1, 1, 0, 0, 1}, 7))
	fmt.Println(prisonAfterNDays([]int{1, 0, 0, 1, 0, 0, 1, 0}, 1000000000))
}
```
