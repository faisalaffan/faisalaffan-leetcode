# 3612 — Process String With Special Operations I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func ProcessStringWithSpecialOperationsI(s string, ops []int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3612: Process String with Special Operations I
// https://leetcode.com/problems/process-string-with-special-operations-i/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", ProcessStringWithSpecialOperationsI("abc", []int{1, 0, 1}))
	// Test case 2
	fmt.Println("Test 2:", ProcessStringWithSpecialOperationsI("ab", []int{1, 1}))
	// Test case 3
	fmt.Println("Test 3:", ProcessStringWithSpecialOperationsI("x", []int{0}))
}

func ProcessStringWithSpecialOperationsI(s string, ops []int) string {
	b := []byte(s)
	for i, op := range ops {
		if i >= len(b) {
			break
		}
		if op == 1 {
			// toggle case
			if b[i] >= 'a' && b[i] <= 'z' {
				b[i] = b[i] - 'a' + 'A'
			} else if b[i] >= 'A' && b[i] <= 'Z' {
				b[i] = b[i] - 'A' + 'a'
			}
		}
		// 0 = no operation
	}
	return string(b)
}
```
