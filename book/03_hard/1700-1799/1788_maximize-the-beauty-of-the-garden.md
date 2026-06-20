# 1788 — Maximize The Beauty Of The Garden

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumBeauty(flowers []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1788: Maximize the Beauty of the Garden
// https://leetcode.com/problems/maximize-the-beauty-of-the-garden/
// Difficulty: Hard [Paid]
//
// We have flowers in a row with beauty values (positive = beautiful, negative = ugly).
// Pick a contiguous segment, then remove any subset of flowers (must keep at least one).
// Maximize the sum of kept flowers.
//
// DP solution: at each position i, dp = max(flowers[i], previous_dp + max(0, flowers[i])).
// This captures the ability to "skip" negative values within the segment while keeping
// all positives. Answer is the max dp over all positions.

import "fmt"

func main() {
	// Example 1: all positive, keep all
	fmt.Println(maximumBeauty([]int{1, 2, 3, 4}))
	// Example 2: skip the negative in middle
	fmt.Println(maximumBeauty([]int{1, -2, 3}))
	// Example 3: all negative, keep the least negative
	fmt.Println(maximumBeauty([]int{-1, -2, -3}))
	// Mixed
	fmt.Println(maximumBeauty([]int{4, -1, 3, -10, 5}))
	// Single positive
	fmt.Println(maximumBeauty([]int{7}))
	// Single negative
	fmt.Println(maximumBeauty([]int{-7}))
	// All negative various
	fmt.Println(maximumBeauty([]int{-5, -1, -3}))
	// Empty
	fmt.Println(maximumBeauty([]int{}))
	// All zeros
	fmt.Println(maximumBeauty([]int{0, 0, 0}))
	// Alternating signs
	fmt.Println(maximumBeauty([]int{-3, 2, -1, 4, -5, 6}))
}

func maximumBeauty(flowers []int) int {
	if len(flowers) == 0 {
		return 0
	}
	dp := flowers[0]
	ans := dp
	for i := 1; i < len(flowers); i++ {
		dp = max(flowers[i], dp+max(0, flowers[i]))
		ans = max(ans, dp)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
