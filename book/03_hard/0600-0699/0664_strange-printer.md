# 0664 — Strange Printer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func strangePrinter(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #664: Strange Printer
// https://leetcode.com/problems/strange-printer/
// Difficulty: Hard

import "fmt"

func main() {
	// Test cases
	testCases := []struct {
		s    string
		want int
	}{
		{"aaabbb", 2},
		{"aba", 2},
		{"", 0},
		{"a", 1},
		{"aa", 1},
		{"ab", 2},
		{"abcabc", 5},
		{"aaabbaaa", 2},
		{"ababab", 4},
		{"leetcode", 6},
		{"tbgtgb", 4},
		{"aaaaaaaaaaaaaaaaaaaa", 1},
	}

	for _, tc := range testCases {
		got := strangePrinter(tc.s)
		status := "PASS"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("%s: strangePrinter(%q) = %d (want %d)\n", status, tc.s, got, tc.want)
	}
}

func strangePrinter(s string) int {
	n := len(s)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}

	// dp[i][j] = minimum turns to print s[i..j]
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1 // single character needs 1 turn
	}

	// Process intervals by length
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1

			// Worst case: print last char separately
			dp[i][j] = dp[i][j-1] + 1

			// Try to merge with a matching character
			for k := i; k < j; k++ {
				if s[k] == s[j] {
					cost := dp[i][k] + dp[k+1][j-1]
					if cost < dp[i][j] {
						dp[i][j] = cost
					}
				}
			}
		}
	}

	return dp[0][n-1]
}
```
