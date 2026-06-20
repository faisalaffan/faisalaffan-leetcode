# 0334 — Increasing Triplet Subsequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func increasingTriplet(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Bitmask

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Bitmask** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #334: Increasing Triplet Subsequence
// https://leetcode.com/problems/increasing-triplet-subsequence/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func increasingTriplet(nums []int) bool {
	first, second := 1<<31-1, 1<<31-1
	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			return true
		}
	}
	return false
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", increasingTriplet([]int{1, 2, 3, 4, 5}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", increasingTriplet([]int{5, 4, 3, 2, 1}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
	// Expected: true
}
```
