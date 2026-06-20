# 2011 — Final Value Of Variable After Performing Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FinalValueOfVariableAfterPerformingOperations(operations []string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2011: Final Value of Variable After Performing Operations
// https://leetcode.com/problems/final-value-of-variable-after-performing-operations/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FinalValueOfVariableAfterPerformingOperations([]string{"--X", "X++", "X++"}))     // 1
	fmt.Println(FinalValueOfVariableAfterPerformingOperations([]string{"++X", "++X", "X++"}))     // 3
	fmt.Println(FinalValueOfVariableAfterPerformingOperations([]string{"X++", "++X", "--X", "X--"})) // 0
}

// Time: O(n), Space: O(1)
func FinalValueOfVariableAfterPerformingOperations(operations []string) int {
	x := 0
	for _, op := range operations {
		if op[1] == '+' {
			x++
		} else {
			x--
		}
	}
	return x
}
```
