# 1440 — Evaluate Boolean Expression

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func evaluate(leftVal int, op string, rightVal int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1440: Evaluate Boolean Expression
// https://leetcode.com/problems/evaluate-boolean-expression/
// Difficulty: Medium (listed here as Hard) [Paid]
//
// Given tables Variables (name, value) and Expressions (left_operand, operator, right_operand),
// evaluate each expression. Operators: <, >, =.

import "fmt"

// evaluate returns the result of a boolean expression.
func evaluate(leftVal int, op string, rightVal int) bool {
	switch op {
	case "<":
		return leftVal < rightVal
	case ">":
		return leftVal > rightVal
	case "=":
		return leftVal == rightVal
	default:
		return false
	}
}

// evaluateExpressions evaluates all expressions given variable values.
func evaluateExpressions(variables map[string]int, expressions []struct {
	Left  string
	Op    string
	Right string
}) []string {
	var results []string
	for _, expr := range expressions {
		leftVal, leftOk := variables[expr.Left]
		rightVal, rightOk := variables[expr.Right]
		if !leftOk || !rightOk {
			results = append(results, "unknown")
			continue
		}
		result := evaluate(leftVal, expr.Op, rightVal)
		if result {
			results = append(results, "true")
		} else {
			results = append(results, "false")
		}
	}
	return results
}

func main() {
	// Sample variables
	variables := map[string]int{
		"x": 10,
		"y": 20,
		"z": 10,
	}

	expressions := []struct {
		Left  string
		Op    string
		Right string
	}{
		{"x", ">", "y"},
		{"x", "<", "y"},
		{"x", "=", "z"},
		{"y", ">", "z"},
		{"x", "=", "y"},
	}

	results := evaluateExpressions(variables, expressions)
	fmt.Println("Expression results:")
	for i, expr := range expressions {
		fmt.Printf("  %s %s %s => %s\n", expr.Left, expr.Op, expr.Right, results[i])
	}

	// Test 2: Edge cases
	variables2 := map[string]int{
		"a": 0,
		"b": -5,
		"c": 100,
	}

	expressions2 := []struct {
		Left  string
		Op    string
		Right string
	}{
		{"a", ">", "b"},
		{"b", "<", "c"},
		{"a", "=", "b"},
		{"a", "=", "a"},
		{"c", ">", "c"},
		{"c", "=", "c"},
	}

	results2 := evaluateExpressions(variables2, expressions2)
	fmt.Println("\nExpression results 2:")
	for i, expr := range expressions2 {
		fmt.Printf("  %s %s %s => %s\n", expr.Left, expr.Op, expr.Right, results2[i])
	}

	// Test 3: Missing variable
	variables3 := map[string]int{
		"x": 5,
	}
	expressions3 := []struct {
		Left  string
		Op    string
		Right string
	}{
		{"x", ">", "y"},
	}
	results3 := evaluateExpressions(variables3, expressions3)
	fmt.Println("\nExpression results 3 (missing var):")
	for i, expr := range expressions3 {
		fmt.Printf("  %s %s %s => %s\n", expr.Left, expr.Op, expr.Right, results3[i])
	}
}
```
