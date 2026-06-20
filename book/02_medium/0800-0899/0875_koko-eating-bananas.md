# 0875 — Koko Eating Bananas

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func KokoEatingBananas(piles []int, h int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log m) where m = max pile  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #875: Koko Eating Bananas
// https://leetcode.com/problems/koko-eating-bananas/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(KokoEatingBananas([]int{3, 6, 7, 11}, 8))
	fmt.Println(KokoEatingBananas([]int{30, 11, 23, 4, 20}, 5))
	fmt.Println(KokoEatingBananas([]int{30, 11, 23, 4, 20}, 6))
}

// Time: O(n log m) where m = max pile | Space: O(1)
func KokoEatingBananas(piles []int, h int) int {
	maxPile := 0
	for _, p := range piles {
		if p > maxPile {
			maxPile = p
		}
	}

	left, right := 1, maxPile
  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := left + (right-left)/2
		hours := 0
		for _, p := range piles {
			hours += (p + mid - 1) / mid
		}
		if hours <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
```
