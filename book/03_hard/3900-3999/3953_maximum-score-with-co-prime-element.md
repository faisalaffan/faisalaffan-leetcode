# 3953 — Maximum Score With Co Prime Element

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumScore(nums []int, maxVal int) int
```

> **💡 Hint:** For each value <= maxVal, compute its score (how many

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, GCD / Matematika, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3953: Maximum Score with Co-Prime Element
// https://leetcode.com/problems/maximum-score-with-co-prime-element/
// Difficulty: Hard
//
// Given array nums and a maxVal, select any number of elements
// such that selected elements are pairwise co-prime (gcd = 1).
// Maximize sum of selected elements, each <= maxVal.
//
// Approach: For each value <= maxVal, compute its score (how many
// times it appears in nums). Use DP over subsets of prime factors
// (or use inclusion-exclusion with prime mask). Since maxVal is
// small, iterate over values.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maximumScore([]int{2, 3, 5}, 10))
	// Example 2
	fmt.Println(maximumScore([]int{4, 6, 8}, 10))
	// Edge: single value
	fmt.Println(maximumScore([]int{7}, 10))
}

func maximumScore(nums []int, maxVal int) int {
	// Count frequency of each value
  // Alokasi slice integer
	freq := make([]int, maxVal+1)
	for _, v := range nums {
		if v <= maxVal {
			freq[v]++
		}
	}

	// For each value, compute its "score" = value * frequency
  // Alokasi slice integer
	score := make([]int, maxVal+1)
	for v := 1; v <= maxVal; v++ {
		score[v] = v * freq[v]
	}

	// DP over values: for each value, can we pick it with picked set?
	// Since we need pairwise coprime, we track prime factor mask.
	// Precompute prime mask for each value
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
  // Alokasi slice integer
	maskOf := make([]int, maxVal+1)
	for v := 2; v <= maxVal; v++ {
		m := 0
		tmp := v
		for pi, p := range primes {
			if tmp%p == 0 {
				m |= (1 << pi)
				for tmp%p == 0 {
					tmp /= p
				}
			}
		}
		maskOf[v] = m
	}

	// dp[mask] = max score with prime factors covered by mask
  // Alokasi slice integer
	dp := make([]int, 1<<len(primes))
	for v := 1; v <= maxVal; v++ {
		if score[v] == 0 {
			continue
		}
		m := maskOf[v]
		// Try to add v to any existing set that doesn't share primes
		for mask := (1 << len(primes)) - 1; mask >= 0; mask-- {
			if mask&m == 0 {
				nm := mask | m
				if dp[mask]+score[v] > dp[nm] {
					dp[nm] = dp[mask] + score[v]
				}
			}
		}
	}

	ans := 0
	for _, v := range dp {
		if v > ans {
			ans = v
		}
	}
	return ans
}
```
