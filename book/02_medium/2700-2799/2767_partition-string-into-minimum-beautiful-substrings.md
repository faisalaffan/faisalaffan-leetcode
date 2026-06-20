# 2767 — Partition String Into Minimum Beautiful Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func PartitionStringIntoMinimumBeautifulSubstrings(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP

**Waktu:** O(2^n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2767: Partition String Into Minimum Beautiful Substrings
// https://leetcode.com/problems/partition-string-into-minimum-beautiful-substrings/
// Difficulty: Medium
// Time: O(2^n) | Space: O(n)

import (
	"fmt"
	"math"
)

func PartitionStringIntoMinimumBeautifulSubstrings(s string) int {
	n := len(s)
  // HashMap: O(1) lookup
	powers := make(map[string]bool)
	for i := 0; i <= 10; i++ {
		p := int(math.Pow(5, float64(i)))
		b := fmt.Sprintf("%b", p)
		if len(b) <= 15 {
			powers[b] = true
		}
	}

  // Alokasi slice
	memo := make([]int, n)
  // Range loop
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int
	dfs = func(start int) int {
		if start == n {
			return 0
		}
		if memo[start] != -1 {
			return memo[start]
		}
		best := math.MaxInt32
		for end := start + 1; end <= n; end++ {
			sub := s[start:end]
			if powers[sub] {
				subResult := dfs(end)
				if subResult != -1 && subResult+1 < best {
					best = subResult + 1
				}
			}
		}
		if best == math.MaxInt32 {
			memo[start] = -1
		} else {
			memo[start] = best
		}
		return memo[start]
	}

	result := dfs(0)
	if result == math.MaxInt32 || result <= 0 {
		return -1
	}
	return result
}

func main() {
	fmt.Println(PartitionStringIntoMinimumBeautifulSubstrings("1011"))
	fmt.Println(PartitionStringIntoMinimumBeautifulSubstrings("111"))
}
```
