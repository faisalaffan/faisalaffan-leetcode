# 0888 — Fair Candy Swap

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func fairCandySwap(aliceSizes []int, bobSizes []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n + m). Space: O(m).  
**Kompleksitas Ruang:** O(m).

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #888: Fair Candy Swap
// https://leetcode.com/problems/fair-candy-swap/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(fairCandySwap([]int{1, 1}, []int{2, 2}))       // [1,2]
	fmt.Println(fairCandySwap([]int{1, 2}, []int{2, 3}))       // [1,2]
	fmt.Println(fairCandySwap([]int{2}, []int{1, 3}))          // [2,3]
}

// fairCandySwap finds a candy swap that makes both Alice and Bob have equal candy.
// Time: O(n + m). Space: O(m).
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sumA, sumB := 0, 0
  // Membuat map (HashMap) — pencarian O(1)
	bSet := make(map[int]bool)
	for _, v := range aliceSizes {
		sumA += v
	}
	for _, v := range bobSizes {
		sumB += v
		bSet[v] = true
	}
	diff := (sumB - sumA) / 2
	for _, a := range aliceSizes {
		if bSet[a+diff] {
			return []int{a, a + diff}
		}
	}
	return nil
}
```
