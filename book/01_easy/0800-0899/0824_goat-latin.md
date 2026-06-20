# 0824 — Goat Latin

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func toGoatLatin(sentence string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #824: Goat Latin
// https://leetcode.com/problems/goat-latin/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))  // "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
	fmt.Println(toGoatLatin("The quick brown fox")) // "heTmaa uickqmaaa rainbowmaaaa oxfmaaaaa"
}

// toGoatLatin converts a sentence to Goat Latin.
// Time: O(n). Space: O(n).
func toGoatLatin(sentence string) string {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	words := strings.Fields(sentence)
	var result []string
	for i, word := range words {
		var transformed string
		if vowels[word[0]] {
			transformed = word + "ma"
		} else {
			transformed = word[1:] + string(word[0]) + "ma"
		}
		transformed += strings.Repeat("a", i+1)
		result = append(result, transformed)
	}
	return strings.Join(result, " ")
}
```
