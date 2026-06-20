# 0667 — Beautiful Arrangement Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func constructArray(n int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #667: Beautiful Arrangement II
// https://leetcode.com/problems/beautiful-arrangement-ii/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(constructArray(3, 1))
	fmt.Println(constructArray(3, 2))
}

func constructArray(n int, k int) []int {
  // Alokasi slice integer
	result := make([]int, n)
	left, right := 1, k+1

	for i := 0; i <= k; i++ {
		if i%2 == 0 {
			result[i] = left
			left++
		} else {
			result[i] = right
			right--
		}
	}

	for i := k + 1; i < n; i++ {
		result[i] = i + 1
	}

	return result
}
```
