# 1433 — Check If A String Can Break Another String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func checkIfCanBreak(s1 string, s2 string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1433: Check If a String Can Break Another String
// https://leetcode.com/problems/check-if-a-string-can-break-another-string/
// Difficulty: Medium (listed here as Hard)
//
// A string a can break string b if after sorting both strings,
// for every i, a[i] >= b[i] (or b[i] >= a[i] for all i).

import (
	"fmt"
	"sort"
)

// checkIfCanBreak returns true if s1 can break s2 or s2 can break s1.
func checkIfCanBreak(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// Sort both strings
	b1 := []byte(s1)
	b2 := []byte(s2)
  // Custom sort
	sort.Slice(b1, func(i, j int) bool { return b1[i] < b1[j] })
  // Custom sort
	sort.Slice(b2, func(i, j int) bool { return b2[i] < b2[j] })

	// Check if s1 can break s2
	s1BreaksS2 := true
  // Linear scan O(n)
	for i := 0; i < len(b1); i++ {
		if b1[i] < b2[i] {
			s1BreaksS2 = false
			break
		}
	}
	if s1BreaksS2 {
		return true
	}

	// Check if s2 can break s1
	s2BreaksS1 := true
  // Linear scan O(n)
	for i := 0; i < len(b1); i++ {
		if b2[i] < b1[i] {
			s2BreaksS1 = false
			break
		}
	}
	return s2BreaksS1
}

func main() {
	// Test case 1
	s1, s2 := "abc", "xya"
	result1 := checkIfCanBreak(s1, s2)
	fmt.Printf("Test 1: s1=%q, s2=%q => %v (expected true)\n", s1, s2, result1)

	// Test case 2
	s1, s2 = "abe", "acd"
	result2 := checkIfCanBreak(s1, s2)
	fmt.Printf("Test 2: s1=%q, s2=%q => %v (expected false)\n", s1, s2, result2)

	// Test case 3
	s1, s2 = "leetcodee", "interview"
	result3 := checkIfCanBreak(s1, s2)
	fmt.Printf("Test 3: s1=%q, s2=%q => %v (expected true)\n", s1, s2, result3)

	// Test case 4: equal strings
	s1, s2 = "abc", "abc"
	result4 := checkIfCanBreak(s1, s2)
	fmt.Printf("Test 4: s1=%q, s2=%q => %v (expected true)\n", s1, s2, result4)

	// Test case 5: single character
	s1, s2 = "a", "b"
	result5 := checkIfCanBreak(s1, s2)
	fmt.Printf("Test 5: s1=%q, s2=%q => %v (expected true, b breaks a)\n", s1, s2, result5)
}
```
