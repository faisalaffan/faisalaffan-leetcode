# 2746 — Decremental String Concatenation

## Deskripsi

**Soal:** [2746. Decremental String Concatenation](https://leetcode.com/problems/decremental-string-concatenation/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func DecrementalStringConcatenation(words []string) int`

## Solusi Go

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
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, 26)
  // Iterasi seluruh elemen
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
  // Membuat slice 2D untuk DP/tabel
		ndp := make([][]int, 26)
  // Iterasi seluruh elemen
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
