# 0087 — Scramble String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func isScramble(s1 string, s2 string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #87: Scramble String
// https://leetcode.com/problems/scramble-string/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("87. Scramble String")
	fmt.Println("great, rgeat ->", isScramble("great", "rgeat"), "(expected true)")
	fmt.Println("abcde, caebd ->", isScramble("abcde", "caebd"), "(expected false)")
	fmt.Println("a, a ->", isScramble("a", "a"), "(expected true)")
	fmt.Println("ab, ba ->", isScramble("ab", "ba"), "(expected true)")
}

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}

	n := len(s1)

	// memo[key] where key encodes (s1_start, s1_end, s2_start, s2_end)
	// but we can use a simpler DP: dp[k][i][j] = isScramble(s1[i:i+k], s2[j:j+k])
  // Membuat matriks/slice 2D untuk DP
	dp := make([][][]bool, n+1)
	for k := 0; k <= n; k++ {
		dp[k] = make([][]bool, n)
		for i := 0; i < n; i++ {
			dp[k][i] = make([]bool, n)
		}
	}

	// k=1 base case
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[1][i][j] = s1[i] == s2[j]
		}
	}

	// k from 2 to n
	for k := 2; k <= n; k++ {
		for i := 0; i <= n-k; i++ {
			for j := 0; j <= n-k; j++ {
				for split := 1; split < k; split++ {
					// no swap: s1[i:i+split] ~ s2[j:j+split] && s1[i+split:i+k] ~ s2[j+split:j+k]
					if dp[split][i][j] && dp[k-split][i+split][j+split] {
						dp[k][i][j] = true
						break
					}
					// swap: s1[i:i+split] ~ s2[j+k-split:j+k] && s1[i+split:i+k] ~ s2[j:j+k-split]
					if dp[split][i][j+k-split] && dp[k-split][i+split][j] {
						dp[k][i][j] = true
						break
					}
				}
			}
		}
	}

	return dp[n][0][0]
}
```
