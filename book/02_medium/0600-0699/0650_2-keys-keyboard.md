# 0650 — 2 Keys Keyboard

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minSteps(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n sqrt(n))  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #650: 2 Keys Keyboard
// https://leetcode.com/problems/2-keys-keyboard/
// Difficulty: Medium
// Time: O(n sqrt(n))
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(minSteps(3))
	fmt.Println(minSteps(1))
	fmt.Println(minSteps(10))
}

func minSteps(n int) int {
	result := 0
	d := 2

	for n > 1 {
		for n%d == 0 {
			result += d
			n /= d
		}
		d++
	}

	return result
}
```
