# 1969 — Minimum Non Zero Product Of The Array Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinNonZeroProduct(p int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(p), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1969: Minimum Non-Zero Product of the Array Elements
// https://leetcode.com/problems/minimum-non-zero-product-of-the-array-elements/
// Difficulty: Medium

import "fmt"

const mod1969 = 1000000007

func main() {
	fmt.Println(MinNonZeroProduct(1))
	fmt.Println(MinNonZeroProduct(2))
	fmt.Println(MinNonZeroProduct(3))
}

// Time: O(p), Space: O(1)
func MinNonZeroProduct(p int) int {
	maxVal := (int64(1) << uint(p)) - 1
	base := maxVal - 1
	exp := (int64(1) << uint(p-1)) - 1
	result := int(maxVal % mod1969)
	result = int(int64(result) * powMod1969(base%int64(mod1969), exp) % mod1969)
	return result
}

func powMod1969(base int64, exp int64) int64 {
	result := int64(1)
	b := base % int64(mod1969)
	e := exp
	for e > 0 {
		if e&1 == 1 {
			result = (result * b) % int64(mod1969)
		}
		b = (b * b) % int64(mod1969)
		e >>= 1
	}
	return result
}
```
