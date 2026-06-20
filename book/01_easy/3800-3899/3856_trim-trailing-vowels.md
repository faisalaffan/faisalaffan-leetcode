# 3856 — Trim Trailing Vowels

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TrimTrailingVowels(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3856: Trim Trailing Vowels
// https://leetcode.com/problems/trim-trailing-vowels/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TrimTrailingVowels("idea"))
	fmt.Println(TrimTrailingVowels("day"))
	fmt.Println(TrimTrailingVowels("aeiou"))
}

// Time: O(n)
// Space: O(n)
func TrimTrailingVowels(s string) string {
	i := len(s) - 1
	for i >= 0 && (s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u') {
		i--
	}
	return s[:i+1]
}
```
