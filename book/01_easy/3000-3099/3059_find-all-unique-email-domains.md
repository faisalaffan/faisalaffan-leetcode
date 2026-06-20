# 3059 — Find All Unique Email Domains

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func FindAllUniqueEmailDomains(emails []string) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

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
  // HashMap: O(1) lookup
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
