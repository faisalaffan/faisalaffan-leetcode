# 1529 — Minimum Suffix Flips

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinFlips(target string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1529: Minimum Suffix Flips
// https://leetcode.com/problems/minimum-suffix-flips/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinFlips("10111"))
	fmt.Println(MinFlips("101"))
	fmt.Println(MinFlips("00000"))
}

func MinFlips(target string) int {
	// Time: O(N), Space: O(1)
	// Count transitions from 0 to 1 or 1 to 0
	flips := 0
	curr := byte('0') // current state of flipped prefix

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(target); i++ {
		if target[i] != curr {
			flips++
			curr = target[i]
		}
	}

	return flips
}
```
