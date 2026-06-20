# 2840 — Check If Strings Can Be Made Equal With Operations Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CheckIfStringsCanBeMadeEqualWithOperationsIi(s1 string, s2 string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2840: Check if Strings Can be Made Equal With Operations II
// https://leetcode.com/problems/check-if-strings-can-be-made-equal-with-operations-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func CheckIfStringsCanBeMadeEqualWithOperationsIi(s1 string, s2 string) bool {
	n := len(s1)
	// We can swap characters at positions with same parity
	// Count character frequencies at even and odd positions
  // Alokasi slice
	even := make([]int, 26)
  // Alokasi slice
	odd := make([]int, 26)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			even[s1[i]-'a']++
			even[s2[i]-'a']--
		} else {
			odd[s1[i]-'a']++
			odd[s2[i]-'a']--
		}
	}

	for i := 0; i < 26; i++ {
		if even[i] != 0 || odd[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(CheckIfStringsCanBeMadeEqualWithOperationsIi("abcd", "cdab"))
	fmt.Println(CheckIfStringsCanBeMadeEqualWithOperationsIi("abcd", "dcba"))
}
```
