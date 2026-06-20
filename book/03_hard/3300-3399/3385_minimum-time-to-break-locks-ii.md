# 3385 — Minimum Time To Break Locks Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumTimeToBreakLocksIi(strength []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3385: Minimum Time to Break Locks II
// https://leetcode.com/problems/minimum-time-to-break-locks-ii/
// Difficulty: Hard [Paid]
//
// Bitmask DP. Each lock broken adds 1 to power multiplier.
// time = ceil(strength[i] / power).

import "fmt"

func main() {
	fmt.Println(MinimumTimeToBreakLocksIi([]int{3, 4, 1}, 2))
}

func MinimumTimeToBreakLocksIi(strength []int, k int) int {
	n := len(strength)
	m := 1 << n
  // Alokasi slice integer
	dp := make([]int, m)
	for mask := 1; mask < m; mask++ {
		dp[mask] = 1 << 60
	}

	dp[0] = 0
	for mask := 0; mask < m; mask++ {
		broken := 0
		for b := 0; b < n; b++ {
			if mask>>b&1 == 1 {
				broken++
			}
		}
		power := k + broken
		for b := 0; b < n; b++ {
			if mask>>b&1 == 1 {
				continue
			}
			need := (strength[b] + power - 1) / power
			nmask := mask | (1 << b)
			if dp[mask]+need < dp[nmask] {
				dp[nmask] = dp[mask] + need
			}
		}
	}
	return dp[m-1]
}
```
