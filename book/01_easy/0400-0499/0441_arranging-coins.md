# 0441 — Arranging Coins

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ArrangingCoins(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #441: Arranging Coins
// https://leetcode.com/problems/arranging-coins/
// Difficulty: Easy

import "fmt"

// Time: O(log n), Space: O(1)
func ArrangingCoins(n int) int {
	lo, hi := 1, n
	for lo <= hi {
		mid := lo + (hi-lo)/2
		sum := mid * (mid + 1) / 2
		if sum == n {
			return mid
		} else if sum < n {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi
}

func main() {
	fmt.Println(ArrangingCoins(5))
	fmt.Println(ArrangingCoins(8))
	fmt.Println(ArrangingCoins(1))
}
```
