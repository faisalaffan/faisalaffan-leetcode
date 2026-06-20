# 3790 — Smallest All Ones Multiple

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func smallestAllOnesMultiple(k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(k)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3790: Smallest All-Ones Multiple
// https://leetcode.com/problems/smallest-all-ones-multiple/
// Difficulty: Medium
// Time: O(k) | Space: O(1)

import "fmt"

func smallestAllOnesMultiple(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	rem := 0
	for n := 1; n <= k; n++ {
		rem = (rem*10 + 1) % k
		if rem == 0 {
			return n
		}
	}
	return -1
}

func main() {
	fmt.Println(smallestAllOnesMultiple(3))
	fmt.Println(smallestAllOnesMultiple(7))
	fmt.Println(smallestAllOnesMultiple(2))
}
```
