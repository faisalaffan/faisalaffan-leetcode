# 0726 — Number Of Atoms

## Deskripsi

**Soal:** [0726. Number Of Atoms](https://leetcode.com/problems/number-of-atoms/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func countOfAtoms(formula string) string`

## Solusi Go

```go
package main

// LeetCode #726: Number of Atoms
// https://leetcode.com/problems/number-of-atoms/
// Difficulty: Hard
//
// Algorithm: Stack-based parsing
// 1. Parse the formula string
// 2. Use a stack of maps to handle nested parentheses
// 3. When we see '(', push a new map onto the stack
// 4. When we see ')', pop the top map and multiply by the following number
// 5. When we see an element, add it to the current map
// 6. At the end, combine all maps and sort by element name

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func countOfAtoms(formula string) string {
	// Stack of maps: each map is element -> count
	stack := []map[string]int{{}}
	i := 0
	n := len(formula)

	for i < n {
		c := formula[i]

		if c == '(' {
			// Push a new map for the new scope
			stack = append(stack, map[string]int{})
			i++
		} else if c == ')' {
			// Pop the top map
			i++
			numStart := i
			for i < n && unicode.IsDigit(rune(formula[i])) {
				i++
			}
			multiplier := 1
			if numStart < i {
				multiplier, _ = strconv.Atoi(formula[numStart:i])
			}

			// Pop the top map and multiply, then merge into the previous map
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for elem, count := range top {
				stack[len(stack)-1][elem] += count * multiplier
			}
		} else if unicode.IsUpper(rune(c)) {
			// Parse element name
			i++
			elemStart := i - 1
			for i < n && unicode.IsLower(rune(formula[i])) {
				i++
			}
			elemName := formula[elemStart:i]

			// Parse count
			numStart := i
			for i < n && unicode.IsDigit(rune(formula[i])) {
				i++
			}
			count := 1
			if numStart < i {
				count, _ = strconv.Atoi(formula[numStart:i])
			}

			stack[len(stack)-1][elemName] += count
		} else {
			// Should not happen for valid input
			i++
		}
	}

	// Final map is the last (and only) item on stack
	elemCount := stack[0]

	// Sort element names
  // Membuat slice untuk menyimpan hasil
	elements := make([]string, 0, len(elemCount))
	for elem := range elemCount {
		elements = append(elements, elem)
	}
	sort.Strings(elements)

	var sb strings.Builder
	for _, elem := range elements {
		sb.WriteString(elem)
		if elemCount[elem] > 1 {
			sb.WriteString(strconv.Itoa(elemCount[elem]))
		}
	}

	return sb.String()
}

func main() {
	// Example from problem
	formula1 := "Mg(OH)2"
	result1 := countOfAtoms(formula1)
	fmt.Printf("Input: %s\nOutput: %s (expected: H2MgO2)\n\n", formula1, result1)

	// Test case 2
	formula2 := "H2O"
	result2 := countOfAtoms(formula2)
	fmt.Printf("Input: %s\nOutput: %s (expected: H2O)\n\n", formula2, result2)

	// Test case 3
	formula3 := "K4(ON(SO3)2)2"
	result3 := countOfAtoms(formula3)
	fmt.Printf("Input: %s\nOutput: %s (expected: K4N2O14S4)\n\n", formula3, result3)

	// Test case 4
	formula4 := "Be32"
	result4 := countOfAtoms(formula4)
	fmt.Printf("Input: %s\nOutput: %s (expected: Be32)\n\n", formula4, result4)

	// Test case 5: nested parentheses
	formula5 := "((H)2(O))3"
	result5 := countOfAtoms(formula5)
	fmt.Printf("Input: %s\nOutput: %s (expected: H6O3)\n\n", formula5, result5)

	// Test case 6: no parentheses
	formula6 := "NaCl"
	result6 := countOfAtoms(formula6)
	fmt.Printf("Input: %s\nOutput: %s (expected: ClNa)\n\n", formula6, result6)

	// Test case 7: complex example
	formula7 := "Mg(H2O)N"
	result7 := countOfAtoms(formula7)
	fmt.Printf("Input: %s\nOutput: %s\n", formula7, result7)
}
```
