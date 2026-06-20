# 0869 — Reordered Power Of 2

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ReorderedPowerOfTwo(n int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #869: Reordered Power of 2
// https://leetcode.com/problems/reordered-power-of-2/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ReorderedPowerOfTwo(1))
	fmt.Println(ReorderedPowerOfTwo(10))
	fmt.Println(ReorderedPowerOfTwo(46))
}

// Time: O(log n) | Space: O(1)
func ReorderedPowerOfTwo(n int) bool {
	sig := signature(n)
	for i := 1; i <= 1_000_000_000; i <<= 1 {
		if signature(i) == sig {
			return true
		}
	}
	return false
}

func signature(x int) [10]int {
	var cnt [10]int
	for x > 0 {
		cnt[x%10]++
		x /= 10
	}
	return cnt
}
```
