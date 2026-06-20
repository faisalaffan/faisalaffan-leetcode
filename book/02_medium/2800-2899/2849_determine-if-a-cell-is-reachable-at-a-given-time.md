# 2849 — Determine If A Cell Is Reachable At A Given Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DetermineIfACellIsReachableAtAGivenTime(sx int, sy int, fx int, fy int, t int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2849: Determine if a Cell Is Reachable at a Given Time
// https://leetcode.com/problems/determine-if-a-cell-is-reachable-at-a-given-time/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func DetermineIfACellIsReachableAtAGivenTime(sx int, sy int, fx int, fy int, t int) bool {
	dx := sx - fx
	if dx < 0 {
		dx = -dx
	}
	dy := sy - fy
	if dy < 0 {
		dy = -dy
	}
	minDist := dx
	if dy > minDist {
		minDist = dy
	}
	if minDist == 0 {
		return t != 1
	}
	return t >= minDist
}

func main() {
	fmt.Println(DetermineIfACellIsReachableAtAGivenTime(1, 2, 3, 4, 3))
	fmt.Println(DetermineIfACellIsReachableAtAGivenTime(1, 2, 1, 2, 1))
	fmt.Println(DetermineIfACellIsReachableAtAGivenTime(1, 1, 1, 1, 0))
}
```
