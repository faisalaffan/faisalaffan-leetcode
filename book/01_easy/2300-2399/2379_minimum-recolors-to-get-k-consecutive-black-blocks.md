# 2379 — Minimum Recolors To Get K Consecutive Black Blocks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumRecolorsToGetKConsecutiveBlackBlocks(blocks string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2379: Minimum Recolors to Get K Consecutive Black Blocks
// https://leetcode.com/problems/minimum-recolors-to-get-k-consecutive-black-blocks/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(MinimumRecolorsToGetKConsecutiveBlackBlocks("WBBWWBBWBW", 7)) // 3
	fmt.Println(MinimumRecolorsToGetKConsecutiveBlackBlocks("WBWBBBW", 2))    // 0
}

func MinimumRecolorsToGetKConsecutiveBlackBlocks(blocks string, k int) int {
	wCount := 0
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			wCount++
		}
	}
	minOps := wCount
	for i := k; i < len(blocks); i++ {
		if blocks[i-k] == 'W' {
			wCount--
		}
		if blocks[i] == 'W' {
			wCount++
		}
		if wCount < minOps {
			minOps = wCount
		}
	}
	return minOps
}
```
