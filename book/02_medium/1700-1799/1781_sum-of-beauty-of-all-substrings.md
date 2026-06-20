# 1781 — Sum Of Beauty Of All Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func beautySum(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2), Space: O(26)  |  **Ruang:** O(26)


## 💻 Solusi Go

```go
package main

// LeetCode #1781: Sum of Beauty of All Substrings
// https://leetcode.com/problems/sum-of-beauty-of-all-substrings/
// Difficulty: Medium
// Time: O(n^2), Space: O(26)

import "fmt"

func beautySum(s string) int {
	n := len(s)
	result := 0

	for i := 0; i < n; i++ {
  // Alokasi slice
		count := make([]int, 26)
		for j := i; j < n; j++ {
			count[s[j]-'a']++
			minFreq, maxFreq := n, 0
			for _, f := range count {
				if f > 0 {
					if f < minFreq {
						minFreq = f
					}
					if f > maxFreq {
						maxFreq = f
					}
				}
			}
			result += maxFreq - minFreq
		}
	}
	return result
}

func main() {
	fmt.Println(beautySum("aabcb")) // Expected: 5
	fmt.Println(beautySum("aabcbaa")) // Expected: 17
	fmt.Println(beautySum("x")) // Expected: 0
}
```
