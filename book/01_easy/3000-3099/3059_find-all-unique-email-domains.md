# 3059 — Find All Unique Email Domains

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindAllUniqueEmailDomains(emails []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3059: Find All Unique Email Domains
// https://leetcode.com/problems/find-all-unique-email-domains/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent logic: extract and count unique email domains.

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// LeetCode name: findUniqueDomains
	emails := []string{"alice@leetcode.com", "bob@leetcode.com", "charlie@gmail.com"}
	fmt.Println(FindAllUniqueEmailDomains(emails))
	// [gmail.com leetcode.com]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: findUniqueDomains
func FindAllUniqueEmailDomains(emails []string) []string {
  // Membuat map (HashMap) — pencarian O(1)
	domains := make(map[string]bool)
	for _, email := range emails {
		parts := strings.Split(email, "@")
		if len(parts) == 2 {
			domains[parts[1]] = true
		}
	}
	result := make([]string, 0, len(domains))
	for d := range domains {
		result = append(result, d)
	}
	sort.Strings(result)
	return result
}
```
