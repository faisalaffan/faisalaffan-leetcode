# 1434 — Number Of Ways To Wear Different Hats To Each Other

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberWays(hats [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1434: Number of Ways to Wear Different Hats to Each Other
// https://leetcode.com/problems/number-of-ways-to-wear-different-hats-to-each-other/
// Difficulty: Hard

import "fmt"

const mod1434 = 1_000_000_007

func numberWays(hats [][]int) int {
	n := len(hats)
	// Map each hat (1..40) to people who like it
  // Membuat matriks/slice 2D untuk DP
	hatToPeople := make([][]int, 41)
	for person, list := range hats {
		for _, hat := range list {
			hatToPeople[hat] = append(hatToPeople[hat], person)
		}
	}

	totalMasks := 1 << n
  // Alokasi slice integer
	dp := make([]int, totalMasks)
	dp[0] = 1

	for hat := 1; hat <= 40; hat++ {
		if len(hatToPeople[hat]) == 0 {
			continue
		}
		// Iterate masks in reverse to avoid reusing the same hat
		for mask := totalMasks - 1; mask >= 0; mask-- {
			for _, person := range hatToPeople[hat] {
				if mask&(1<<person) != 0 {
					continue
				}
				nextMask := mask | (1 << person)
				dp[nextMask] = (dp[nextMask] + dp[mask]) % mod1434
			}
		}
	}
	return dp[totalMasks-1]
}

func main() {
	// Example: hats = [[3,4],[4,5],[5]] -> 1
	fmt.Println(numberWays([][]int{{3, 4}, {4, 5}, {5}}))
}
```
