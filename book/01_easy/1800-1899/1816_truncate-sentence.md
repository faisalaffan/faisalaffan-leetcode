# 1816 — Truncate Sentence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TruncateSentence(s string, k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1816: Truncate Sentence
// https://leetcode.com/problems/truncate-sentence/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n), Space: O(n)
func TruncateSentence(s string, k int) string {
	words := strings.Split(s, " ")
	return strings.Join(words[:k], " ")
}

func main() {
	fmt.Println(TruncateSentence("Hello how are you Contestant", 4))
	fmt.Println(TruncateSentence("What is the solution to this problem", 4))
	fmt.Println(TruncateSentence("chopper is not a tanuki", 5))
}
```
