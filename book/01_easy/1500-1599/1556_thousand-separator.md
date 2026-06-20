# 1556 — Thousand Separator

## Deskripsi

**Soal:** [1556. Thousand Separator](https://leetcode.com/problems/thousand-separator/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n), Space: O(log n)  
**Kompleksitas Ruang:** O(log n)

**Algoritma:** —

**Fungsi Solusi:** `func thousandSeparator(n int) string`

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
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
