# 2270 — Number Of Ways To Split Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func waysToSplitArray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2270: Number of Ways to Split Array
// https://leetcode.com/problems/number-of-ways-to-split-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func waysToSplitArray(nums []int) int {
	n := len(nums)
	total := int64(0)
	for _, v := range nums {
		total += int64(v)
	}

	var prefix int64 = 0
	ways := 0
	for i := 0; i < n-1; i++ {
		prefix += int64(nums[i])
		if prefix >= total-prefix {
			ways++
		}
	}
	return ways
}

func main() {
	// Test case 1
	fmt.Println(waysToSplitArray([]int{10, 4, -8, 7}))
	// Expected: 2

	// Test case 2
	fmt.Println(waysToSplitArray([]int{2, 3, 1, 0}))
	// Expected: 2

	// Test case 3
	fmt.Println(waysToSplitArray([]int{-1, -2, -3, -4}))
	// Expected: 0
}
```
