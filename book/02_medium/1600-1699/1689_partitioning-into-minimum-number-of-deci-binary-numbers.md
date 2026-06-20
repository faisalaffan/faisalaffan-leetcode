# 1689 — Partitioning Into Minimum Number Of Deci Binary Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minPartitions(n string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1689: Partitioning Into Minimum Number Of Deci-Binary Numbers
// https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func minPartitions(n string) int {
	maxDigit := 0
	for _, ch := range n {
		digit := int(ch - '0')
		if digit > maxDigit {
			maxDigit = digit
		}
		if maxDigit == 9 {
			break // can't get higher than 9
		}
	}
	return maxDigit
}

func main() {
	fmt.Println(minPartitions("32"))     // Expected: 3
	fmt.Println(minPartitions("82734"))  // Expected: 8
	fmt.Println(minPartitions("27346209830709182346")) // Expected: 9
}
```
