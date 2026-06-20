# 2083 — Substrings That Begin And End With The Same Letter

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func numberOfSubstrings(s string) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2083: Substrings That Begin and End With the Same Letter
// https://leetcode.com/problems/substrings-that-begin-and-end-with-the-same-letter/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func numberOfSubstrings(s string) int64 {
  // Alokasi slice
	freq := make([]int64, 26)
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}
	var result int64 = 0
	for _, f := range freq {
		result += f * (f + 1) / 2
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfSubstrings("abc"))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", numberOfSubstrings("abacaba"))
	// Expected: 12

	// Test case 3
	fmt.Println("Test 3:", numberOfSubstrings("aa"))
	// Expected: 3
}
```
