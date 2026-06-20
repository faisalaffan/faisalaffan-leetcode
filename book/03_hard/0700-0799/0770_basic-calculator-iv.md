# 0770 — Basic Calculator Iv

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func basicCalculatorIV(expression string, evalvars []string, evalints []int) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Merge Sort

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #770: Basic Calculator IV
// https://leetcode.com/problems/basic-calculator-iv/
// Difficulty: Hard
//
// Given an expression like "e + 8 - a + 5" with evalvars/evalints,
// substitute variables, evaluate, and return the polynomial result
// as a list of terms sorted by degree (desc) then lexicographically.

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Polynomial represented as map[variableSignature]coefficient
// variableSignature is sorted (e.g., "a*b*c")
type Poly map[string]int

func main() {
	// Example 1
	fmt.Println(basicCalculatorIV("e + 8 - a + 5", []string{"e"}, []int{1}))
	// Output: ["-1*a","14"]

	// Example 2
	fmt.Println(basicCalculatorIV("e - 8 + temperature - pressure",
		[]string{"e", "temperature"}, []int{1, 12}))
	// Output: ["-1*pressure","5"]

	// Example 3
	fmt.Println(basicCalculatorIV("(e + 8) * (e - 8)",
		[]string{}, []int{}))
	// Output: ["1*e*e","-64"]
}

func basicCalculatorIV(expression string, evalvars []string, evalints []int) []string {
	// Build substitution map
  // Membuat map (HashMap) — pencarian O(1)
	subst := make(map[string]int)
	for i, v := range evalvars {
		subst[v] = evalints[i]
	}

	// Tokenize
	tokens := tokenize(expression)

	// Parse and evaluate
	idx := 0
	poly := parseExpr(tokens, &idx, subst)

	// Format result
	return formatPoly(poly)
}

func tokenize(s string) []string {
	var tokens []string
	i := 0
	for i < len(s) {
		if s[i] == ' ' {
			i++
			continue
		}
		if s[i] == '(' || s[i] == ')' || s[i] == '+' || s[i] == '-' || s[i] == '*' {
			tokens = append(tokens, string(s[i]))
			i++
		} else {
			// Number or variable
			j := i
			for j < len(s) && s[j] != ' ' && s[j] != '(' && s[j] != ')' &&
				s[j] != '+' && s[j] != '-' && s[j] != '*' {
				j++
			}
			tokens = append(tokens, s[i:j])
			i = j
		}
	}
	return tokens
}

// Grammar:
// expr   := term (('+' | '-') term)*
// term   := factor (('*' factor)*)
// factor := NUMBER | VARIABLE | '(' expr ')'

func parseExpr(tokens []string, idx *int, subst map[string]int) Poly {
	left := parseTerm(tokens, idx, subst)
	for *idx < len(tokens) {
		op := tokens[*idx]
		if op != "+" && op != "-" {
			break
		}
		*idx++
		right := parseTerm(tokens, idx, subst)
		if op == "+" {
			left = addPoly(left, right)
		} else {
			left = subPoly(left, right)
		}
	}
	return left
}

func parseTerm(tokens []string, idx *int, subst map[string]int) Poly {
	left := parseFactor(tokens, idx, subst)
	for *idx < len(tokens) && tokens[*idx] == "*" {
		*idx++
		right := parseFactor(tokens, idx, subst)
		left = mulPoly(left, right)
	}
	return left
}

func parseFactor(tokens []string, idx *int, subst map[string]int) Poly {
	tok := tokens[*idx]
	*idx++

	if tok == "(" {
		poly := parseExpr(tokens, idx, subst)
		// expect ")"
		*idx++
		return poly
	}

	// Number or variable
	if isNum(tok) {
		n, _ := strconv.Atoi(tok)
		return Poly{"": n}
	}

	// Variable: substitute if in evalvars
	if val, ok := subst[tok]; ok {
		return Poly{"": val}
	}

	// Keep as variable
	return Poly{tok: 1}
}

func isNum(s string) bool {
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return len(s) > 0
}

func addPoly(a, b Poly) Poly {
	res := make(Poly)
	for k, v := range a {
		res[k] = v
	}
	for k, v := range b {
		res[k] += v
	}
	return res
}

func subPoly(a, b Poly) Poly {
	res := make(Poly)
	for k, v := range a {
		res[k] = v
	}
	for k, v := range b {
		res[k] -= v
	}
	return res
}

func mulPoly(a, b Poly) Poly {
	res := make(Poly)
	for ka, va := range a {
		for kb, vb := range b {
			sig := combineSig(ka, kb)
			res[sig] += va * vb
		}
	}
	return res
}

func combineSig(a, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}

	// Merge sorted variable names separated by "*"
	vars := mergeVars(strings.Split(a, "*"), strings.Split(b, "*"))
	return strings.Join(vars, "*")
}

func mergeVars(a, b []string) []string {
	res := make([]string, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else if a[i] > b[j] {
			res = append(res, b[j])
			j++
		} else {
			// For multiplication, keep ALL copies (don't deduplicate).
			// "e" * "e" should produce "e*e" (degree 2).
			res = append(res, a[i])
			res = append(res, b[j])
			i++
			j++
		}
	}
	res = append(res, a[i:]...)
	res = append(res, b[j:]...)
	return res
}

func formatPoly(p Poly) []string {
	// Filter out zero coefficients, collect terms
	var terms []struct {
		sig  string
		coef int
	}
	for sig, coef := range p {
		if coef == 0 {
			continue
		}
		terms = append(terms, struct {
			sig  string
			coef int
		}{sig, coef})
	}

	// Sort by degree desc, then lexicographically
  // Custom sort dengan comparator
	sort.Slice(terms, func(i, j int) bool {
		degI := 0
		if terms[i].sig != "" {
			degI = strings.Count(terms[i].sig, "*") + 1
		}
		degJ := 0
		if terms[j].sig != "" {
			degJ = strings.Count(terms[j].sig, "*") + 1
		}
		if degI != degJ {
			return degI > degJ
		}
		return terms[i].sig < terms[j].sig
	})

	res := make([]string, len(terms))
	for i, t := range terms {
		if t.sig == "" {
			res[i] = strconv.Itoa(t.coef)
		} else {
			res[i] = fmt.Sprintf("%d*%s", t.coef, t.sig)
		}
	}
	return res
}
```
