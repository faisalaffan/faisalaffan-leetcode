# 2803 — Factorial Generator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FactorialGenerator() func() int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) per call  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2803: Factorial Generator
// https://leetcode.com/problems/factorial-generator/
// Difficulty: Easy [Paid]
// Time: O(1) per call | Space: O(1)
// Note: JS problem, adapted to Go. Generator that yields factorials.

import "fmt"

func main() {
	gen := FactorialGenerator()
	for i := 0; i < 5; i++ {
		fmt.Println(gen())
	}
}

func FactorialGenerator() func() int {
	n := 0
	curr := 1
	return func() int {
  // Edge case: input kosong — langsung return
		if n == 0 {
			n++
			return 1
		}
		curr *= n
		n++
		return curr
	}
}
```
