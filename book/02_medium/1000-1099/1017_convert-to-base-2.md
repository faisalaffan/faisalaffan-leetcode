# 1017 — Convert To Base 2

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func baseNeg2(n int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(log n)


## 💻 Solusi Go

```go
package main

// LeetCode #1017: Convert to Base -2
// https://leetcode.com/problems/convert-to-base-2/
// Difficulty: Medium
//
// Approach: Repeated division by -2. Handle negative remainder.
// Time: O(log n)
// Space: O(log n)

import "fmt"

func main() {
	fmt.Println(baseNeg2(2))  // "110"
	fmt.Println(baseNeg2(3))  // "111"
	fmt.Println(baseNeg2(4))  // "100"
}

func baseNeg2(n int) string {
  // Edge case: input kosong
	if n == 0 {
		return "0"
	}

	result := ""
	for n != 0 {
		remainder := n % -2
		n /= -2
		if remainder < 0 {
			remainder += 2
			n++
		}
		result = string(rune('0'+remainder)) + result
	}

	return result
}
```
