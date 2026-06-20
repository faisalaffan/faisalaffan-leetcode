# 2825 — Make String A Subsequence Using Cyclic Increments

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MakeStringASubsequenceUsingCyclicIncrements(str1 string, str2 string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2825: Make String a Subsequence Using Cyclic Increments
// https://leetcode.com/problems/make-string-a-subsequence-using-cyclic-increments/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func MakeStringASubsequenceUsingCyclicIncrements(str1 string, str2 string) bool {
	j := 0
  // Linear scan O(n)
	for i := 0; i < len(str1) && j < len(str2); i++ {
		if str1[i] == str2[j] || (str1[i]-'a'+1)%26 == str2[j]-'a' {
			j++
		}
	}
	return j == len(str2)
}

func main() {
	fmt.Println(MakeStringASubsequenceUsingCyclicIncrements("abc", "bcd"))
	fmt.Println(MakeStringASubsequenceUsingCyclicIncrements("abc", "ad"))
}
```
