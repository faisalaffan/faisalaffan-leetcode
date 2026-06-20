# 3174 — Clear Digits

## Deskripsi

**Soal:** [3174. Clear Digits](https://leetcode.com/problems/clear-digits/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** Stack (tumpukan LIFO)

## Solusi Go

```go
package main

// LeetCode #3174: Clear Digits
// https://leetcode.com/problems/clear-digits/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ClearDigits("abc"))
	fmt.Println(ClearDigits("cb34"))
	fmt.Println(ClearDigits("a1b2c3"))
}

// ClearDigits removes all digits and their nearest non-digit character to the left.
// Time: O(n). Space: O(n).
func ClearDigits(s string) string {
  // Membuat slice untuk menyimpan hasil
	stack := make([]byte, 0, len(s))
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
```
