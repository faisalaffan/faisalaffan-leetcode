# 1758 — Minimum Changes To Make Alternating Binary String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MinOperations(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1758: Minimum Changes to Make Alternating Binary String
// https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func MinOperations(s string) int {
	changes := 0
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		expected := byte('0' + i%2)
		if s[i] != expected {
			changes++
		}
	}
	if changes < len(s)-changes {
		return changes
	}
	return len(s) - changes
}

func main() {
	fmt.Println(MinOperations("0100"))
	fmt.Println(MinOperations("10"))
	fmt.Println(MinOperations("1111"))
}
```
