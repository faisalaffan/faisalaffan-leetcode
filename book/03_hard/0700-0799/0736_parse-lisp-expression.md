# 0736 — Parse Lisp Expression

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func evaluate(expression string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #736: Parse Lisp Expression
// https://leetcode.com/problems/parse-lisp-expression/
// Difficulty: Hard
//
// Algorithm: Recursive Descent with Scope Stack
// Supports three operations:
//   (let v1 e1 v2 e2 ... body) - bind variables, evaluate body
//   (add e1 e2) - evaluate e1 + e2 (both must be numbers)
//   (mult e1 e2) - evaluate e1 * e2 (both must be numbers)
//   integer - return the integer value
//   variable - look up variable in current scope(s)
//
// Variables are scoped: let creates a new scope that shadows outer scopes.

import (
	"fmt"
	"strconv"
	"unicode"
)

type ExprParser struct {
	s      string
	pos    int
	scopes []map[string]int
}

func (p *ExprParser) peek() byte {
	if p.pos >= len(p.s) {
		return 0
	}
	return p.s[p.pos]
}

func (p *ExprParser) consume() byte {
	c := p.peek()
	p.pos++
	return c
}

func (p *ExprParser) skipWhitespace() {
	for p.pos < len(p.s) && (p.s[p.pos] == ' ' || p.s[p.pos] == '\t' || p.s[p.pos] == '\n') {
		p.pos++
	}
}

func (p *ExprParser) parseIdent() string {
	start := p.pos
	for p.pos < len(p.s) && (unicode.IsLetter(rune(p.s[p.pos])) || unicode.IsDigit(rune(p.s[p.pos]))) {
		p.pos++
	}
	return p.s[start:p.pos]
}

func (p *ExprParser) parseInt() int {
	start := p.pos
	if p.peek() == '-' {
		p.pos++
	}
	for p.pos < len(p.s) && unicode.IsDigit(rune(p.s[p.pos])) {
		p.pos++
	}
	num, _ := strconv.Atoi(p.s[start:p.pos])
	return num
}

func (p *ExprParser) parse() int {
	p.skipWhitespace()

	if p.peek() != '(' {
		// Token: integer or variable
		if p.peek() == '-' || unicode.IsDigit(rune(p.peek())) {
			return p.parseInt()
		}
		// Variable
		name := p.parseIdent()
		// Look up in scopes from innermost to outermost
		for i := len(p.scopes) - 1; i >= 0; i-- {
			if val, ok := p.scopes[i][name]; ok {
				return val
			}
		}
		return 0 // should not happen for valid input
	}

	// It's a parenthesized expression
	p.consume() // consume '('
	p.skipWhitespace()

	keyword := p.parseIdent()
	p.skipWhitespace()

	var result int

	switch keyword {
	case "let":
		// Push a new scope for this let
		p.scopes = append(p.scopes, make(map[string]int))
		// Parse (var expr)* body
		for {
			p.skipWhitespace()
			if p.peek() == ')' {
				break
			}
			savePos := p.pos
			name := p.parseIdent()
			p.skipWhitespace()
			if p.peek() == ')' {
				// name is the body (variable reference)
				for i := len(p.scopes) - 1; i >= 0; i-- {
					if val, ok := p.scopes[i][name]; ok {
						result = val
						break
					}
				}
				break
			}
			if name != "" {
				// name is a variable in a binding: name = expr
				val := p.parse()
				p.scopes[len(p.scopes)-1][name] = val
			} else {
				// Not an identifier, must be body expression
				p.pos = savePos
				result = p.parse()
				break
			}
		}
		// Pop scope
		p.scopes = p.scopes[:len(p.scopes)-1]

	case "add":
		e1 := p.parse()
		p.skipWhitespace()
		e2 := p.parse()
		result = e1 + e2

	case "mult":
		e1 := p.parse()
		p.skipWhitespace()
		e2 := p.parse()
		result = e1 * e2
	}

	p.skipWhitespace()
	p.consume() // consume ')'
	return result
}

func evaluate(expression string) int {
	parser := &ExprParser{
		s:      expression,
		pos:    0,
		scopes: []map[string]int{},
	}
	return parser.parse()
}

func main() {
	// Example from problem
	expr1 := "(let x 2 (mult x 5))"
	result1 := evaluate(expr1)
	fmt.Printf("Input: %s\nOutput: %d (expected: 10)\n\n", expr1, result1)

	// Test case 2
	expr2 := "(let x 2 (mult x (let x 3 (add x x))))"
	result2 := evaluate(expr2)
	fmt.Printf("Input: %s\nOutput: %d (expected: 12)\n\n", expr2, result2)

	// Test case 3
	expr3 := "(let a1 3 b2 (add a1 1) b2)"
	result3 := evaluate(expr3)
	fmt.Printf("Input: %s\nOutput: %d (expected: 4)\n\n", expr3, result3)

	// Test case 4: add
	expr4 := "(add 1 2)"
	result4 := evaluate(expr4)
	fmt.Printf("Input: %s\nOutput: %d (expected: 3)\n\n", expr4, result4)

	// Test case 5: mult
	expr5 := "(mult 3 7)"
	result5 := evaluate(expr5)
	fmt.Printf("Input: %s\nOutput: %d (expected: 21)\n\n", expr5, result5)

	// Test case 6: nested operations
	expr6 := "(mult (add 2 3) (add 4 5))"
	result6 := evaluate(expr6)
	fmt.Printf("Input: %s\nOutput: %d (expected: 45)\n\n", expr6, result6)

	// Test case 7: plain number
	expr7 := "123"
	result7 := evaluate(expr7)
	fmt.Printf("Input: %s\nOutput: %d (expected: 123)\n\n", expr7, result7)

	// Test case 8: negative numbers
	expr8 := "(let x -5 (add x 3))"
	result8 := evaluate(expr8)
	fmt.Printf("Input: %s\nOutput: %d (expected: -2)\n\n", expr8, result8)

	// Test case 9: variable shadowing
	expr9 := "(let x 1 (let x 2 x))"
	result9 := evaluate(expr9)
	fmt.Printf("Input: %s\nOutput: %d (expected: 2)\n\n", expr9, result9)

	// Test case 10
	expr10 := "(let x 1 (let x 2 (let x 3 x)))"
	result10 := evaluate(expr10)
	fmt.Printf("Input: %s\nOutput: %d (expected: 3)\n", expr10, result10)
}
```
