# 1884 — Egg Drop With 2 Eggs And N Floors

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TwoEggDrop(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1884: Egg Drop With 2 Eggs and N Floors
// https://leetcode.com/problems/egg-drop-with-2-eggs-and-n-floors/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(TwoEggDrop(2))
	fmt.Println(TwoEggDrop(100))
	fmt.Println(TwoEggDrop(10))
}

// Time: O(1), Space: O(1)
func TwoEggDrop(n int) int {
	// Solve x(x+1)/2 >= n
	// x = ceil((-1 + sqrt(1 + 8n)) / 2)
	return int(math.Ceil((-1 + math.Sqrt(1+8*float64(n))) / 2))
}
```
