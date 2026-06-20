# 2501 — Longest Square Streak In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func longestSquareStreak(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Dynamic Programming

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2501: Longest Square Streak in an Array
// https://leetcode.com/problems/longest-square-streak-in-an-array/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)
// Sort, use map to track longest streak ending at each value.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestSquareStreak([]int{4, 3, 6, 16, 8, 2})) // 3 (2 -> 4 -> 16)
	fmt.Println(longestSquareStreak([]int{2, 3, 5, 6, 7}))     // -1
}

func longestSquareStreak(nums []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
  // Membuat map (HashMap) — pencarian O(1)
	dp := make(map[int]int)
	ans := -1

	for _, v := range nums {
		root := intSqrt(v)
		if root*root == v {
			if prev, ok := dp[root]; ok {
				dp[v] = prev + 1
			} else {
				dp[v] = 1
			}
		} else {
			dp[v] = 1
		}
		if dp[v] > ans {
			ans = dp[v]
		}
	}
	if ans < 2 {
		return -1
	}
	return ans
}

func intSqrt(n int) int {
	lo, hi := 1, n
	for lo <= hi {
		mid := (lo + hi) / 2
		if mid*mid == n {
			return mid
		} else if mid*mid < n {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return 0
}
```
