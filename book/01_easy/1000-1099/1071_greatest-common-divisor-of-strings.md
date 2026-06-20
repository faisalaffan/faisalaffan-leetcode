# 1071 — Greatest Common Divisor Of Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func gcdOfStrings(str1, str2 string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n + m)  |  **Ruang:** O(n + m)


## 💻 Solusi Go

```go
package main

// LeetCode #1071: Greatest Common Divisor of Strings
// https://leetcode.com/problems/greatest-common-divisor-of-strings/
// Difficulty: Easy
// Time: O(n + m) | Space: O(n + m)

import "fmt"

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC")) // "ABC"
	fmt.Println(gcdOfStrings("ABABAB", "ABAB")) // "AB"
	fmt.Println(gcdOfStrings("LEET", "CODE"))   // ""
}

// LeetCode submission: gcdOfStrings
func gcdOfStrings(str1, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[:gcd(len(str1), len(str2))]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```
