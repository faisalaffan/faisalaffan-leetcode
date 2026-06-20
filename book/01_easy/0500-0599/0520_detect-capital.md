# 0520 — Detect Capital

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DetectCapital(word string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #520: Detect Capital
// https://leetcode.com/problems/detect-capital/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func DetectCapital(word string) bool {
	upperCount := 0
	for _, c := range word {
		if c >= 'A' && c <= 'Z' {
			upperCount++
		}
	}
	return upperCount == len(word) || upperCount == 0 || (upperCount == 1 && word[0] >= 'A' && word[0] <= 'Z')
}

func main() {
	fmt.Println(DetectCapital("USA"))
	fmt.Println(DetectCapital("FlaG"))
	fmt.Println(DetectCapital("Google"))
}
```
