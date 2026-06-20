# 1096 — Brace Expansion Ii

## Deskripsi

**Soal:** [1096. Brace Expansion Ii](https://leetcode.com/problems/brace-expansion-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1096: Brace Expansion II
// https://leetcode.com/problems/brace-expansion-ii/
// Difficulty: Hard
//
// Recursive descent parser for grammar:
//   expr   → term (',' term)*
//   term   → factor (factor)*   (concatenation = product)
//   factor → '{' expr '}' | letter+
// Returns sorted unique strings.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(braceExpansionII("{a,b}{c,{d,e}}"))
}

func braceExpansionII(expression string) []string {
	type parser func() []string

	idx := 0

	var parseExpr, parseTerm, parseFactor parser

	parseExpr = func() []string {
		result := parseTerm()
		for idx < len(expression) && expression[idx] == ',' {
			idx++
			result = union(result, parseTerm())
		}
		return result
	}

	parseTerm = func() []string {
		result := []string{""}
		for idx < len(expression) && (expression[idx] == '{' || isLetter(expression[idx])) {
			result = product(result, parseFactor())
		}
		return result
	}

	parseFactor = func() []string {
		if expression[idx] == '{' {
			idx++ // skip '{'
			result := parseExpr()
			idx++ // skip '}'
			return result
		}
		// read consecutive letters
		j := idx
		for j < len(expression) && isLetter(expression[j]) {
			j++
		}
		result := []string{expression[idx:j]}
		idx = j
		return result
	}

	result := parseExpr()
	sort.Strings(result)
	return result
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z'
}

func union(a, b []string) []string {
  // Membuat map untuk pencarian O(1): key → value
	set := make(map[string]bool)
	for _, s := range a {
		set[s] = true
	}
	for _, s := range b {
		set[s] = true
	}
  // Membuat slice untuk menyimpan hasil
	res := make([]string, 0, len(set))
	for s := range set {
		res = append(res, s)
	}
	return res
}

func product(a, b []string) []string {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
  // Membuat map untuk pencarian O(1): key → value
	set := make(map[string]bool)
	for _, s1 := range a {
		for _, s2 := range b {
			set[s1+s2] = true
		}
	}
  // Membuat slice untuk menyimpan hasil
	res := make([]string, 0, len(set))
	for s := range set {
		res = append(res, s)
	}
	return res
}
```
