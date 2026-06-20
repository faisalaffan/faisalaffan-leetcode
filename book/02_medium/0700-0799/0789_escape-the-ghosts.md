# 0789 — Escape The Ghosts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func escapeGhosts(ghosts [][]int, target []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #789: Escape The Ghosts
// https://leetcode.com/problems/escape-the-ghosts/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(escapeGhosts([][]int{{1, 0}, {0, 3}}, []int{0, 1}))
	fmt.Println(escapeGhosts([][]int{{1, 0}}, []int{2, 0}))
}

func escapeGhosts(ghosts [][]int, target []int) bool {
	myDist := abs(target[0]) + abs(target[1])

	for _, g := range ghosts {
		ghostDist := abs(g[0]-target[0]) + abs(g[1]-target[1])
		if ghostDist <= myDist {
			return false
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
