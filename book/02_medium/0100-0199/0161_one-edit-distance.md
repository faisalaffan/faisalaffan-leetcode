# 0161 — One Edit Distance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func isOneEditDistance(s string, t string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #161: One Edit Distance
// https://leetcode.com/problems/one-edit-distance/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func isOneEditDistance(s string, t string) bool {
	ns, nt := len(s), len(t)
	if abs(ns-nt) > 1 {
		return false
	}

	if ns > nt {
		s, t = t, s
		ns, nt = nt, ns
	}

	for i := 0; i < ns; i++ {
		if s[i] != t[i] {
			if ns == nt {
				return s[i+1:] == t[i+1:]
			}
			return s[i:] == t[i+1:]
		}
	}

	return ns+1 == nt
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(isOneEditDistance("ab", "acb"))
	fmt.Println(isOneEditDistance("", ""))
	fmt.Println(isOneEditDistance("a", ""))
}
```
