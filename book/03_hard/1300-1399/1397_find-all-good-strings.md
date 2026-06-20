# 1397 — Find All Good Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfGoodStrings(n int, s1 string, s2 string, evil string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1397: Find All Good Strings
// https://leetcode.com/problems/find-all-good-strings/
// Difficulty: Hard

import "fmt"

const mod = 1_000_000_007

// numberOfGoodStrings returns the number of strings of length n
// that are lexicographically between s1 and s2 (inclusive) and
// do not contain evil as a substring.
func numberOfGoodStrings(n int, s1 string, s2 string, evil string) int {
	// Build KMP LPS array for evil string
	m := len(evil)
  // Alokasi slice integer
	lps := make([]int, m)
	for i := 1; i < m; i++ {
		j := lps[i-1]
		for j > 0 && evil[i] != evil[j] {
			j = lps[j-1]
		}
		if evil[i] == evil[j] {
			j++
		}
		lps[i] = j
	}

	// nextState[k][c] = next KMP state after adding character c in state k
  // Alokasi slice integer
	nextState := make([][26]int, m)
	for k := 0; k < m; k++ {
		for c := 0; c < 26; c++ {
			if k < m && int(evil[k]-'a') == c {
				nextState[k][c] = k + 1
			} else if k == 0 {
				nextState[k][c] = 0
			} else {
				j := k
				for j > 0 && int(evil[j]-'a') != c {
					j = lps[j-1]
				}
				if int(evil[j]-'a') == c {
					j++
				}
				nextState[k][c] = j
			}
		}
	}

	// dp[pos][state][tightLow][tightHigh]
	// Count strings from position pos to end
  // Membuat matriks/slice 2D untuk DP
	dp := make([][][2][2]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([][2][2]int, m+1)
	}

	// Initialize dp for pos == n (past the end)
	for state := 0; state <= m; state++ {
		dp[n][state][0][0] = 1
		dp[n][state][0][1] = 1
		dp[n][state][1][0] = 1
		dp[n][state][1][1] = 1
	}

	for pos := n - 1; pos >= 0; pos-- {
		for state := 0; state < m; state++ {
			for tightLow := 0; tightLow <= 1; tightLow++ {
				for tightHigh := 0; tightHigh <= 1; tightHigh++ {
					lowChar := byte('a')
					if tightLow == 1 {
						lowChar = s1[pos]
					}
					highChar := byte('z')
					if tightHigh == 1 {
						highChar = s2[pos]
					}

					var total int
					for c := lowChar; c <= highChar; c++ {
						nextSt := nextState[state][c-'a']
						if nextSt == m {
							// This path would contain evil substring, skip
							continue
						}
						nextTightLow := tightLow
						if c > s1[pos] {
							nextTightLow = 0
						}
						nextTightHigh := tightHigh
						if c < s2[pos] {
							nextTightHigh = 0
						}
						total = (total + dp[pos+1][nextSt][nextTightLow][nextTightHigh]) % mod
					}
					dp[pos][state][tightLow][tightHigh] = total
				}
			}
		}
	}

	return dp[0][0][1][1]
}

func main() {
	// Test case 1
	n1, s1_1, s2_1, evil1 := 2, "aa", "da", "b"
	result1 := numberOfGoodStrings(n1, s1_1, s2_1, evil1)
	fmt.Printf("Test 1: n=%d, s1=%q, s2=%q, evil=%q => %d (expected 51)\n", n1, s1_1, s2_1, evil1, result1)

	// Test case 2
	n2, s1_2, s2_2, evil2 := 3, "aa", "az", "b"
	result2 := numberOfGoodStrings(n2, s1_2, s2_2, evil2)
	fmt.Printf("Test 2: n=%d, s1=%q, s2=%q, evil=%q => %d\n", n2, s1_2, s2_2, evil2, result2)

	// Test case 3
	n3, s1_3, s2_3, evil3 := 8, "leetcode", "leetgoes", "leet"
	result3 := numberOfGoodStrings(n3, s1_3, s2_3, evil3)
	fmt.Printf("Test 3: n=%d, s1=%q, s2=%q, evil=%q => %d\n", n3, s1_3, s2_3, evil3, result3)

	// Test case 4: from LeetCode example 1
	n4, s1_4, s2_4, evil4 := 2, "aa", "bb", "ab"
	result4 := numberOfGoodStrings(n4, s1_4, s2_4, evil4)
	fmt.Printf("Test 4: n=%d, s1=%q, s2=%q, evil=%q => %d (expected 0)\n", n4, s1_4, s2_4, evil4, result4)

	// Test case 5: small edge case
	n5, s1_5, s2_5, evil5 := 1, "a", "c", "b"
	result5 := numberOfGoodStrings(n5, s1_5, s2_5, evil5)
	fmt.Printf("Test 5: n=%d, s1=%q, s2=%q, evil=%q => %d (expected 2: \"a\", \"c\")\n", n5, s1_5, s2_5, evil5, result5)
}
```
