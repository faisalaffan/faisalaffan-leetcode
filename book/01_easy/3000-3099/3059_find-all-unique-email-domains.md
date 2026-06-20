# 3059 — Find All Unique Email Domains

## Deskripsi

**Soal:** [3059. Find All Unique Email Domains](https://leetcode.com/problems/find-all-unique-email-domains/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
	domains := make(map[string]bool)
	for _, email := range emails {
		parts := strings.Split(email, "@")
		if len(parts) == 2 {
			domains[parts[1]] = true
		}
	}
  // Membuat slice untuk menyimpan hasil
	result := make([]string, 0, len(domains))
	for d := range domains {
		result = append(result, d)
	}
	sort.Strings(result)
	return result
}
```
