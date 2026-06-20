# 3634 — Minimum Removals To Balance Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumRemovalsToBalanceArray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3634: Minimum Removals to Balance Array
// https://leetcode.com/problems/minimum-removals-to-balance-array/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MinimumRemovalsToBalanceArray([]int{1, 2, 3, 4, 5}))
	// Test case 2
	fmt.Println("Test 2:", MinimumRemovalsToBalanceArray([]int{1, 1, 1, 1}))
	// Test case 3
	fmt.Println("Test 3:", MinimumRemovalsToBalanceArray([]int{5, 4, 3, 2, 1}))
}

func MinimumRemovalsToBalanceArray(nums []int) int {
	// Minimum removals so that sum of first half equals sum of second half
	n := len(nums)
	if n%2 != 0 {
		// Remove middle element for odd length
		return 1
	}
	half := n / 2
	sum1, sum2 := 0, 0
	for i := 0; i < half; i++ {
		sum1 += nums[i]
	}
	for i := half; i < n; i++ {
		sum2 += nums[i]
	}
	if sum1 == sum2 {
		return 0
	}
	// Remove one element from the larger half
	return 1
}
```
