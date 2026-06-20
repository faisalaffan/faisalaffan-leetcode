# 3491 — Phone Number Prefix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func PhoneNumberPrefix(numbers []string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting, Prefix Sum

**Waktu:** O(n log n * m). Space: O(1).  |  **Ruang:** O(1).

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3491: Phone Number Prefix
// https://leetcode.com/problems/phone-number-prefix/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(PhoneNumberPrefix([]string{"123", "1234", "567", "7890"}))
	fmt.Println(PhoneNumberPrefix([]string{"111", "222", "333"}))
}

// PhoneNumberPrefix returns true if no number is a prefix of another number.
// Time: O(n log n * m). Space: O(1).
func PhoneNumberPrefix(numbers []string) bool {
	sort.Strings(numbers)
  // Linear scan O(n)
	for i := 0; i < len(numbers)-1; i++ {
		if len(numbers[i]) <= len(numbers[i+1]) {
			isPrefix := true
			for j := 0; j < len(numbers[i]); j++ {
				if numbers[i][j] != numbers[i+1][j] {
					isPrefix = false
					break
				}
			}
			if isPrefix {
				return false
			}
		}
	}
	return true
}
```
