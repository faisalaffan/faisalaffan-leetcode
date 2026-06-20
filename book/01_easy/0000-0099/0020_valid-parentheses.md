# 0020 — Valid Parentheses

## Deskripsi

**Soal:** [0020. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func IsValid(s string) bool`

## Solusi Go

```go
package main

// LeetCode #20: Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(n)
func IsValid(s string) bool {
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
  // Membuat slice untuk menyimpan hasil
	stack := make([]byte, 0, len(s))
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(IsValid("()"))
	fmt.Println(IsValid("()[]{}"))
	fmt.Println(IsValid("(]"))
}
```
