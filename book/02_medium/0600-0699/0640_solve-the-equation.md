# 0640 — Solve The Equation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func solveEquation(equation string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n) where n is length of equation string  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #640: Solve the Equation
// https://leetcode.com/problems/solve-the-equation/
// Difficulty: Medium
// Time: O(n) where n is length of equation string
// Space: O(1)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(solveEquation("x+5-3+x=6+x-2"))
	fmt.Println(solveEquation("x=x"))
	fmt.Println(solveEquation("2x=x"))
}

func solveEquation(equation string) string {
	eqIdx := findEqual(equation)
	leftCoeff, leftConst := parse(equation[:eqIdx])
	rightCoeff, rightConst := parse(equation[eqIdx+1:])

	coeff := leftCoeff - rightCoeff
	constant := rightConst - leftConst

	if coeff == 0 && constant == 0 {
		return "Infinite solutions"
	}
	if coeff == 0 {
		return "No solution"
	}
	return "x=" + strconv.Itoa(constant/coeff)
}

func findEqual(s string) int {
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if s[i] == '=' {
			return i
		}
	}
	return -1
}

func parse(expr string) (int, int) {
	coeff, constant := 0, 0
	sign := 1
	i := 0
	n := len(expr)

	for i < n {
		if expr[i] == '+' {
			sign = 1
			i++
		} else if expr[i] == '-' {
			sign = -1
			i++
		} else {
			start := i
			for i < n && expr[i] >= '0' && expr[i] <= '9' {
				i++
			}
			if i < n && expr[i] == 'x' {
				if start == i {
					coeff += sign
				} else {
					val, _ := strconv.Atoi(expr[start:i])
					coeff += sign * val
				}
				i++
			} else {
				val, _ := strconv.Atoi(expr[start:i])
				constant += sign * val
			}
		}
	}
	return coeff, constant
}
```
