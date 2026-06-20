# 3941 — Password Strength

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func PasswordStrength(password string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(N)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3941: Password Strength
// https://leetcode.com/problems/password-strength/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Track distinct characters per category. Sum points:
// lowercase=1, uppercase=2, digit=3, special=5.

import "fmt"

func PasswordStrength(password string) int {
  // HashMap: O(1) lookup
	lower := make(map[byte]bool)
  // HashMap: O(1) lookup
	upper := make(map[byte]bool)
  // HashMap: O(1) lookup
	digit := make(map[byte]bool)
  // HashMap: O(1) lookup
	special := make(map[byte]bool)

  // Linear scan O(n)
	for i := 0; i < len(password); i++ {
		ch := password[i]
		if ch >= 'a' && ch <= 'z' {
			lower[ch] = true
		} else if ch >= 'A' && ch <= 'Z' {
			upper[ch] = true
		} else if ch >= '0' && ch <= '9' {
			digit[ch] = true
		} else {
			// Special: ! @ # $
			special[ch] = true
		}
	}

	return len(lower)*1 + len(upper)*2 + len(digit)*3 + len(special)*5
}

func main() {
	// Example 1
	fmt.Println(PasswordStrength("aA1!")) // Expected: 11

	// Example 2
	fmt.Println(PasswordStrength("bbB11#")) // Expected: 11
}
```
