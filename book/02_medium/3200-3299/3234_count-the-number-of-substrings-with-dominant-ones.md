# 3234 — Count The Number Of Substrings With Dominant Ones

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func numberOfSubstrings(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * sqrt(n))  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3234: Count the Number of Substrings With Dominant Ones
// https://leetcode.com/problems/count-the-number-of-substrings-with-dominant-ones/
// Difficulty: Medium
// Time: O(n * sqrt(n)) | Space: O(1)

import (
	"fmt"
	"math"
)

func numberOfSubstrings(s string) int {
	n := len(s)
	ans := 0
	maxZeros := int(math.Sqrt(float64(n)))

	for l := 0; l < n; l++ {
		zeros := 0
		ones := 0
		for r := l; r < n; r++ {
			if s[r] == '0' {
				zeros++
				if zeros > maxZeros {
					break
				}
			} else {
				ones++
			}
			if ones >= zeros*zeros {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfSubstrings("00011")) // Expected: 5
	fmt.Println(numberOfSubstrings("101"))    // Expected: 3
}
```
