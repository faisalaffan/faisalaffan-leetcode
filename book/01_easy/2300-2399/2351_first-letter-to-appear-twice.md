# 2351 — First Letter To Appear Twice

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FirstLetterToAppearTwice(s string) byte
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2351: First Letter to Appear Twice
// https://leetcode.com/problems/first-letter-to-appear-twice/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(string(FirstLetterToAppearTwice("abccbaacz"))) // "c"
	fmt.Println(string(FirstLetterToAppearTwice("abcdd")))      // "d"
}

func FirstLetterToAppearTwice(s string) byte {
	seen := [26]bool{}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if seen[idx] {
			return s[i]
		}
		seen[idx] = true
	}
	return 0
}
```
