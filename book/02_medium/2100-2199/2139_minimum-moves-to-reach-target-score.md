# 2139 — Minimum Moves To Reach Target Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minMoves(target int, maxDoubles int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log target)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2139: Minimum Moves to Reach Target Score
// https://leetcode.com/problems/minimum-moves-to-reach-target-score/
// Difficulty: Medium
// Time: O(log target) | Space: O(1)

import "fmt"

func minMoves(target int, maxDoubles int) int {
	moves := 0
	for target > 1 {
		if maxDoubles == 0 {
			moves += target - 1
			break
		}
		if target%2 == 1 {
			target--
			moves++
		} else {
			target /= 2
			maxDoubles--
			moves++
		}
	}
	return moves
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minMoves(5, 0))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", minMoves(19, 2))
	// Expected: 7

	// Test case 3
	fmt.Println("Test 3:", minMoves(10, 4))
	// Expected: 4
}
```
