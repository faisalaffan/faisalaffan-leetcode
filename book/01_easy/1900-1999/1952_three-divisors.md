# 1952 — Three Divisors

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ThreeDivisors(n int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(sqrt(n)), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1952: Three Divisors
// https://leetcode.com/problems/three-divisors/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ThreeDivisors(2))  // false
	fmt.Println(ThreeDivisors(4))  // true
	fmt.Println(ThreeDivisors(81)) // false
}

// Time: O(sqrt(n)), Space: O(1)
func ThreeDivisors(n int) bool {
	// n has exactly 3 divisors iff n is a perfect square of a prime
	if n < 4 {
		return false
	}

	// Check if sqrt(n) is integer
	root := 1
	for root*root < n {
		root++
	}
	if root*root != n {
		return false
	}

	// Check if root is prime
	for i := 2; i*i <= root; i++ {
		if root%i == 0 {
			return false
		}
	}
	return root > 1
}
```
