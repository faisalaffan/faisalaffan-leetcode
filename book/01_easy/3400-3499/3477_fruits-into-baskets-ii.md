# 3477 — Fruits Into Baskets Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FruitsIntoBasketsIi(fruits []int, baskets []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3477: Fruits Into Baskets II
// https://leetcode.com/problems/fruits-into-baskets-ii/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FruitsIntoBasketsIi([]int{4, 2, 5}, []int{3, 5, 4}))
	fmt.Println(FruitsIntoBasketsIi([]int{3, 6, 1}, []int{6, 4, 7}))
}

// FruitsIntoBasketsIi counts fruits that cannot be placed into baskets.
// Each fruit i can go into basket j if fruits[i] <= baskets[j].
// Time: O(n * m). Space: O(1).
func FruitsIntoBasketsIi(fruits []int, baskets []int) int {
	used := make([]bool, len(baskets))
	unplaced := 0
	for _, f := range fruits {
		placed := false
		for j, b := range baskets {
			if !used[j] && f <= b {
				used[j] = true
				placed = true
				break
			}
		}
		if !placed {
			unplaced++
		}
	}
	return unplaced
}
```
