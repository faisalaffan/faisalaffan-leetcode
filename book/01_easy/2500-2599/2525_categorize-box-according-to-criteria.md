# 2525 — Categorize Box According To Criteria

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CategorizeBoxAccordingToCriteria(length int, width int, height int, mass int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2525: Categorize Box According to Criteria
// https://leetcode.com/problems/categorize-box-according-to-criteria/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CategorizeBoxAccordingToCriteria(1000, 35, 700, 300)) // "Heavy"
	fmt.Println(CategorizeBoxAccordingToCriteria(200, 50, 800, 50))   // "Neither"
}

func CategorizeBoxAccordingToCriteria(length int, width int, height int, mass int) string {
	volume := length * width * height
	isBulky := length >= 10000 || width >= 10000 || height >= 10000 || volume >= 1000000000
	isHeavy := mass >= 100

	if isBulky && isHeavy {
		return "Both"
	}
	if isBulky {
		return "Bulky"
	}
	if isHeavy {
		return "Heavy"
	}
	return "Neither"
}
```
