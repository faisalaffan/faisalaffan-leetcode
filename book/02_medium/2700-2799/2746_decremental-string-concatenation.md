# 2746 — Decremental String Concatenation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func DecrementalStringConcatenation(words []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2746: Decremental String Concatenation
// https://leetcode.com/problems/decremental-string-concatenation/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func DecrementalStringConcatenation(words []string) int {
	n := len(words)
	// dp[first][last] = min length
	// Use 26 letters
	const INF = 1 << 30
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, 26)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, 26)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}

	f, l := int(words[0][0]-'a'), int(words[0][len(words[0])-1]-'a')
	dp[f][l] = len(words[0])

	for i := 1; i < n; i++ {
		w := words[i]
		cf, cl := int(w[0]-'a'), int(w[len(w)-1]-'a')
  // Membuat matriks/slice 2D untuk DP
		ndp := make([][]int, 26)
  // Range loop: iterasi dengan indeks + nilai
		for i := range ndp {
			ndp[i] = make([]int, 26)
			for j := range ndp[i] {
				ndp[i][j] = INF
			}
		}

		for a := 0; a < 26; a++ {
			for b := 0; b < 26; b++ {
				if dp[a][b] == INF {
					continue
				}
				// Append w: a...b + cf...cl
				if b == cf {
					if dp[a][b]+len(w)-1 < ndp[a][cl] {
						ndp[a][cl] = dp[a][b] + len(w) - 1
					}
				} else {
					if dp[a][b]+len(w) < ndp[a][cl] {
						ndp[a][cl] = dp[a][b] + len(w)
					}
				}
				// Prepend w: cf...cl + a...b
				if cl == a {
					if dp[a][b]+len(w)-1 < ndp[cf][b] {
						ndp[cf][b] = dp[a][b] + len(w) - 1
					}
				} else {
					if dp[a][b]+len(w) < ndp[cf][b] {
						ndp[cf][b] = dp[a][b] + len(w)
					}
				}
			}
		}
		dp = ndp
	}

	best := INF
	for a := 0; a < 26; a++ {
		for b := 0; b < 26; b++ {
			if dp[a][b] < best {
				best = dp[a][b]
			}
		}
	}
	return best
}

func main() {
	fmt.Println(DecrementalStringConcatenation([]string{"abc", "cde", "efg"}))
	fmt.Println(DecrementalStringConcatenation([]string{"aa", "ab", "bc"}))
}
```
