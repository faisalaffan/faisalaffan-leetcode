# 0455 — Assign Cookies

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func AssignCookies(g, s []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n + m log m), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #455: Assign Cookies
// https://leetcode.com/problems/assign-cookies/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n + m log m), Space: O(1)
func AssignCookies(g, s []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(g)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
		}
		j++
	}
	return i
}

func main() {
	fmt.Println(AssignCookies([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(AssignCookies([]int{1, 2}, []int{1, 2, 3}))
}
```
