# 0830 — Positions Of Large Groups

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func largeGroupPositions(s string) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1) excluding output.  
**Kompleksitas Ruang:** O(1) excluding output.

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #830: Positions of Large Groups
// https://leetcode.com/problems/positions-of-large-groups/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(largeGroupPositions("abbxxxxzzy"))    // [[3,6]]
	fmt.Println(largeGroupPositions("abc"))           // []
	fmt.Println(largeGroupPositions("abcdddeeeeaabbbcd")) // [[3,5],[6,9],[12,14]]
}

// largeGroupPositions finds all large groups (consecutive identical characters of length >= 3).
// Time: O(n). Space: O(1) excluding output.
func largeGroupPositions(s string) [][]int {
  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, 0)
	start := 0
	for i := 1; i <= len(s); i++ {
		if i == len(s) || s[i] != s[start] {
			if i-start >= 3 {
				result = append(result, []int{start, i - 1})
			}
			start = i
		}
	}
	return result
}
```
