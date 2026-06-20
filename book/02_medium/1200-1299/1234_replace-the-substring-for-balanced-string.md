# 1234 — Replace The Substring For Balanced String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func balancedString(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer, Sliding Window

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1234: Replace the Substring for Balanced String
// https://leetcode.com/problems/replace-the-substring-for-balanced-string/
// Difficulty: Medium

// Sliding window. Find smallest substring such that outside it,
// each of Q,W,E,R appears at most n/4 times.

// Time: O(n)
// Space: O(1)

func balancedString(s string) int {
	n := len(s)
	target := n / 4
  // HashMap: O(1) lookup
	count := make(map[byte]int)
	for i := 0; i < n; i++ {
		count[s[i]]++
	}

	// Check if already balanced
	balanced := true
	for _, c := range []byte{'Q', 'W', 'E', 'R'} {
		if count[c] > target {
			balanced = false
			break
		}
	}
	if balanced {
		return 0
	}

	left := 0
	minLen := n

	for right := 0; right < n; right++ {
		count[s[right]]--

  // Binary search loop
		for left <= right {
			ok := true
			for _, c := range []byte{'Q', 'W', 'E', 'R'} {
				if count[c] > target {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			count[s[left]]++
			left++
		}
	}

	return minLen
}

func main() {
	fmt.Printf("%d (expected: 0)\n", balancedString("QWER"))
	fmt.Printf("%d (expected: 1)\n", balancedString("QQWE"))
	fmt.Printf("%d (expected: 2)\n", balancedString("QQQW"))
}
```
