# 0032 — Longest Valid Parentheses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func longestValidParentheses(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #32: Longest Valid Parentheses
// https://leetcode.com/problems/longest-valid-parentheses/
// Difficulty: Hard

import "fmt"

// longestValidParentheses finds the length of the longest valid parentheses substring.
// Uses a stack-based approach.
//
// Complexity: O(n) time, O(n) space
func longestValidParentheses(s string) int {
	stack := []int{-1} // base index for valid substring calculation
	maxLen := 0

	for i, ch := range s {
		if ch == '(' {
			stack = append(stack, i)
		} else {
			// Pop the matching '('
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				// No matching '('; set new base index
				stack = append(stack, i)
			} else {
				// Calculate length of current valid substring
				length := i - stack[len(stack)-1]
				if length > maxLen {
					maxLen = length
				}
			}
		}
	}

	return maxLen
}

func main() {
	// Test cases from LeetCode
	fmt.Println("Test 1: (() ->", longestValidParentheses("(()"))       // 2
	fmt.Println("Test 2: )()()) ->", longestValidParentheses(")()())"))  // 4
	fmt.Println("Test 3: '' ->", longestValidParentheses(""))           // 0
	fmt.Println("Test 4: ()() ->", longestValidParentheses("()()"))     // 4
	fmt.Println("Test 5: (()()) ->", longestValidParentheses("(()())")) // 6
}
```
