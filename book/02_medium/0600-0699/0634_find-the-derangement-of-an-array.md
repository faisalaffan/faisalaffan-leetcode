# 0634 — Find The Derangement Of An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindDerangement(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #634: Find the Derangement of An Array
// https://leetcode.com/problems/find-the-derangement-of-an-array/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(FindDerangement(3))
	fmt.Println(FindDerangement(4))
}

func FindDerangement(n int) int {
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 0
	}

	const mod = 1_000_000_007
	a, b := 0, 1 // D(1)=0, D(2)=1

	for i := 3; i <= n; i++ {
		c := ((i - 1) * (a + b)) % mod
		a, b = b, c
	}

	return b
}
```
