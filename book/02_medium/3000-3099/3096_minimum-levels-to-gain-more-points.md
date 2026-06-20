# 3096 — Minimum Levels To Gain More Points

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumLevels(possible []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3096: Minimum Levels to Gain More Points
// https://leetcode.com/problems/minimum-levels-to-gain-more-points/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumLevels(possible []int) int {
	n := len(possible)
  // Alokasi slice integer
	score := make([]int, n)
	for i, v := range possible {
		if v == 0 {
			score[i] = -1
		} else {
			score[i] = 1
		}
	}

	total := 0
	for _, v := range score {
		total += v
	}

	prefix := 0
	for i := 0; i < n-1; i++ {
		prefix += score[i]
		if prefix > total-prefix {
			return i + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(minimumLevels([]int{1, 0, 1, 0}))       // Expected: 1
	fmt.Println(minimumLevels([]int{1, 1, 1, 1, 1}))    // Expected: 3
	fmt.Println(minimumLevels([]int{0, 0}))              // Expected: -1
}
```
