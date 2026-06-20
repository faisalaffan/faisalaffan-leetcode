# 1896 — Minimum Cost To Change The Final Value Of Expression

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minCostToChange(expression string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1896: Minimum Cost to Change the Final Value of Expression
// https://leetcode.com/problems/minimum-cost-to-change-the-final-value-of-expression/
// Difficulty: Hard

import (
	"fmt"
)

// Pair represents (min cost to make value 0, min cost to make value 1)
type Pair struct{ zero, one int }

func minCostToChange(expression string) int {
	var valStack []Pair
	var opStack []byte

	apply := func() {
		b := valStack[len(valStack)-1]
		valStack = valStack[:len(valStack)-1]
		a := valStack[len(valStack)-1]
		valStack = valStack[:len(valStack)-1]
		op := opStack[len(opStack)-1]
		opStack = opStack[:len(opStack)-1]

		var res Pair
		if op == '&' {
			// cost to make 0:
			// already 0 on left: cost 0 (any right) or change left from 0 to 0
			// We can make left=0 (a.zero) OR right=0 (b.zero) or both
			// Actually for AND: result=0 if any operand is 0
			// Option 1: make left 0 (cost a.zero)
			// Option 2: make right 0 (cost b.zero)
			// Option 2b: change operator (cost 1) but we need to track toggle ops too
			// Option 3: change operator to |, then result 0 if both 0: cost = a.zero + b.zero + 1

			// Make 0: either operand is 0
			res.zero = min(a.zero, b.zero)
			// Also can change operator: a.zero + b.zero + 1 (change & to |, make both 0)
			res.zero = min(res.zero, a.zero+b.zero+1)

			// Make 1: both operands must be 1
			res.one = a.one + b.one
			// Or change operator: min(a.one, b.one) + 1 (change & to |, make one 1)
			res.one = min(res.one, min(a.one, b.one)+1)
		} else { // '|'
			// Make 0: both operands must be 0
			res.zero = a.zero + b.zero
			// Or change operator: min(a.zero, b.zero) + 1 (change | to &, make one 0)
			res.zero = min(res.zero, min(a.zero, b.zero)+1)

			// Make 1: either operand is 1
			res.one = min(a.one, b.one)
			// Or change operator: a.one + b.one + 1 (change | to &, make both 1)
			res.one = min(res.one, a.one+b.one+1)
		}
		valStack = append(valStack, res)
	}

	for _, ch := range expression {
		switch ch {
		case '(':
			opStack = append(opStack, '(')
		case ')':
			for len(opStack) > 0 && opStack[len(opStack)-1] != '(' {
				apply()
			}
			opStack = opStack[:len(opStack)-1] // pop '('
		case '&', '|':
			for len(opStack) > 0 && opStack[len(opStack)-1] != '(' && opStack[len(opStack)-1] != '&' && opStack[len(opStack)-1] != '|' {
				// precedence: same-level operators processed left to right
				// but actually & and | have the same precedence in this problem
				// We process left to right when same operator type
			}
			// Actually for simplicity, since we only have & and | with same precedence,
			// we just process left to right. We'll just evaluate when we see an operator
			// if the previous one was also an operator.
			for len(opStack) > 0 && opStack[len(opStack)-1] != '(' {
				apply()
			}
			opStack = append(opStack, byte(ch))
		case '0', '1':
			v := int(ch - '0')
			if v == 0 {
				valStack = append(valStack, Pair{0, 1}) // cost 0 to make 0, cost 1 to make 1
			} else {
				valStack = append(valStack, Pair{1, 0}) // cost 1 to make 0, cost 0 to make 1
			}
		}
	}

	for len(opStack) > 0 {
		apply()
	}

	res := valStack[len(valStack)-1]
	// The input expression evaluates to some value.
	// We need the min cost to flip that value.
	// First, determine the current value.
	// Re-evaluate the expression without the cost to find current value:
	currVal := evaluate(expression)
	if currVal == 0 {
		return res.one
	}
	return res.zero
}

func evaluate(expr string) int {
	// Simple stack-based evaluation
	var vals []int
	var ops []byte
	for _, ch := range expr {
		switch ch {
		case '(':
			ops = append(ops, '(')
		case ')':
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				b := vals[len(vals)-1]
				vals = vals[:len(vals)-1]
				a := vals[len(vals)-1]
				vals = vals[:len(vals)-1]
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1]
				if op == '&' {
					vals = append(vals, a&b)
				} else {
					vals = append(vals, a|b)
				}
			}
			ops = ops[:len(ops)-1] // pop '('
		case '0', '1':
			vals = append(vals, int(ch-'0'))
		case '&', '|':
			ops = append(ops, byte(ch))
		}
	}
	return vals[0]
}

func main() {
	// Example: "1&(0|1)" -> 1 (change & to |)
	fmt.Println(minCostToChange("1&(0|1)"))

	// Additional tests
	fmt.Println(minCostToChange("0&0"))
	fmt.Println(minCostToChange("1|0"))
}
```
