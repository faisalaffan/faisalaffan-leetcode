# 2381 — Shifting Letters Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func shiftingLetters(s string, shifts [][]int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2381: Shifting Letters II
// https://leetcode.com/problems/shifting-letters-ii/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n)
// Difference array to apply range shifts efficiently.

import "fmt"

func main() {
	fmt.Println(shiftingLetters("abc", [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}})) // "ace"
	fmt.Println(shiftingLetters("dztz", [][]int{{0, 0, 0}, {1, 1, 1}}))       // "catz"
}

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
  // Alokasi slice integer
	diff := make([]int, n+1)
	for _, sh := range shifts {
		start, end, dir := sh[0], sh[1], sh[2]
		if dir == 1 {
			diff[start]++
			diff[end+1]--
		} else {
			diff[start]--
			diff[end+1]++
		}
	}

	cur := 0
	res := make([]byte, n)
	for i, ch := range s {
		cur += diff[i]
		shift := ((int(ch-'a')+cur)%26 + 26) % 26
		res[i] = byte('a' + shift)
	}
	return string(res)
}
```
