# 3860 — Unique Email Groups

## Deskripsi

**Soal:** [3860. Unique Email Groups](https://leetcode.com/problems/unique-email-groups/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N * M)  
**Kompleksitas Ruang:** O(N * M)

**Algoritma:** —

**Fungsi Solusi:** `func UniqueEmailGroups(emails []string) int`

> **Ide Kunci:** Normalize each email by processing local part (ignore dots,

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
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
