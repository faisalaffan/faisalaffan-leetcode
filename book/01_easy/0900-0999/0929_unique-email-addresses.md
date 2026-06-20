# 0929 — Unique Email Addresses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func numUniqueEmails(emails []string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * m). Space: O(n).  |  **Ruang:** O(n).

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #929: Unique Email Addresses
// https://leetcode.com/problems/unique-email-addresses/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"})) // 2
	fmt.Println(numUniqueEmails([]string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}))                                                    // 3
}

// numUniqueEmails counts unique email addresses after applying normalization rules.
// Time: O(n * m). Space: O(n).
func numUniqueEmails(emails []string) int {
  // HashMap: O(1) lookup
	set := make(map[string]bool)
	for _, email := range emails {
		parts := strings.Split(email, "@")
		local := parts[0]
		domain := parts[1]

		// Remove dots and everything after '+'
		cleaned := strings.ReplaceAll(local, ".", "")
		if plusIdx := strings.Index(cleaned, "+"); plusIdx != -1 {
			cleaned = cleaned[:plusIdx]
		}
		set[cleaned+"@"+domain] = true
	}
	return len(set)
}
```
