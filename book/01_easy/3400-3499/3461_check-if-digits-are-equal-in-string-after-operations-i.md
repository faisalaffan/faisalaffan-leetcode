# 3461 — Check If Digits Are Equal In String After Operations I

## Deskripsi

**Soal:** [3461. Check If Digits Are Equal In String After Operations I](https://leetcode.com/problems/check-if-digits-are-equal-in-string-after-operations-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^2). Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3461: Check If Digits Are Equal in String After Operations I
// https://leetcode.com/problems/check-if-digits-are-equal-in-string-after-operations-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfDigitsAreEqualInStringAfterOperationsI("1234"))
	fmt.Println(CheckIfDigitsAreEqualInStringAfterOperationsI("1111"))
}

// CheckIfDigitsAreEqualInStringAfterOperationsI repeatedly replaces adjacent digit pairs with (sum % 10) until 2 digits remain, then checks equality.
// Time: O(n^2). Space: O(n).
func CheckIfDigitsAreEqualInStringAfterOperationsI(s string) bool {
  // Membuat slice untuk menyimpan hasil
	digits := make([]int, len(s))
	for i, ch := range s {
		digits[i] = int(ch - '0')
	}

	for len(digits) > 2 {
  // Membuat slice untuk menyimpan hasil
		next := make([]int, len(digits)-1)
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(digits)-1; i++ {
			next[i] = (digits[i] + digits[i+1]) % 10
		}
		digits = next
	}
	return digits[0] == digits[1]
}
```
