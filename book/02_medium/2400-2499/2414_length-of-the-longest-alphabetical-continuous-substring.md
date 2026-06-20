# 2414 — Length Of The Longest Alphabetical Continuous Substring

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func longestContinuousSubstring(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2414: Length of the Longest Alphabetical Continuous Substring
// https://leetcode.com/problems/length-of-the-longest-alphabetical-continuous-substring/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Scan, count consecutive chars where s[i] == s[i-1] + 1.

import "fmt"

func main() {
	fmt.Println(longestContinuousSubstring("abacaba")) // 2 ("ab")
	fmt.Println(longestContinuousSubstring("abcde"))   // 5
}

func longestContinuousSubstring(s string) int {
	ans, cur := 0, 0
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if i == 0 || s[i] == s[i-1]+1 {
			cur++
		} else {
			cur = 1
		}
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
```
