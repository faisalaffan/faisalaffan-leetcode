# 3206 — Alternating Groups I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func AlternatingGroupsI(colors []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3206: Alternating Groups I
// https://leetcode.com/problems/alternating-groups-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(AlternatingGroupsI([]int{1, 1, 1}))
	fmt.Println(AlternatingGroupsI([]int{0, 1, 0, 0, 1}))
}

// AlternatingGroupsI counts the number of groups of 3 adjacent elements where all three are alternating.
// Time: O(n). Space: O(1).
func AlternatingGroupsI(colors []int) int {
	n := len(colors)
	count := 0
	for i := 0; i < n; i++ {
		if colors[i] == colors[(i+2)%n] && colors[i] != colors[(i+1)%n] {
			count++
		}
	}
	return count
}
```
