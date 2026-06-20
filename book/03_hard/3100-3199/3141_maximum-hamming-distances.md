# 3141 — Maximum Hamming Distances

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxHammingDistances(nums []int, m int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3141: Maximum Hamming Distances
// https://leetcode.com/problems/maximum-hamming-distances/
// Difficulty: Hard [Paid]
//
// Given an array of integers, for each element find the maximum Hamming distance
// (number of differing bits) to any other element in the array.
//
// HammingDist(x, y) = popcount(x ^ y)
// For an element x, max distance = m - min distance from complement to the set.
// Use DP/bitmask: dp[mask] = min Hamming distance from mask to any number in nums.

import (
	"fmt"
	"math"
)

func maxHammingDistances(nums []int, m int) []int {
	size := 1 << m
  // Alokasi slice integer
	dp := make([]int, size)
	for i := 0; i < size; i++ {
		dp[i] = math.MaxInt32
	}

	// Set distance 0 for all numbers present
	for _, v := range nums {
		dp[v] = 0
	}

	// DP: for each bit, relax distances
	for i := 0; i < m; i++ {
		bit := 1 << i
		for mask := 0; mask < size; mask++ {
			if dp[mask^bit]+1 < dp[mask] {
				dp[mask] = dp[mask^bit] + 1
			}
		}
	}

	fullMask := size - 1
  // Alokasi slice integer
	ans := make([]int, len(nums))
	for idx, v := range nums {
		comp := fullMask ^ v
		ans[idx] = m - dp[comp]
	}
	return ans
}

func main() {
	// Test case 1: 2-bit numbers
	nums := []int{0, 1, 3}
	fmt.Println("Test 1:", maxHammingDistances(nums, 2))
	// Expected: [2, 2, 2] (0^3=2=popcount 2, 1^3=2=popcount 2, 3^0=2=popcount 2)

	// Test case 2: 3-bit numbers
	nums2 := []int{0, 7}
	fmt.Println("Test 2:", maxHammingDistances(nums2, 3))
	// Expected: [3, 3] (0^7=7=popcount 3)

	// Test case 3: single element
	nums3 := []int{5}
	fmt.Println("Test 3:", maxHammingDistances(nums3, 3))
	// Expected: [0] (no other element to compare)

	// Test case 4: 4-bit
	nums4 := []int{0, 1, 2, 4}
	fmt.Println("Test 4:", maxHammingDistances(nums4, 3))
	// Expected: various distances
}
```
