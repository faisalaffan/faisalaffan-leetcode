# 1854 — Maximum Population Year

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumPopulation(logs [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + range), Space: O(range)  
**Kompleksitas Ruang:** O(range)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1854: Maximum Population Year
// https://leetcode.com/problems/maximum-population-year/
// Difficulty: Easy

import "fmt"

// Time: O(n + range), Space: O(range)
func MaximumPopulation(logs [][]int) int {
  // Alokasi slice integer
	delta := make([]int, 101) // 1950 to 2050
	for _, log := range logs {
		delta[log[0]-1950]++
		delta[log[1]-1950]--
	}
	maxPop := 0
	currentPop := 0
	bestYear := 1950
	for i := 0; i < 101; i++ {
		currentPop += delta[i]
		if currentPop > maxPop {
			maxPop = currentPop
			bestYear = 1950 + i
		}
	}
	return bestYear
}

func main() {
	fmt.Println(MaximumPopulation([][]int{{1993, 1999}, {2000, 2010}}))
	fmt.Println(MaximumPopulation([][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}}))
}
```
