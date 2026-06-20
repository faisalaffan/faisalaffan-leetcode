# 3875 — Construct Uniform Parity Array I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ConstructUniformParityArrayI(nums1 []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3875: Construct Uniform Parity Array I
// https://leetcode.com/problems/construct-uniform-parity-array-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ConstructUniformParityArrayI([]int{2, 3}))
	fmt.Println(ConstructUniformParityArrayI([]int{4, 6}))
}

// Time: O(1)
// Space: O(1)
func ConstructUniformParityArrayI(nums1 []int) bool {
	return true
}
```
