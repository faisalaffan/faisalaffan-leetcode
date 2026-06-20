# 1759 — Count Number Of Homogenous Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func countHomogenous(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1759: Count Number of Homogenous Substrings
// https://leetcode.com/problems/count-number-of-homogenous-substrings/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

const mod = 1_000_000_007

func countHomogenous(s string) int {
	result := 0
	count := 0

  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}
		result = (result + count) % mod
	}
	return result
}

func main() {
	fmt.Println(countHomogenous("abbcccaa")) // Expected: 13
	fmt.Println(countHomogenous("xy"))        // Expected: 2
	fmt.Println(countHomogenous("zzzzz"))     // Expected: 15
}
```
