# 2118 — Build The Equation

## Deskripsi

**Soal:** [2118. Build The Equation](https://leetcode.com/problems/build-the-equation/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** LIS (Longest Increasing Subsequence)

## Solusi Go

```go
package main

// LeetCode #2118: Build the Equation
// https://leetcode.com/problems/build-the-equation/
// Difficulty: Hard [Paid]
//
// Given a list of terms (power, coefficient), build the polynomial equation
// string. Terms are sorted by power descending. Format: "+2x^3-4x^1+5" = "0".
// Handle special cases: coefficient = 1 or -1, power = 0 or 1.

import (
	"fmt"
	"sort"
)

// term represents a polynomial term with coefficient and power
type term struct {
	coefficient int
	power       int
}

func main() {
	// Example: y = 2x^3 - 4x + 5  =>  "2x^3-4x+5=0"
	fmt.Println(buildEquation([]term{{2, 3}, {-4, 1}, {5, 0}}))

	// Example: y = -x^2 + x - 1 =>  "-x^2+x-1=0"
	fmt.Println(buildEquation([]term{{-1, 2}, {1, 1}, {-1, 0}}))

	// Example: y = 3x^2 =>  "3x^2=0"
	fmt.Println(buildEquation([]term{{3, 2}}))

	// Single constant
	fmt.Println(buildEquation([]term{{7, 0}}))

	// No terms (empty equation)
	fmt.Println(buildEquation([]term{}))
}

func buildEquation(terms []term) string {
	if len(terms) == 0 {
		return "0=0"
	}

	// Sort by power descending
	sort.Slice(terms, func(i, j int) bool {
		return terms[i].power > terms[j].power
	})

	var result string

	for i, t := range terms {
		coef := t.coefficient
		pow := t.power

		if coef == 0 {
			continue
		}

		// Sign handling
		if i == 0 {
			// First term: no leading '+' for positive
			if coef < 0 {
				result += "-"
			}
		} else {
			if coef > 0 {
				result += "+"
			} else {
				result += "-"
			}
		}

		absCoef := coef
		if absCoef < 0 {
			absCoef = -absCoef
		}

		// Coefficient part (don't print 1 or -1 unless power is 0)
		if pow == 0 || absCoef != 1 {
			result += fmt.Sprintf("%d", absCoef)
		}

		// Variable part
		if pow > 0 {
			result += "x"
			if pow > 1 {
				result += fmt.Sprintf("^%d", pow)
			}
		}
	}

	if result == "" {
		return "0=0"
	}

	return result + "=0"
}
```
