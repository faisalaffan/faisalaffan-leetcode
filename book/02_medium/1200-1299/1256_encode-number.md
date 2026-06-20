# 1256 — Encode Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func encode(num int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(log n)


## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1256: Encode Number
// https://leetcode.com/problems/encode-number/
// Difficulty: Medium [Paid]

// Encode n as binary of n+1, then remove first bit.
// n=0 -> "0" (binary of 1 -> "1", remove first -> "")
// Wait: n=0 -> "". Let me check the pattern.
// 0: "" (1->"1", drop first->"")
// 1: "0" (2->"10", drop first->"0")
// 2: "1" (3->"11", drop first->"1")
// 3: "00" (4->"100", drop first->"00")
// 4: "01" (5->"101", drop first->"01")

// Time: O(log n)
// Space: O(log n)

func encode(num int) string {
	if num == 0 {
		return ""
	}

	// num+1 in binary, then drop first bit
	n := num + 1
	result := ""

	for n > 1 {
		if n%2 == 0 {
			result = "0" + result
		} else {
			result = "1" + result
		}
		n /= 2
	}

	return result
}

func main() {
	fmt.Printf("%q (expected: %q)\n", encode(0), "")
	fmt.Printf("%q (expected: %q)\n", encode(1), "0")
	fmt.Printf("%q (expected: %q)\n", encode(2), "1")
	fmt.Printf("%q (expected: %q)\n", encode(3), "00")
}
```
