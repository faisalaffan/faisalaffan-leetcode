# 3374 — First Letter Capitalization Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FirstLetterCapitalizationIi(title string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3374: First Letter Capitalization II
// https://leetcode.com/problems/first-letter-capitalization-ii/
// Difficulty: Hard
//
// Capitalize first letter of each word after punctuation separators.

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(FirstLetterCapitalizationIi("hello world"))
	fmt.Println(FirstLetterCapitalizationIi("Leetcode is fun"))
}

func FirstLetterCapitalizationIi(title string) string {
	r := []rune(title)
	n := len(r)
	capitalize := true
	for i := 0; i < n; i++ {
		if unicode.IsLetter(r[i]) {
			if capitalize {
				r[i] = unicode.ToUpper(r[i])
				capitalize = false
			} else {
				r[i] = unicode.ToLower(r[i])
			}
		} else {
			capitalize = true
		}
	}
	return string(r)
}
```
