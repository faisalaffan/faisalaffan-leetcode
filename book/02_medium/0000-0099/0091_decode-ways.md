# 0091 — Decode Ways

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numDecodings(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #91: Decode Ways
// https://leetcode.com/problems/decode-ways/
// Difficulty: Medium

import "fmt"

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	n := len(s)
  // Alokasi slice integer
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		oneDigit := int(s[i-1] - '0')
		if oneDigit >= 1 {
			dp[i] += dp[i-1]
		}

		twoDigits := int(s[i-2]-'0')*10 + oneDigit
		if twoDigits >= 10 && twoDigits <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func main() {
	// Test case 1
	fmt.Println(numDecodings("12")) // 2

	// Test case 2
	fmt.Println(numDecodings("226")) // 3

	// Test case 3
	fmt.Println(numDecodings("06")) // 0
}

// Time: O(n) | Space: O(n)
```
