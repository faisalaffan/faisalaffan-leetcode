# 0306 — Additive Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func isAdditiveNumber(num string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n^2), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #306: Additive Number
// https://leetcode.com/problems/additive-number/
// Difficulty: Medium
// Time: O(n^2), Space: O(n)

import (
	"fmt"
	"strconv"
	"strings"
)

func isAdditiveNumber(num string) bool {
	n := len(num)

	for firstEnd := 1; firstEnd <= n/2; firstEnd++ {
		if num[0] == '0' && firstEnd > 1 {
			break
		}
		num1, _ := strconv.ParseInt(num[:firstEnd], 10, 64)

		for secondEnd := firstEnd + 1; max(firstEnd, secondEnd-firstEnd) <= n-secondEnd; secondEnd++ {
			if num[firstEnd] == '0' && secondEnd-firstEnd > 1 {
				break
			}
			num2, _ := strconv.ParseInt(num[firstEnd:secondEnd], 10, 64)

			if isValid(num, num1, num2, secondEnd) {
				return true
			}
		}
	}

	return false
}

func isValid(num string, num1, num2 int64, start int) bool {
	if start == len(num) {
		return false
	}

	for start < len(num) {
		sum := num1 + num2
		sumStr := strconv.FormatInt(sum, 10)

		if !strings.HasPrefix(num[start:], sumStr) {
			return false
		}

		start += len(sumStr)
		num1, num2 = num2, sum
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(isAdditiveNumber("112358"))
	fmt.Println(isAdditiveNumber("199100199"))
	fmt.Println(isAdditiveNumber("12"))
}
```
