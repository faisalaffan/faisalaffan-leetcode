# 1078 — Occurrences After Bigram

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findOcurrences(text, first, second string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1078: Occurrences After Bigram
// https://leetcode.com/problems/occurrences-after-bigram/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findOcurrences("alice is a good girl she is a good student", "a", "good"))
	// ["girl","student"]
	fmt.Println(findOcurrences("we will we will rock you", "we", "will"))
	// ["we","rock"]
}

// LeetCode submission: findOcurrences
func findOcurrences(text, first, second string) []string {
	words := strings.Fields(text)
	var ans []string
	for i := 2; i < len(words); i++ {
		if words[i-2] == first && words[i-1] == second {
			ans = append(ans, words[i])
		}
	}
	return ans
}
```
