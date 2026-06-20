# 2839 — Check If Strings Can Be Made Equal With Operations I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CheckIfStringsCanBeMadeEqualWithOperationsI(s1 string, s2 string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2839: Check if Strings Can be Made Equal With Operations I
// https://leetcode.com/problems/check-if-strings-can-be-made-equal-with-operations-i/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(CheckIfStringsCanBeMadeEqualWithOperationsI("abcd", "cdab"))
	fmt.Println(CheckIfStringsCanBeMadeEqualWithOperationsI("abcd", "dacb"))
}

func CheckIfStringsCanBeMadeEqualWithOperationsI(s1 string, s2 string) bool {
	// Can swap characters at even indices (0<->2) and odd indices (1<->3)
	// Check that multiset of chars at same parity positions match
	return (s1[0] == s2[0] || s1[0] == s2[2]) &&
		(s1[2] == s2[2] || s1[2] == s2[0]) &&
		(s1[1] == s2[1] || s1[1] == s2[3]) &&
		(s1[3] == s2[3] || s1[3] == s2[1])
}
```
