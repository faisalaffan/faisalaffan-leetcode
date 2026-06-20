# 1400 — Construct K Palindrome Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah bilangan bulat (integer). Tugasmu adalah menentukan apakah bilangan tersebut adalah **palindrome** — dibaca sama dari depan maupun dari belakang.

Contoh: `121` → palindrome. `-121` → bukan (tanda minus!). `10` → bukan.

**Cara berpikir:** Balik setengah digit secara matematika menggunakan modulo (`%`) dan pembagian (`/`). Tidak perlu konversi ke string.

**Fungsi Solusi:** `func canConstruct(s string, k int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n) where n = length of string  |  **Ruang:** O(1) - fixed array of 26


## 💻 Solusi Go

```go
package main

// LeetCode #1400: Construct K Palindrome Strings
// https://leetcode.com/problems/construct-k-palindrome-strings/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(canConstruct("annabelle", 2)) // true

	// Test case 2
	fmt.Println(canConstruct("leetcode", 3)) // false

	// Test case 3
	fmt.Println(canConstruct("true", 4)) // true

	// Test case 4
	fmt.Println(canConstruct("yzyzyzyzyzyzyzy", 2)) // true
}

// Time: O(n) where n = length of string
// Space: O(1) - fixed array of 26
func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}

  // Alokasi slice
	freq := make([]int, 26)
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	// Count characters with odd frequency
	oddCount := 0
	for _, f := range freq {
		if f%2 == 1 {
			oddCount++
		}
	}

	// Each palindrome can have at most 1 odd-count character
	return oddCount <= k
}
```
