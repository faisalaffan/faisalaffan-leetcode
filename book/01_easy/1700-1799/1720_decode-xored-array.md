# 1720 — Decode Xored Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Decode(encoded []int, first int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1720: Decode XORed Array
// https://leetcode.com/problems/decode-xored-array/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func Decode(encoded []int, first int) []int {
  // Alokasi slice integer
	result := make([]int, len(encoded)+1)
	result[0] = first
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(encoded); i++ {
		result[i+1] = result[i] ^ encoded[i]
	}
	return result
}

func main() {
	fmt.Println(Decode([]int{1, 2, 3}, 1))
	fmt.Println(Decode([]int{6, 2, 7, 3}, 4))
}
```
