# 1230 — Toss Strange Coins

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func probabilityOfHeads(prob []float64, target int) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n * target)  
**Kompleksitas Ruang:** O(target)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1230: Toss Strange Coins
// https://leetcode.com/problems/toss-strange-coins/
// Difficulty: Medium [Paid]

// Probability that exactly target coins land heads.
// dp[j] = probability of j heads after processing i coins.

// Time: O(n * target)
// Space: O(target)

func probabilityOfHeads(prob []float64, target int) float64 {
	n := len(prob)
	dp := make([]float64, target+1)
	dp[0] = 1.0

	for i := 0; i < n; i++ {
		for j := min(target, i+1); j >= 0; j-- {
			if j > 0 {
				dp[j] = dp[j-1]*prob[i] + dp[j]*(1-prob[i])
			} else {
				dp[0] = dp[0] * (1 - prob[i])
			}
		}
	}

	return dp[target]
}

func main() {
	fmt.Printf("%.5f (expected: 0.40000)\n",
		probabilityOfHeads([]float64{0.4}, 1))

	fmt.Printf("%.5f (expected: 0.40000)\n",
		probabilityOfHeads([]float64{0.5, 0.5, 0.5, 0.5, 0.5}, 0))
}
```
