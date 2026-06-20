# 3186 — Maximum Total Damage With Spell Casting

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumTotalDamage(power []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Dynamic Programming

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3186: Maximum Total Damage With Spell Casting
// https://leetcode.com/problems/maximum-total-damage-with-spell-casting/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maximumTotalDamage(power []int) int64 {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, v := range power {
		freq[v]++
	}

  // Alokasi slice integer
	vals := make([]int, 0, len(freq))
	for k := range freq {
		vals = append(vals, k)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(vals)

	n := len(vals)
  // Alokasi slice integer
	dp := make([]int64, n)

	for i := 0; i < n; i++ {
		val := vals[i]
		count := freq[val]
		dp[i] = int64(val) * int64(count)

		// Find prev valid (val - 2)
		for j := i - 1; j >= 0; j-- {
			if vals[j] < val-2 {
				if dp[j] > dp[i] {
					dp[i] = dp[j]
				}
				break
			}
			if vals[j] <= val-2 {
				dp[i] = maxInt64(dp[i], dp[j]+int64(val)*int64(count))
			}
		}

		if i > 0 && dp[i-1] > dp[i] {
			dp[i] = dp[i-1]
		}
	}

	return dp[n-1]
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumTotalDamage([]int{1, 1, 3, 4}))          // Expected: 6
	fmt.Println(maximumTotalDamage([]int{7, 1, 6, 3}))           // Expected: 10
}
```
