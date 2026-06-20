# 3290 — Maximum Multiplication Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxScore(a []int, b []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n) Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3290: Maximum Multiplication Score
// https://leetcode.com/problems/maximum-multiplication-score/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import (
	"fmt"
)

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4}, []int{5, 6, 7, 8}))          // 70
	fmt.Println(maxScore([]int{-1, -2, -3, -4}, []int{1, 2, 3, 4}))      // -20
	fmt.Println(maxScore([]int{3, 2, 1, 4}, []int{2, 3, 4, 5, 6}))       // 52
}

func maxScore(a []int, b []int) int64 {
	const negInf int64 = -1e18
	dp := [4]int64{negInf, negInf, negInf, negInf}

	for _, bi := range b {
		for i := 3; i >= 0; i-- {
			var prev int64
			if i > 0 {
				prev = dp[i-1]
			}
			val := prev + int64(a[i])*int64(bi)
			if val > dp[i] {
				dp[i] = val
			}
		}
	}

	return dp[3]
}
```
