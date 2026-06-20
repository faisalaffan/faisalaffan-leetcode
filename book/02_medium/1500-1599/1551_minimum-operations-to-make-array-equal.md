# 1551 — Minimum Operations To Make Array Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinOperations(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1551: Minimum Operations to Make Array Equal
// https://leetcode.com/problems/minimum-operations-to-make-array-equal/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinOperations(3))
	fmt.Println(MinOperations(6))
	fmt.Println(MinOperations(1))
}

func MinOperations(n int) int {
	// Time: O(1), Space: O(1)
	// arr[i] = 2*i + 1 for i in [0..n-1]
	// Target value = n (average of first n odd numbers)
	// Total operations = sum of (target - arr[i]) for arr[i] < target
	// = sum of (n - (2*i+1)) for i where 2*i+1 < n
	// = sum of (n - 2*i - 1) for i < n/2
	//
	// For n even: n=2k, sum_{i=0}^{k-1} (2k-2i-1) = k^2
	// For n odd: n=2k+1, sum_{i=0}^{k-1} (2k+1-2i-1) = k*(k+1) = k^2 + k = k(k+1)

	return n * n / 4
}
```
