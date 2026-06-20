# 1903 — Largest Odd Number In String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func LargestOddNumberInString(num string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1903: Largest Odd Number in String
// https://leetcode.com/problems/largest-odd-number-in-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(LargestOddNumberInString("52"))      // "5"
	fmt.Println(LargestOddNumberInString("4206"))     // ""
	fmt.Println(LargestOddNumberInString("35427"))    // "35427"
}

// Time: O(n), Space: O(1)
func LargestOddNumberInString(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if (num[i]-'0')%2 == 1 {
			return num[:i+1]
		}
	}
	return ""
}
```
