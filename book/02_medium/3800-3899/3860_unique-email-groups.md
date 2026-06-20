# 3860 — Unique Email Groups

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func UniqueEmailGroups(emails []string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(N * M)  |  **Ruang:** O(N * M)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3860: Unique Email Groups
// https://leetcode.com/problems/unique-email-groups/
// Difficulty: Medium [Paid]
// Time: O(N * M) | Space: O(N * M)
// Approach: Normalize each email by processing local part (ignore dots,
// ignore after +) and lowercase everything. Count unique normalized forms.

import (
	"fmt"
	"strings"
)

func UniqueEmailGroups(emails []string) int {
  // HashMap: O(1) lookup
	seen := make(map[string]bool)

	for _, email := range emails {
		parts := strings.SplitN(email, "@", 2)
		local, domain := parts[0], parts[1]

		var normLocal strings.Builder
		for _, ch := range local {
			if ch == '+' {
				break
			}
			if ch != '.' {
				normLocal.WriteRune(ch)
			}
		}

		normalized := strings.ToLower(normLocal.String()) + "@" + strings.ToLower(domain)
		seen[normalized] = true
	}

	return len(seen)
}

func main() {
	// Example 1
	emails1 := []string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com",
	}
	fmt.Println(UniqueEmailGroups(emails1)) // Expected: 2

	// Example 2
	emails2 := []string{"A@B.com", "a@b.com", "ab+xy@b.com", "a.b@b.com"}
	fmt.Println(UniqueEmailGroups(emails2)) // Expected: 2

	// Example 3
	emails3 := []string{
		"a.b+c.d+e@DoMain.com",
		"ab+xyz@domain.com",
		"ab@domain.com",
	}
	fmt.Println(UniqueEmailGroups(emails3)) // Expected: 1
}
```
