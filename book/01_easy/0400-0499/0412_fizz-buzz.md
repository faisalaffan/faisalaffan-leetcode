# 0412 — Fizz Buzz

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FizzBuzz(n int) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #412: Fizz Buzz
// https://leetcode.com/problems/fizz-buzz/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

// Time: O(n), Space: O(n)
func FizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result[i-1] = "FizzBuzz"
		case i%3 == 0:
			result[i-1] = "Fizz"
		case i%5 == 0:
			result[i-1] = "Buzz"
		default:
			result[i-1] = strconv.Itoa(i)
		}
	}
	return result
}

func main() {
	fmt.Println(FizzBuzz(3))
	fmt.Println(FizzBuzz(5))
	fmt.Println(FizzBuzz(15))
}
```
