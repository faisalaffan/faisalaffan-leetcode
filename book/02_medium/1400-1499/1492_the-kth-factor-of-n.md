# 1492 — The Kth Factor Of N

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func KthFactor(n int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(sqrt(N)), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1492: The kth Factor of n
// https://leetcode.com/problems/the-kth-factor-of-n/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(KthFactor(12, 3))
	fmt.Println(KthFactor(7, 2))
	fmt.Println(KthFactor(4, 4))
}

func KthFactor(n int, k int) int {
	// Time: O(sqrt(N)), Space: O(1)
	// Count factors from 1 to sqrt(n)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			k--
			if k == 0 {
				return i
			}
		}
	}

	// Count factors from sqrt(n) down to 1 (the paired factors)
	// Start from the largest paired factor
	for i := intSqrt(n); i >= 1; i-- {
		if n%i == 0 && i*i != n { // don't double count perfect square root
			k--
			if k == 0 {
				return n / i
			}
		}
	}

	return -1
}

func intSqrt(n int) int {
	for i := 1; i*i <= n; i++ {
		if i*i == n {
			return i
		}
	}
	// floor sqrt
	result := 0
	for result*result <= n {
		result++
	}
	return result - 1
}
```
