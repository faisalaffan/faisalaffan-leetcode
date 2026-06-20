# 3368 — First Letter Capitalization

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FirstLetterCapitalization(title string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3368: First Letter Capitalization
// https://leetcode.com/problems/first-letter-capitalization/
// Difficulty: Hard [Paid]
//
// Capitalize first letter of each word, lowercase rest.

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(FirstLetterCapitalization("hello world"))
	fmt.Println(FirstLetterCapitalization("Leetcode is fun"))
}

func FirstLetterCapitalization(title string) string {
	words := strings.Fields(title)
	for i, w := range words {
		r := []rune(w)
		for j := range r {
			if j == 0 {
				r[j] = unicode.ToUpper(r[j])
			} else {
				r[j] = unicode.ToLower(r[j])
			}
		}
		words[i] = string(r)
	}
	return strings.Join(words, " ")
}
```
