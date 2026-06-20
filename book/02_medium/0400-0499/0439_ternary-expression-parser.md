# 0439 — Ternary Expression Parser

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func parseTernary(expression string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Stack

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #439: Ternary Expression Parser
// https://leetcode.com/problems/ternary-expression-parser/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func parseTernary(expression string) string {
	stack := make([]byte, 0)
	// Process from right to left
	for i := len(expression) - 1; i >= 0; i-- {
		ch := expression[i]
		if ch >= '0' && ch <= '9' || ch == 'T' || ch == 'F' {
			stack = append(stack, ch)
		} else if ch == '?' {
			// Evaluate: preceding char is condition
			trueVal := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			falseVal := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			cond := expression[i-1]
			i-- // skip the condition character (already consumed)
			if cond == 'T' {
				stack = append(stack, trueVal)
			} else {
				stack = append(stack, falseVal)
			}
		}
		// Skip ':'
	}
	return string(stack[0])
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", parseTernary("T?2:3"))
	// Expected: "2"

	// Test case 2
	fmt.Println("Test 2:", parseTernary("F?1:T?4:5"))
	// Expected: "4"

	// Test case 3
	fmt.Println("Test 3:", parseTernary("T?T?F:5:3"))
	// Expected: "F"
}
```
