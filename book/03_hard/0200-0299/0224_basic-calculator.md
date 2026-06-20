# 0224 — Basic Calculator

## Deskripsi

**Soal:** [0224. Basic Calculator](https://leetcode.com/problems/basic-calculator/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func calculate(s string) int`

## Solusi Go

```go
package main

// LeetCode #224: Basic Calculator
// https://leetcode.com/problems/basic-calculator/
// Difficulty: Hard

import "fmt"

func calculate(s string) int {
  // Membuat slice untuk menyimpan hasil
	stack := make([]int, 0)
	result := 0
	sign := 1
	num := 0

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			num = 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			result += sign * num
			i--
		} else if ch == '+' {
			sign = 1
		} else if ch == '-' {
			sign = -1
		} else if ch == '(' {
			stack = append(stack, result, sign)
			result = 0
			sign = 1
		} else if ch == ')' {
			result = stack[len(stack)-2] + stack[len(stack)-1]*result
			stack = stack[:len(stack)-2]
		}
	}

	return result
}

func main() {
	fmt.Println(calculate("1 + 1"))
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}
```
