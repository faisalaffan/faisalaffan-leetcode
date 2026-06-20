# 1556 — Thousand Separator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func thousandSeparator(n int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n), Space: O(log n)  |  **Ruang:** O(log n)


## 💻 Solusi Go

```go
package main

// LeetCode #1556: Thousand Separator
// https://leetcode.com/problems/thousand-separator/
// Difficulty: Easy
//
// LeetCode submission: func thousandSeparator(n int) string

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ThousandSeparator(987))      // "987"
	fmt.Println(ThousandSeparator(1234))     // "1.234"
	fmt.Println(ThousandSeparator(1000000))  // "1.000.000"
}

// Time: O(log n), Space: O(log n)
func ThousandSeparator(n int) string {
	s := strconv.Itoa(n)
	res := make([]byte, 0, len(s)+len(s)/3)
	for i, ch := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			res = append(res, '.')
		}
		res = append(res, byte(ch))
	}
	return string(res)
}
```
