# 0575 — Distribute Candies

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DistributeCandies(candyType []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #575: Distribute Candies
// https://leetcode.com/problems/distribute-candies/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func DistributeCandies(candyType []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	types := make(map[int]bool)
	for _, c := range candyType {
		types[c] = true
	}
	maxAllowed := len(candyType) / 2
	if len(types) < maxAllowed {
		return len(types)
	}
	return maxAllowed
}

func main() {
	fmt.Println(DistributeCandies([]int{1, 1, 2, 2, 3, 3}))
	fmt.Println(DistributeCandies([]int{1, 1, 2, 3}))
	fmt.Println(DistributeCandies([]int{6, 6, 6, 6}))
}
```
