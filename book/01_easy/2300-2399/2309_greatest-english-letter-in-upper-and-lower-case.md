# 2309 — Greatest English Letter In Upper And Lower Case

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func GreatestEnglishLetterInUpperAndLowerCase(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2309: Greatest English Letter in Upper and Lower Case
// https://leetcode.com/problems/greatest-english-letter-in-upper-and-lower-case/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(GreatestEnglishLetterInUpperAndLowerCase("lEeTcOdE")) // "E"
	fmt.Println(GreatestEnglishLetterInUpperAndLowerCase("arRAzFif")) // "R"
	fmt.Println(GreatestEnglishLetterInUpperAndLowerCase("AbCdEfGhIjK")) // ""
}

func GreatestEnglishLetterInUpperAndLowerCase(s string) string {
	seen := [26]bool{}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			seen[s[i]-'a'] = true
		}
	}
	best := byte(0)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			idx := s[i] - 'A'
			if seen[idx] && s[i] > best {
				best = s[i]
			}
		}
	}
	if best == 0 {
		return ""
	}
	return string(best)
}
```
