# 0372 — Super Pow

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func superPow(a int, b []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #372: Super Pow
// https://leetcode.com/problems/super-pow/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

const mod = 1337

func superPow(a int, b []int) int {
	a %= mod
	result := 1

	for _, digit := range b {
		result = (powMod(result, 10) * powMod(a, digit)) % mod
	}
	return result
}

func powMod(base, exp int) int {
	result := 1
	base %= mod
	for i := 0; i < exp; i++ {
		result = (result * base) % mod
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", superPow(2, []int{3}))
	// Expected: 8

	// Test case 2
	fmt.Println("Test 2:", superPow(2, []int{1, 0}))
	// Expected: 1024

	// Test case 3
	fmt.Println("Test 3:", superPow(1, []int{4, 3, 3, 8, 5, 2}))
	// Expected: 1
}
```
