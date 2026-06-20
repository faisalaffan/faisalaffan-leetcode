# 3332 — Maximum Points Tourist Can Earn

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxScore(n int, k int, stayScore [][]int, travelScore [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(k * n^2) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3332: Maximum Points Tourist Can Earn
// https://leetcode.com/problems/maximum-points-tourist-can-earn/
// Difficulty: Medium
// Time: O(k * n^2) Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxScore(2, 1, [][]int{{2, 3}}, [][]int{{0, 2}, {1, 0}})) // 3
	fmt.Println(maxScore(3, 2, [][]int{{3, 4, 2}, {2, 1, 3}}, [][]int{{0, 2, 1}, {2, 0, 2}, {1, 3, 0}})) // 8
}

func maxScore(n int, k int, stayScore [][]int, travelScore [][]int) int {
  // Alokasi slice integer
	dp := make([]int, n)
	for i := 0; i < k; i++ {
  // Alokasi slice integer
		ndp := make([]int, n)
		// Copy dp and add stay score
		for curr := 0; curr < n; curr++ {
			ndp[curr] = dp[curr] + stayScore[i][curr]
		}
		// Try travel from any city to any city
		for curr := 0; curr < n; curr++ {
			for dest := 0; dest < n; dest++ {
				val := dp[curr] + travelScore[curr][dest]
				if val > ndp[dest] {
					ndp[dest] = val
				}
			}
		}
		dp = ndp
	}

	maxVal := 0
	for _, v := range dp {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
```
