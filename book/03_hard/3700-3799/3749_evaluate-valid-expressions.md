# 3749 — Evaluate Valid Expressions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func evaluateExpression(expression string) int64
```

> **💡 Hint:** Two-stack algorithm (Shunting-yard / recursive descent).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3749: Evaluate Valid Expressions
// https://leetcode.com/problems/evaluate-valid-expressions/
// Difficulty: Hard [Paid]
//
// Evaluate a mathematical expression with +, -, *, and parentheses.
// All operations are integer-based. No division.
//
// Approach: Two-stack algorithm (Shunting-yard / recursive descent).

import (
	"fmt"
	"unicode"
)

func main() {
	// Example 1
	fmt.Println(evaluateExpression("1+2*3"))
	// Example 2
	fmt.Println(evaluateExpression("(1+2)*3"))
	// Example 3
	fmt.Println(evaluateExpression("2*(3+4)-5"))
	// Edge: single number
	fmt.Println(evaluateExpression("42"))
	// Edge: nested parens
	fmt.Println(evaluateExpression("((1+2)*(3+4))"))
}

func evaluateExpression(expression string) int64 {
	if len(expression) == 0 {
		return 0
	}
	return int64(calc(expression))
}

func calc(s string) int {
	n := len(s)
  // Alokasi slice integer
	nums := make([]int, 0)
	ops := make([]byte, 0)
	prec := func(op byte) int {
		if op == '+' || op == '-' {
			return 1
		}
		if op == '*' {
			return 2
		}
		return 0
	}
	apply := func() int {
		b := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		a := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]
		switch op {
		case '+':
			return a + b
		case '-':
			return a - b
		case '*':
			return a * b
		}
		return 0
	}

	for i := 0; i < n; i++ {
		ch := s[i]
		if ch == ' ' {
			continue
		}
		if unicode.IsDigit(rune(ch)) {
			val := 0
			for i < n && unicode.IsDigit(rune(s[i])) {
				val = val*10 + int(s[i]-'0')
				i++
			}
			nums = append(nums, val)
			i--
		} else if ch == '(' {
			ops = append(ops, ch)
		} else if ch == ')' {
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				nums = append(nums, apply())
			}
			if len(ops) > 0 {
				ops = ops[:len(ops)-1] // pop '('
			}
		} else {
			// +, -, *
			for len(ops) > 0 && prec(ops[len(ops)-1]) >= prec(ch) {
				nums = append(nums, apply())
			}
			ops = append(ops, ch)
		}
	}

	for len(ops) > 0 {
		nums = append(nums, apply())
	}

	return nums[0]
}
```
