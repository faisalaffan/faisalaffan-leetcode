# 3675 — Minimum Operations To Transform String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minimumOperationsToTransformString(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3675: Minimum Operations to Transform String
// https://leetcode.com/problems/minimum-operations-to-transform-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumOperationsToTransformString(s string) int {
	minChar := byte('z' + 1)
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != 'a' && c < minChar {
			minChar = c
			if minChar == 'b' {
				break
			}
		}
	}
	if minChar > 'z' {
		return 0
	}
	return int('z' + 1 - minChar)
}

func main() {
	fmt.Println(minimumOperationsToTransformString("yz"))
	fmt.Println(minimumOperationsToTransformString("a"))
	fmt.Println(minimumOperationsToTransformString("abc"))
}
```
