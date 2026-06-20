# 3726 — Remove Zeros In Decimal Representation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func RemoveZerosInDecimalRepresentation(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n) - number of digits  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3726: Remove Zeros in Decimal Representation
// https://leetcode.com/problems/remove-zeros-in-decimal-representation/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(RemoveZerosInDecimalRepresentation(1020030))
	fmt.Println(RemoveZerosInDecimalRepresentation(1000))
}

// Time: O(log n) - number of digits
// Space: O(1)
func RemoveZerosInDecimalRepresentation(n int) int {
	ans := 0
	k := 1
	for n > 0 {
		x := n % 10
		if x > 0 {
			ans = k*x + ans
			k *= 10
		}
		n /= 10
	}
	return ans
}
```
