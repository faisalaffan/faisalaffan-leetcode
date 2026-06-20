# 1663 — Smallest String With A Given Numeric Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func GetSmallestString(n int, k int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N), Space: O(N)  |  **Ruang:** O(N)


## 💻 Solusi Go

```go
package main

// LeetCode #1663: Smallest String With A Given Numeric Value
// https://leetcode.com/problems/smallest-string-with-a-given-numeric-value/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GetSmallestString(3, 27))
	fmt.Println(GetSmallestString(5, 73))
	fmt.Println(GetSmallestString(1, 26))
}

func GetSmallestString(n int, k int) string {
	// Time: O(N), Space: O(N)
	// Greedy: fill from the end with 'z' as much as possible
	result := make([]byte, n)
  // Range loop
	for i := range result {
		result[i] = 'a'
	}

	k -= n // All positions have at least 'a' (value 1)

	for i := n - 1; i >= 0 && k > 0; i-- {
		add := 25 // 'z' - 'a' = 25
		if k < add {
			add = k
		}
		result[i] = byte('a' + add)
		k -= add
	}

	return string(result)
}
```
