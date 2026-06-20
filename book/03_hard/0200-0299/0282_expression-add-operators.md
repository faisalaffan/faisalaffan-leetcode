# 0282 — Expression Add Operators

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func addOperators(num string, target int) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Backtracking

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Backtracking** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #282: Expression Add Operators
// https://leetcode.com/problems/expression-add-operators/
// Difficulty: Hard
//
// Approach: Backtracking.
//  1. Try all possible splits of the string into operands.
//  2. At each step, try '+', '-', '*' operators.
//  3. For '*', we need to track the last operand to handle precedence:
//     current result = (result - lastOperand) + (lastOperand * currentNum)
//  4. Avoid numbers with leading zeros (e.g., "05" is invalid).

import (
	"fmt"
	"strconv"
)

func main() {
	// Example 1: "123", 6 -> ["1+2+3","1*2*3"]
	fmt.Println("\"123\", 6 ->", addOperators("123", 6))

	// Example 2: "232", 8 -> ["2*3+2","2+3*2"]
	fmt.Println("\"232\", 8 ->", addOperators("232", 8))

	// Example 3: "3456237490", 9191 -> []
	fmt.Println("\"3456237490\", 9191 ->", addOperators("3456237490", 9191))

	// Edge: "105", 5 -> ["1*0+5","10-5"]
	fmt.Println("\"105\", 5 ->", addOperators("105", 5))

	// Edge: "00", 0 -> ["0+0","0-0","0*0"]
	fmt.Println("\"00\", 0 ->", addOperators("00", 0))
}

// addOperators returns all possible expressions that evaluate to target.
func addOperators(num string, target int) []string {
	var result []string
	if len(num) == 0 {
		return result
	}
	backtrack(num, target, 0, 0, 0, "", &result)
	return result
}

// backtrack explores all possible expression constructions.
//   - num: remaining string to process
//   - target: target value
//   - index: current position in num
//   - currentValue: value of the expression built so far
//   - lastOperand: last operand added to the expression (for multiplication precedence)
//   - expression: the expression string built so far
//   - result: collects valid expressions
func backtrack(num string, target int, index int, currentValue int, lastOperand int, expression string, result *[]string) {
	if index == len(num) {
		if currentValue == target {
			*result = append(*result, expression)
		}
		return
	}

	for i := index; i < len(num); i++ {
		// Avoid numbers with leading zeros.
		if i > index && num[index] == '0' {
			break
		}

		// Parse current number from num[index:i+1].
		currentNum, _ := strconv.Atoi(num[index : i+1])

		if index == 0 {
			// First operand: no operator needed.
			backtrack(num, target, i+1, currentNum, currentNum, strconv.Itoa(currentNum), result)
		} else {
			// Try addition.
			backtrack(num, target, i+1, currentValue+currentNum, currentNum, expression+"+"+strconv.Itoa(currentNum), result)

			// Try subtraction.
			backtrack(num, target, i+1, currentValue-currentNum, -currentNum, expression+"-"+strconv.Itoa(currentNum), result)

			// Try multiplication: undo last operation, then multiply.
			// newValue = currentValue - lastOperand + (lastOperand * currentNum)
			backtrack(num, target, i+1, currentValue-lastOperand+lastOperand*currentNum, lastOperand*currentNum, expression+"*"+strconv.Itoa(currentNum), result)
		}
	}
}

// Stub compatibility.
func ExpressionAddOperators() any {
	return addOperators("123", 6)
}
```
