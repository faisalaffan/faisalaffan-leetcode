# 0811 — Subdomain Visit Count

## Deskripsi

**Soal:** [0811. Subdomain Visit Count](https://leetcode.com/problems/subdomain-visit-count/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #811: Subdomain Visit Count
// https://leetcode.com/problems/subdomain-visit-count/
// Difficulty: Medium

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(SubdomainVisitCount([]string{"9001 discuss.leetcode.com"}))
	fmt.Println(SubdomainVisitCount([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}

func SubdomainVisitCount(cpdomains []string) []string {
  // Membuat map untuk pencarian O(1): key → value
	counts := make(map[string]int)

	for _, cpdomain := range cpdomains {
		parts := strings.SplitN(cpdomain, " ", 2)
		count, _ := strconv.Atoi(parts[0])
		domain := parts[1]

		subdomains := strings.Split(domain, ".")
  // Iterasi seluruh elemen
		for i := range subdomains {
			sub := strings.Join(subdomains[i:], ".")
			counts[sub] += count
		}
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]string, 0, len(counts))
	for sub, cnt := range counts {
		result = append(result, strconv.Itoa(cnt)+" "+sub)
	}

	return result
}
```
