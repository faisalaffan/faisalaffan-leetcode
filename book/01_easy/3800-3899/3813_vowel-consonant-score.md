# 3813 — Vowel Consonant Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func VowelConsonantScore(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3813: Vowel-Consonant Score
// https://leetcode.com/problems/vowel-consonant-score/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(VowelConsonantScore("cooear"))
	fmt.Println(VowelConsonantScore("axeyizou"))
	fmt.Println(VowelConsonantScore("au 123"))
}

// Time: O(n)
// Space: O(1)
func VowelConsonantScore(s string) int {
	vowels := 0
	consonants := 0
	for _, ch := range s {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			vowels++
		} else if ch >= 'a' && ch <= 'z' {
			consonants++
		}
	}
	if consonants == 0 {
		return 0
	}
	return vowels / consonants
}
```
