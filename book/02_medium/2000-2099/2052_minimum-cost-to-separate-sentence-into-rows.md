# 2052 — Minimum Cost To Separate Sentence Into Rows

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumCost(sentence string, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n * k)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2052: Minimum Cost to Separate Sentence Into Rows
// https://leetcode.com/problems/minimum-cost-to-separate-sentence-into-rows/
// Difficulty: Medium [Paid]
// Time: O(n * k) | Space: O(n)

import "fmt"

func minimumCost(sentence string, k int) int {
	words := []string{}
	start := 0
	for i := 0; i <= len(sentence); i++ {
		if i == len(sentence) || sentence[i] == ' ' {
			words = append(words, sentence[start:i])
			start = i + 1
		}
	}

	n := len(words)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	// Check for word longer than k
	for _, w := range words {
		if len(w) > k {
			return -1 // impossible
		}
	}

	// DP[i] = min cost to place words[i:]
  // Alokasi slice
	dp := make([]int, n+1)
	dp[n] = 0

	for i := n - 1; i >= 0; i-- {
		lineLen := len(words[i])
		best := int(1e9)
		j := i
		for j < n && lineLen <= k {
			if j == n-1 {
				// Last word, no extra cost
				if 0 < best {
					best = 0
				}
			} else {
				cost := (k - lineLen) * (k - lineLen)
				total := cost + dp[j+1]
				if total < best {
					best = total
				}
			}
			j++
			if j < n {
				lineLen += 1 + len(words[j]) // space + word
			}
		}
		dp[i] = best
	}

	return dp[0]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumCost("hello world", 10))
	// Expected: 0 (both words fit on first line, last line no cost)

	// Test case 2
	fmt.Println("Test 2:", minimumCost("a b c d e", 2))
	// Expected: varies based on cost calculation

	// Test case 3
	fmt.Println("Test 3:", minimumCost("hello", 5))
	// Expected: 0
}
```
