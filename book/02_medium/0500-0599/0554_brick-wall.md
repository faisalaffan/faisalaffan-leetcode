# 0554 — Brick Wall

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LeastBricks(wall [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n * m) where n = rows, m = avg bricks per row  
**Kompleksitas Ruang:** O(n * m)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #554: Brick Wall
// https://leetcode.com/problems/brick-wall/
// Difficulty: Medium
// Time: O(n * m) where n = rows, m = avg bricks per row
// Space: O(n * m)

import "fmt"

func main() {
	wall := [][]int{
		{1, 2, 2, 1},
		{3, 1, 2},
		{1, 3, 2},
		{2, 4},
		{3, 1, 2},
		{1, 3, 1, 1},
	}
	fmt.Println(LeastBricks(wall))
}

func LeastBricks(wall [][]int) int {
  // Membuat map (HashMap) — pencarian O(1)
	gapCount := make(map[int]int)

	for _, row := range wall {
		pos := 0
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(row)-1; i++ {
			pos += row[i]
			gapCount[pos]++
		}
	}

	maxGaps := 0
	for _, count := range gapCount {
		if count > maxGaps {
			maxGaps = count
		}
	}

	return len(wall) - maxGaps
}
```
