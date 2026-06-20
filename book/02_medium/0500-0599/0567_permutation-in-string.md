# 0567 — Permutation In String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CheckInclusion(s1 string, s2 string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n + m) where n = len(s1), m = len(s2)  |  **Ruang:** O(1) (fixed 26 chars)


## 💻 Solusi Go

```go
package main

// LeetCode #567: Permutation in String
// https://leetcode.com/problems/permutation-in-string/
// Difficulty: Medium
// Time: O(n + m) where n = len(s1), m = len(s2)
// Space: O(1) (fixed 26 chars)

import "fmt"

func main() {
	fmt.Println(CheckInclusion("ab", "eidbaooo"))
	fmt.Println(CheckInclusion("ab", "eidboaoo"))
}

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	count1 := [26]int{}
	count2 := [26]int{}

  // Linear scan O(n)
	for i := 0; i < len(s1); i++ {
		count1[s1[i]-'a']++
		count2[s2[i]-'a']++
	}

	if count1 == count2 {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		count2[s2[i]-'a']++
		count2[s2[i-len(s1)]-'a']--
		if count1 == count2 {
			return true
		}
	}

	return false
}
```
