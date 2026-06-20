# 1929 — Concatenation Of Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ConcatenationOfArray(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1929: Concatenation of Array
// https://leetcode.com/problems/concatenation-of-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ConcatenationOfArray([]int{1, 2, 1}))       // [1,2,1,1,2,1]
	fmt.Println(ConcatenationOfArray([]int{1, 3, 2, 1}))    // [1,3,2,1,1,3,2,1]
}

// Time: O(n), Space: O(n)
func ConcatenationOfArray(nums []int) []int {
	n := len(nums)
  // Alokasi slice integer
	ans := make([]int, 2*n)
	for i, v := range nums {
		ans[i] = v
		ans[i+n] = v
	}
	return ans
}
```
