# 0631 — Design Excel Sum Formula

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diminta mendesain struktur data kustom dengan operasi spesifik (insert, delete, search). Target: O(1) atau O(log n) per operasi.

**Cara berpikir:** Kombinasikan HashMap + Heap + Linked List sesuai kebutuhan.

**Fungsi Solusi:** `func Constructor(H int, W byte) *Excel`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// LeetCode #631: Design Excel Sum Formula
// https://leetcode.com/problems/design-excel-sum-formula/
// Difficulty: Hard [Paid]
//
// Design an Excel spreadsheet with sum formula support.
// Cells are referenced like "A1", "B2", etc. (col letter + row number, 1-indexed).
// Sum formulas reference ranges like "A1:A3" or individual cells like "A1,C1:D3".

// Excel implements the spreadsheet.
type Excel struct {
	H        int // height (rows)
	W        int // width (cols)
	values   map[string]int    // cell -> direct value
	formulas map[string]string // cell -> sum formula string (e.g., "A1:A3")
	sumCache map[string]int    // cell -> cached sum value
}

// Constructor creates a new Excel with H rows and W cols (A..W).
func Constructor(H int, W byte) *Excel {
	return &Excel{
		H:        H,
		W:        int(W - 'A' + 1),
		values:   make(map[string]int),
		formulas: make(map[string]string),
		sumCache: make(map[string]int),
	}
}

// cellKey converts (row, col) to "A1" format.
// Col is 0-indexed (0 = A), row is 1-indexed.
func cellKey(row int, col int) string {
	return string(rune('A'+col)) + strconv.Itoa(row)
}

// parseCell parses "A1" into (row, col) where col is 0-indexed, row is 1-indexed.
func parseCell(s string) (row int, col int) {
	// Find where the digits start.
	i := 0
	for i < len(s) && unicode.IsLetter(rune(s[i])) {
		i++
	}
	col = int(s[0]-'A') // supports A-Z
	row, _ = strconv.Atoi(s[i:])
	return
}

// parseRange parses a formula like "A1:A3" or "A1,C1:D3,E2".
// Returns list of individual cell keys.
func parseRange(formula string) []string {
	var cells []string
	// Split by comma for top-level range references.
	parts := strings.Split(formula, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if colonIdx := strings.Index(part, ":"); colonIdx >= 0 {
			// Range like "A1:A3" or "A1:C3" or "A1:B2"
			start := part[:colonIdx]
			end := part[colonIdx+1:]
			r1, c1 := parseCell(start)
			r2, c2 := parseCell(end)
			// Iterate over the range.
			for r := r1; r <= r2; r++ {
				for c := c1; c <= c2; c++ {
					cells = append(cells, cellKey(r, c))
				}
			}
		} else {
			// Single cell like "A1".
			cells = append(cells, part)
		}
	}
	return cells
}

// evaluate computes the value of a cell, recursively evaluating sum formulas.
func (ex *Excel) evaluate(cell string) int {
	if formula, ok := ex.formulas[cell]; ok {
		// This cell has a sum formula.
		if val, cached := ex.sumCache[cell]; cached {
			return val
		}
		total := 0
		for _, ref := range parseRange(formula) {
			total += ex.evaluate(ref)
		}
		ex.sumCache[cell] = total
		return total
	}
	// Direct value.
	return ex.values[cell]
}

// Set sets the value of cell (row, col). Row is 1-indexed, col is 0-indexed.
func (ex *Excel) Set(row int, col int, val int) {
	cell := cellKey(row, col)
	ex.values[cell] = val
	delete(ex.formulas, cell)
	ex.invalidateCache()
}

// Get returns the value of cell (row, col).
func (ex *Excel) Get(row int, col int) int {
	return ex.evaluate(cellKey(row, col))
}

// Sum sets a sum formula on cell (row, col) and returns the computed sum.
// numbers is a comma-separated list of cells/ranges like "A1:A3" or "A1,C1:D3".
// Row is 1-indexed, col is 0-indexed.
func (ex *Excel) Sum(row int, col int, numbers []string) int {
	formula := strings.Join(numbers, ",")
	cell := cellKey(row, col)
	ex.formulas[cell] = formula
	delete(ex.values, cell)
	ex.invalidateCache()
	return ex.evaluate(cell)
}

// invalidateCache clears all cached sum values (must be called when any cell changes).
func (ex *Excel) invalidateCache() {
	ex.sumCache = make(map[string]int)
}

// --- Tests -------------------------------------------------------------------

func main() {
	fmt.Println("=== 0631 Design Excel Sum Formula ===")

	// Test 1: Basic get/set.
	ex := Constructor(5, 'E') // 5 rows, 5 cols (A-E)
	ex.Set(1, 0, 10)          // A1 = 10
	ex.Set(2, 0, 20)          // A2 = 20
	fmt.Printf("Test 1 - A1 = %d (expected 10)\n", ex.Get(1, 0))

	// Test 2: Single cell sum.
	ex2 := Constructor(3, 'C')
	ex2.Set(1, 0, 5) // A1 = 5
	ex2.Set(2, 0, 3) // A2 = 3
	sum := ex2.Sum(3, 0, []string{"A1", "A2"}) // A3 = A1+A2
	fmt.Printf("Test 2 - A3 sum = %d (expected 8)\n", sum)

	// Test 3: Range sum.
	ex3 := Constructor(5, 'C')
	ex3.Set(1, 0, 1) // A1 = 1
	ex3.Set(2, 0, 2) // A2 = 2
	ex3.Set(3, 0, 3) // A3 = 3
	ex3.Set(4, 0, 4) // A4 = 4
	sum3 := ex3.Sum(5, 0, []string{"A1:A4"}) // A5 = A1+A2+A3+A4 = 10
	fmt.Printf("Test 3 - A5 range sum = %d (expected 10)\n", sum3)

	// Test 4: Dependencies update after source cell changes.
	ex4 := Constructor(3, 'C')
	ex4.Set(1, 0, 10) // A1 = 10
	ex4.Sum(2, 0, []string{"A1"}) // A2 = A1 = 10
	fmt.Printf("Test 4a - A2 = %d (expected 10)\n", ex4.Get(2, 0))
	ex4.Set(1, 0, 25) // A1 = 25 (should invalidate A2's cache)
	fmt.Printf("Test 4b - A2 after A1 change = %d (expected 25)\n", ex4.Get(2, 0))

	// Test 5: Nested sums (A3 = A1 + A2, where A2 = sum(B1:B2)).
	ex5 := Constructor(5, 'E')
	ex5.Set(1, 0, 5) // A1 = 5
	ex5.Set(1, 1, 3) // B1 = 3
	ex5.Set(2, 1, 7) // B2 = 7
	ex5.Sum(2, 0, []string{"B1", "B2"}) // A2 = B1+B2 = 10
	nested := ex5.Sum(3, 0, []string{"A1", "A2"}) // A3 = A1+A2 = 15
	fmt.Printf("Test 5 - A3 nested sum = %d (expected 15)\n", nested)

	// Test 6: Multi-range formula (A1 + C1:D3).
	ex6 := Constructor(5, 'E')
	ex6.Set(1, 0, 100) // A1 = 100
	ex6.Set(1, 2, 10)  // C1 = 10
	ex6.Set(2, 2, 20)  // C2 = 20
	ex6.Set(3, 2, 30)  // C3 = 30
	sum6 := ex6.Sum(4, 0, []string{"A1", "C1:C3"}) // A4 = 100+10+20+30 = 160
	fmt.Printf("Test 6 - Multi-range sum = %d (expected 160)\n", sum6)

	// Test 7: Override formula cell with Set.
	ex7 := Constructor(3, 'C')
	ex7.Set(1, 0, 5)
	ex7.Sum(2, 0, []string{"A1"})
	ex7.Set(2, 0, 42) // Override A2 formula with direct value
	fmt.Printf("Test 7 - A2 after override = %d (expected 42)\n", ex7.Get(2, 0))
}
```
