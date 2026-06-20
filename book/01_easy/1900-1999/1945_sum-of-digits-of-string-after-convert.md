# 1945 — Sum Of Digits Of String After Convert

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func SumOfDigitsOfStringAfterConvert(s string, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1945: Sum of Digits of String After Convert
// https://leetcode.com/problems/sum-of-digits-of-string-after-convert/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(SumOfDigitsOfStringAfterConvert("iiii", 1))  // 36
	fmt.Println(SumOfDigitsOfStringAfterConvert("leetcode", 2))  // 6
}

// Time: O(n), Space: O(n)
func SumOfDigitsOfStringAfterConvert(s string, k int) int {
	var digits string
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		digits += strconv.Itoa(int(s[i] - 'a' + 1))
	}

	for t := 0; t < k; t++ {
		sum := 0
  // Linear scan O(n)
		for i := 0; i < len(digits); i++ {
			sum += int(digits[i] - '0')
		}
		digits = strconv.Itoa(sum)
	}

	result, _ := strconv.Atoi(digits)
	return result
}
```
