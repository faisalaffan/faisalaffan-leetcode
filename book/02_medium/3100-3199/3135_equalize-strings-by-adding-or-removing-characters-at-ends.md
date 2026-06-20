# 3135 — Equalize Strings By Adding Or Removing Characters At Ends

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minOperations(initial string, target string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * m)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3135: Equalize Strings by Adding or Removing Characters at Ends
// https://leetcode.com/problems/equalize-strings-by-adding-or-removing-characters-at-ends/
// Difficulty: Medium [Paid]
// Time: O(n * m) | Space: O(1)

import "fmt"

func minOperations(initial string, target string) int {
	m, n := len(initial), len(target)
	longest := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k := 0
			for i+k < m && j+k < n && initial[i+k] == target[j+k] {
				k++
			}
			if k > longest {
				longest = k
			}
		}
	}

	return m + n - 2*longest
}

func main() {
	fmt.Println(minOperations("abcdef", "defabc")) // Expected: 0 (lcs "def" or "abc")
	fmt.Println(minOperations("abc", "xyz"))       // Expected: 6
	fmt.Println(minOperations("abcde", "cde"))     // Expected: 2
}
```
