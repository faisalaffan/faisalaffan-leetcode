# 2129 — Capitalize The Title

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CapitalizeTheTitle(title string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2129: Capitalize the Title
// https://leetcode.com/problems/capitalize-the-title/
// Difficulty: Easy

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(CapitalizeTheTitle("capiTalIze tHe titLe")) // "Capitalize The Title"
	fmt.Println(CapitalizeTheTitle("First leTTER of EACH Word")) // "First Letter of Each Word"
	fmt.Println(CapitalizeTheTitle("i lOve leetcode"))           // "i Love Leetcode"
}

// Time: O(n), Space: O(n)
func CapitalizeTheTitle(title string) string {
	words := strings.Fields(title)
	for i, w := range words {
		lower := strings.ToLower(w)
		if len(lower) > 2 {
			runes := []rune(lower)
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		} else {
			words[i] = lower
		}
	}
	return strings.Join(words, " ")
}
```
