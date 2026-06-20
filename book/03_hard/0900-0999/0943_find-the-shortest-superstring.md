# 0943 — Find The Shortest Superstring

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func shortestSuperstring(words []string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Prefix Sum, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #943: Find the Shortest Superstring
// https://leetcode.com/problems/find-the-shortest-superstring/
// Difficulty: Hard
// DP bitmask TSP: dp[mask][last] = shortest superstring for set with given last.
// Precompute overlap[i][j] = overlap of words[i] suffix with words[j] prefix.

import (
	"fmt"
	"math"
)

func shortestSuperstring(words []string) string {
	n := len(words)

	// Precompute overlap
  // Membuat matriks/slice 2D untuk DP
	overlap := make([][]int, n)
	for i := 0; i < n; i++ {
		overlap[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			maxOv := min(len(words[i]), len(words[j]))
			for k := maxOv; k >= 0; k-- {
				if words[i][len(words[i])-k:] == words[j][:k] {
					overlap[i][j] = k
					break
				}
			}
		}
	}

	// dp[mask][last] = length of shortest superstring
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, 1<<n)
  // Membuat matriks/slice 2D untuk DP
	parent := make([][]int, 1<<n)
	for mask := 0; mask < (1 << n); mask++ {
		dp[mask] = make([]int, n)
		parent[mask] = make([]int, n)
		for i := 0; i < n; i++ {
			dp[mask][i] = math.MaxInt32
			parent[mask][i] = -1
		}
	}

	// Initialize: single words
	for i := 0; i < n; i++ {
		dp[1<<i][i] = len(words[i])
	}

	// DP
	for mask := 0; mask < (1 << n); mask++ {
		for last := 0; last < n; last++ {
			if dp[mask][last] == math.MaxInt32 {
				continue
			}
			for next := 0; next < n; next++ {
				if mask&(1<<next) != 0 {
					continue
				}
				newMask := mask | (1 << next)
				newLen := dp[mask][last] + len(words[next]) - overlap[last][next]
				if newLen < dp[newMask][next] {
					dp[newMask][next] = newLen
					parent[newMask][next] = last
				}
			}
		}
	}

	// Find best last for full mask
	fullMask := (1 << n) - 1
	bestLast := 0
	bestLen := math.MaxInt32
	for i := 0; i < n; i++ {
		if dp[fullMask][i] < bestLen {
			bestLen = dp[fullMask][i]
			bestLast = i
		}
	}

	// Reconstruct
	mask := fullMask
	last := bestLast
  // Alokasi slice integer
	order := make([]int, 0, n)
	for last != -1 {
		order = append(order, last)
		prev := parent[mask][last]
		mask ^= (1 << last)
		last = prev
	}

	// Reverse order
	for i, j := 0, len(order)-1; i < j; i, j = i+1, j-1 {
		order[i], order[j] = order[j], order[i]
	}

	// Build result
	result := words[order[0]]
	for i := 1; i < n; i++ {
		result += words[order[i]][overlap[order[i-1]][order[i]]:]
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(shortestSuperstring([]string{"alex","loves","leetcode"})) // Expected: "alexlovesleetcode"
	fmt.Println(shortestSuperstring([]string{"catg","ctaagt","gcta","ttca","atgcatc"})) // Expected: "gctaagttcatgcatc"
	fmt.Println(shortestSuperstring([]string{"a","b","c"})) // Expected: "abc"
}
```
