# 0227 — Basic Calculator Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func calculate(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #227: Basic Calculator II
// https://leetcode.com/problems/basic-calculator-ii/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func calculate(s string) int {
	stack := []int{}
	num := 0
	op := '+'

	for i, ch := range s {
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
		}

		if ch == '+' || ch == '-' || ch == '*' || ch == '/' || i == len(s)-1 {
			if s[i] == ' ' && i != len(s)-1 {
				continue
			}
			switch op {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			op = ch
			num = 0
		}
	}

	result := 0
	for _, v := range stack {
		result += v
	}
	return result
}

func main() {
	fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate(" 3/2 "))
	fmt.Println(calculate(" 3+5 / 2 "))
}
```
