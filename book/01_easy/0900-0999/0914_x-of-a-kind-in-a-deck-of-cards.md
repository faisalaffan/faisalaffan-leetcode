# 0914 — X Of A Kind In A Deck Of Cards

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func hasGroupsSizeX(deck []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, GCD / Matematika

**Kompleksitas Waktu:** O(n log m) where m is max count. Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #914: X of a Kind in a Deck of Cards
// https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1})) // true
	fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3})) // false
	fmt.Println(hasGroupsSizeX([]int{1}))                        // false
}

// hasGroupsSizeX checks if the deck can be split into groups of equal size with same values.
// Time: O(n log m) where m is max count. Space: O(n).
func hasGroupsSizeX(deck []int) bool {
  // Membuat map (HashMap) — pencarian O(1)
	counts := make(map[int]int)
	for _, v := range deck {
		counts[v]++
	}
	g := -1
	for _, c := range counts {
		if g == -1 {
			g = c
		} else {
			g = gcd(g, c)
		}
	}
	return g >= 2
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```
