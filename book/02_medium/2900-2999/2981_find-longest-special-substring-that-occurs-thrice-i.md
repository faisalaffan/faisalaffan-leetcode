# 2981 — Find Longest Special Substring That Occurs Thrice I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func maximumLength2981(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2981: Find Longest Special Substring That Occurs Thrice I
// https://leetcode.com/problems/find-longest-special-substring-that-occurs-thrice-i/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(maximumLength2981("aaaa"))
	fmt.Println(maximumLength2981("abcdef"))
	fmt.Println(maximumLength2981("abcaba"))
}

func maximumLength2981(s string) int {
	n := len(s)
	l, r := 0, n
	check := func(x int) bool {
		cnt := [26]int{}
		for i := 0; i < n; {
			j := i + 1
			for j < n && s[j] == s[i] {
				j++
			}
			k := s[i] - 'a'
			add := j - i - x + 1
			if add > 0 {
				cnt[k] += add
			}
			if cnt[k] >= 3 {
				return true
			}
			i = j
		}
		return false
	}
	for l < r {
		mid := (l + r + 1) >> 1
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if l == 0 {
		return -1
	}
	return l
}
```
