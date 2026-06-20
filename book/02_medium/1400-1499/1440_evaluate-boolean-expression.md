# 1440 — Evaluate Boolean Expression

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func evaluateBoolean(expressions []struct { leftOperand string operator string rightOperand string }, variables []struct { name string value int }) []evalResult`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n) where n = number of expressions  |  **Ruang:** O(k) where k = number of variables

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1440: Evaluate Boolean Expression
// https://leetcode.com/problems/evaluate-boolean-expression/
// Difficulty: Medium

import "fmt"

func main() {
	// SQL problem - simulating in Go
	result := evaluateBoolean(
		[]struct {
			leftOperand  string
			operator     string
			rightOperand string
		}{
			{"x", ">", "y"},
			{"x", "<", "y"},
			{"x", "=", "y"},
			{"y", ">", "x"},
			{"y", "<", "x"},
			{"x", "=", "x"},
		},
		[]struct {
			name  string
			value int
		}{
			{"x", 66},
			{"y", 77},
			{"z", 88},
		},
	)
	for _, r := range result {
		fmt.Printf("%s %s %s %t\n", r.left, r.op, r.right, r.value)
	}
}

type evalResult struct {
	left, op, right string
	value           bool
}

// Time: O(n) where n = number of expressions
// Space: O(k) where k = number of variables
func evaluateBoolean(expressions []struct {
	leftOperand  string
	operator     string
	rightOperand string
}, variables []struct {
	name  string
	value int
}) []evalResult {
  // HashMap: O(1) lookup
	varMap := make(map[string]int)
	for _, v := range variables {
		varMap[v.name] = v.value
	}

	var results []evalResult
	for _, e := range expressions {
		leftVal := varMap[e.leftOperand]
		rightVal := varMap[e.rightOperand]
		var val bool

		switch e.operator {
		case ">":
			val = leftVal > rightVal
		case "<":
			val = leftVal < rightVal
		case "=":
			val = leftVal == rightVal
		case "!=":
			val = leftVal != rightVal
		case ">=":
			val = leftVal >= rightVal
		case "<=":
			val = leftVal <= rightVal
		}

		results = append(results, evalResult{e.leftOperand, e.operator, e.rightOperand, val})
	}

	return results
}
```
