# 0929 — Unique Email Addresses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func numUniqueEmails(emails []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Membuat map (HashMap) — pencarian O(1)
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
