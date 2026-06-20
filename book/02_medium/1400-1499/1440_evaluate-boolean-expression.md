# 1440 — Evaluate Boolean Expression

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func evaluateBoolean(expressions []struct {
	leftOperand  string
	operator     string
	rightOperand string
}, variables []struct {
	name  string
	value int
}) []evalResult
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) where n = number of expressions  
**Kompleksitas Ruang:** O(k) where k = number of variables

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Membuat map (HashMap) — pencarian O(1)
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
