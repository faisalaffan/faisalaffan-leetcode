# 3280 — Convert Date To Binary

## Deskripsi

**Soal:** [3280. Convert Date To Binary](https://leetcode.com/problems/convert-date-to-binary/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3280: Convert Date to Binary
// https://leetcode.com/problems/convert-date-to-binary/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(ConvertDateToBinary("2080-02-29"))
	fmt.Println(ConvertDateToBinary("1900-01-01"))
}

// toBinary converts an integer to its binary string representation without leading zeros.
func toBinary(n int) string {
  // Edge case: input kosong
	if n == 0 {
		return "0"
	}
	s := ""
	for n > 0 {
		s = string('0'+byte(n%2)) + s
		n /= 2
	}
	return s
}

// ConvertDateToBinary converts a date string to binary format.
// Time: O(1). Space: O(1).
func ConvertDateToBinary(date string) string {
	parts := strings.Split(date, "-")
	year, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	day, _ := strconv.Atoi(parts[2])
	return toBinary(year) + "-" + toBinary(month) + "-" + toBinary(day)
}
```
