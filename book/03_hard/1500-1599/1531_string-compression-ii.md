# 1531 — String Compression Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func getLengthOfOptimalCompression(s string, k int) int
```

> **💡 Hint:** DP (Top-Down with memoization)

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1531: String Compression II
// https://leetcode.com/problems/string-compression-ii/
// Difficulty: Hard
//
// Approach: DP (Top-Down with memoization)
// dp[i][k] = minimum compressed length for s[i:] with at most k deletions.
// At position i with character c = s[i]:
//   - Option 1: delete s[i] (if k > 0): dp[i+1][k-1]
//   - Option 2: keep s[i], count consecutive same chars starting at i.
//     For each run length len from 1 to n-i (stopping when we either run out of
//     deletions or change character), we keep 'len' copies of c and delete the rest
//     within the run. The compressed length for this segment is:
//       1 (for the char) + (digits of len if len > 1)
//     Then add dp[i+len+deleted][k-deleted].

import (
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(getLengthOfOptimalCompression("aaabcccd", 2))
	// Expected: 4 ("aa3bccd" -> a2bccd... Let me verify)
	// Actually "aaabcccd" with 2 deletions:
	// delete a at idx 2, delete c at idx 5? Let's trust LeetCode: answer is 4.

	// Example 2
	fmt.Println(getLengthOfOptimalCompression("aabbaa", 2))
	// Expected: 2

	// Example 3
	fmt.Println(getLengthOfOptimalCompression("aaaaaaaaaaa", 0))
	// Expected: 3 ("a11" -> "a11" is 3 chars)
}

func getLengthOfOptimalCompression(s string, k int) int {
	n := len(s)
	// memo[i][k]
  // Membuat matriks/slice 2D untuk DP
	memo := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range memo {
		memo[i] = make([]int, k+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dp func(i, k int) int
	dp = func(i, k int) int {
		if i == n || k >= n-i {
			return 0
		}
		if memo[i][k] != -1 {
			return memo[i][k]
		}

		best := math.MaxInt32

		// Option 1: delete s[i]
		if k > 0 {
			best = minInt(best, dp(i+1, k-1))
		}

		// Option 2: keep s[i]
		c := s[i]
		sameCount := 0
		deleted := 0
		for j := i; j < n && deleted <= k; j++ {
			if s[j] == c {
				sameCount++
			} else {
				deleted++
			}
			if deleted > k {
				break
			}
			// Compressed length for sameCount copies of c
			compLen := compLength(sameCount)
			best = minInt(best, compLen+dp(j+1, k-deleted))
		}

		memo[i][k] = best
		return best
	}

	return dp(0, k)
}

func compLength(cnt int) int {
	// compressed = 1 char + digit length if cnt > 1
	if cnt == 1 {
		return 1
	}
	if cnt < 10 {
		return 2
	}
	if cnt < 100 {
		return 3
	}
	return 4
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
