# 0135 — Candy

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func candy(ratings []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #135: Candy
// https://leetcode.com/problems/candy/
// Difficulty: Hard

import (
	"fmt"
)

func candy(ratings []int) int {
	n := len(ratings)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}

  // Alokasi slice integer
	candies := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range candies {
		candies[i] = 1
	}

	// Left to right
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// Right to left
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	total := 0
	for _, c := range candies {
		total += c
	}
	return total
}

func main() {
	ratings := []int{1, 0, 2}
	result := candy(ratings)
	expected := 5

	fmt.Printf("candy(%v) = %d\n", ratings, result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```
