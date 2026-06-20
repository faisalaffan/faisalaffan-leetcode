# 3723 — Maximize Sum Of Squares Of Digits

## Deskripsi

**Soal:** [3723. Maximize Sum Of Squares Of Digits](https://leetcode.com/problems/maximize-sum-of-squares-of-digits/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func maximizeSumOfSquaresOfDigits(num int, total int) string`

## Solusi Go

```go
package main

// LeetCode #3723: Maximize Sum of Squares of Digits
// https://leetcode.com/problems/maximize-sum-of-squares-of-digits/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func maximizeSumOfSquaresOfDigits(num int, total int) string {
	if num*9 < total {
		return ""
	}
	nines := total / 9
	rem := total % 9
	var sb strings.Builder
	sb.WriteString(strings.Repeat("9", nines))
	if rem > 0 {
		sb.WriteByte(byte(rem) + '0')
	}
	for sb.Len() < num {
		sb.WriteByte('0')
	}
	return sb.String()
}

func main() {
	fmt.Println(maximizeSumOfSquaresOfDigits(2, 3))
	fmt.Println(maximizeSumOfSquaresOfDigits(2, 17))
	fmt.Println(maximizeSumOfSquaresOfDigits(1, 10))
}
```
