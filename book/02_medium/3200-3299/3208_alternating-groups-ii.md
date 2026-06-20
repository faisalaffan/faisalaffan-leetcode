# 3208 — Alternating Groups Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfAlternatingGroups(colors []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3208: Alternating Groups II
// https://leetcode.com/problems/alternating-groups-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func numberOfAlternatingGroups(colors []int, k int) int {
	n := len(colors)
	ans := 0
	len := 1

	for i := 1; i < n+k-1; i++ {
		if colors[i%n] != colors[(i-1)%n] {
			len++
		} else {
			len = 1
		}
		if len >= k {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfAlternatingGroups([]int{0, 1, 0, 1, 0}, 3)) // Expected: 3
	fmt.Println(numberOfAlternatingGroups([]int{0, 1, 0, 0, 1}, 3)) // Expected: 2
}
```
