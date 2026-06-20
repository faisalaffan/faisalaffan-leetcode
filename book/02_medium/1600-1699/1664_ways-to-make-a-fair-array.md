# 1664 — Ways To Make A Fair Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func WaysToMakeFair(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1664: Ways to Make a Fair Array
// https://leetcode.com/problems/ways-to-make-a-fair-array/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(WaysToMakeFair([]int{2, 1, 6, 4}))
	fmt.Println(WaysToMakeFair([]int{1, 1, 1}))
	fmt.Println(WaysToMakeFair([]int{1, 2, 3, 4, 5}))
}

func WaysToMakeFair(nums []int) int {
	// Time: O(N), Space: O(1)

	// Calculate total sum at even and odd indices
	totalEven := 0
	totalOdd := 0
	for i, num := range nums {
		if i%2 == 0 {
			totalEven += num
		} else {
			totalOdd += num
		}
	}

	result := 0
	prefixEven := 0
	prefixOdd := 0

	for i, num := range nums {
		if i%2 == 0 {
			totalEven -= num
		} else {
			totalOdd -= num
		}

		// After removing nums[i], all indices shift:
		// Elements to the right of i swap parity
		// Even sum = prefixEven + totalOdd
		// Odd sum = prefixOdd + totalEven
		if prefixEven+totalOdd == prefixOdd+totalEven {
			result++
		}

		if i%2 == 0 {
			prefixEven += num
		} else {
			prefixOdd += num
		}
	}

	return result
}
```
