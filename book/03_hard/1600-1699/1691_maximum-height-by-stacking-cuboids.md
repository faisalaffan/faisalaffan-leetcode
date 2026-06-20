# 1691 — Maximum Height By Stacking Cuboids

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxHeight(cuboids [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1691: Maximum Height by Stacking Cuboids
// https://leetcode.com/problems/maximum-height-by-stacking-cuboids/
// Difficulty: Hard
// Strategy: Sort each cuboid so height is the max dimension,
// then sort all cuboids and run LIS (DP).

import (
	"fmt"
	"sort"
)

func maxHeight(cuboids [][]int) int {
	// For each cuboid, sort dimensions so the largest is height
	for _, c := range cuboids {
  // Urutkan secara ascending — O(n log n)
		sort.Ints(c)
	}
	// Sort cuboids by dimensions (width, depth, height)
  // Custom sort dengan comparator
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][0] != cuboids[j][0] {
			return cuboids[i][0] < cuboids[j][0]
		}
		if cuboids[i][1] != cuboids[j][1] {
			return cuboids[i][1] < cuboids[j][1]
		}
		return cuboids[i][2] < cuboids[j][2]
	})

	n := len(cuboids)
  // Alokasi slice integer
	dp := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		dp[i] = cuboids[i][2] // height
		for j := 0; j < i; j++ {
			// Check if cuboid j can be placed below cuboid i
			if cuboids[j][0] <= cuboids[i][0] &&
				cuboids[j][1] <= cuboids[i][1] &&
				cuboids[j][2] <= cuboids[i][2] {
				if dp[j]+cuboids[i][2] > dp[i] {
					dp[i] = dp[j] + cuboids[i][2]
				}
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

func main() {
	// Example 1: [[50,45,20],[95,37,53],[45,23,12]] -> 190
	cuboids1 := [][]int{{50, 45, 20}, {95, 37, 53}, {45, 23, 12}}
	fmt.Printf("maxHeight(%v) = %d (expected 190)\n", cuboids1, maxHeight(cuboids1))

	// Example 2: [[38,25,45],[76,35,3]] -> 76
	cuboids2 := [][]int{{38, 25, 45}, {76, 35, 3}}
	fmt.Printf("maxHeight(%v) = %d (expected 76)\n", cuboids2, maxHeight(cuboids2))

	// Example 3: [[7,11,17],[7,17,11],[11,7,17],[11,17,7],[17,7,11],[17,11,7]] -> 102
	cuboids3 := [][]int{{7, 11, 17}, {7, 17, 11}, {11, 7, 17}, {11, 17, 7}, {17, 7, 11}, {17, 11, 7}}
	fmt.Printf("maxHeight(%v) = %d (expected 102)\n", cuboids3, maxHeight(cuboids3))

	// Single cuboid
	cuboids4 := [][]int{{1, 2, 3}}
	fmt.Printf("maxHeight(%v) = %d (expected 3)\n", cuboids4, maxHeight(cuboids4))
}
```
