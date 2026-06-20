# 0811 — Subdomain Visit Count

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func SubdomainVisitCount(cpdomains []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

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
  // Membuat map (HashMap) — pencarian O(1)
	counts := make(map[string]int)

	for _, cpdomain := range cpdomains {
		parts := strings.SplitN(cpdomain, " ", 2)
		count, _ := strconv.Atoi(parts[0])
		domain := parts[1]

		subdomains := strings.Split(domain, ".")
  // Range loop: iterasi dengan indeks + nilai
		for i := range subdomains {
			sub := strings.Join(subdomains[i:], ".")
			counts[sub] += count
		}
	}

	result := make([]string, 0, len(counts))
	for sub, cnt := range counts {
		result = append(result, strconv.Itoa(cnt)+" "+sub)
	}

	return result
}
```
