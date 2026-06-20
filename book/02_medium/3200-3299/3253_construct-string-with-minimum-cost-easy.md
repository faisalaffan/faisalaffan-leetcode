# 3253 — Construct String With Minimum Cost Easy

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumCost(target string, words []string, costs []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n * m * L)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3253: Construct String with Minimum Cost (Easy)
// https://leetcode.com/problems/construct-string-with-minimum-cost-easy/
// Difficulty: Medium [Paid]
// Time: O(n * m * L) | Space: O(n)

import (
	"fmt"
	"math"
)

func minimumCost(target string, words []string, costs []int) int {
	n := len(target)
  // Alokasi slice integer
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 0; i < n; i++ {
		if dp[i] == math.MaxInt32 {
			continue
		}
		for j, w := range words {
			if i+len(w) <= n && target[i:i+len(w)] == w {
				if dp[i]+costs[j] < dp[i+len(w)] {
					dp[i+len(w)] = dp[i] + costs[j]
				}
			}
		}
	}

	if dp[n] == math.MaxInt32 {
		return -1
	}
	return dp[n]
}

func main() {
	fmt.Println(minimumCost("abc", []string{"a", "bc", "abc"}, []int{1, 2, 3})) // Expected: 3
	fmt.Println(minimumCost("xyz", []string{"ab", "cd"}, []int{1, 2}))          // Expected: -1
}
```
