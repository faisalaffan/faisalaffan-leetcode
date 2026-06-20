# 1392 — Longest Happy Prefix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func longestPrefix(s string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1392: Longest Happy Prefix
// https://leetcode.com/problems/longest-happy-prefix/
// Difficulty: Hard

import "fmt"

// longestPrefix finds the longest happy prefix using KMP prefix function.
// A happy prefix is a non-empty proper prefix that is also a suffix.
func longestPrefix(s string) string {
	n := len(s)
	if n <= 1 {
		return ""
	}

	// Build KMP LPS (Longest Proper Prefix which is also Suffix) array
  // Alokasi slice
	lps := make([]int, n)
	for i := 1; i < n; i++ {
		j := lps[i-1]
		for j > 0 && s[i] != s[j] {
			j = lps[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		lps[i] = j
	}

	// lps[n-1] is the length of the longest proper prefix that is also a suffix
	length := lps[n-1]
	return s[:length]
}

func main() {
	// Test case 1
	fmt.Println("Test 1: s = \"level\"")
	result1 := longestPrefix("level")
	expected1 := "l"
	fmt.Printf("Result: %q (expected: %q)\n", result1, expected1)

	// Test case 2
	fmt.Println("\nTest 2: s = \"ababab\"")
	result2 := longestPrefix("ababab")
	expected2 := "abab"
	fmt.Printf("Result: %q (expected: %q)\n", result2, expected2)

	// Test case 3
	fmt.Println("\nTest 3: s = \"leetcodeleet\"")
	result3 := longestPrefix("leetcodeleet")
	expected3 := "leet"
	fmt.Printf("Result: %q (expected: %q)\n", result3, expected3)

	// Test case 4
	fmt.Println("\nTest 4: s = \"a\"")
	result4 := longestPrefix("a")
	expected4 := ""
	fmt.Printf("Result: %q (expected: %q)\n", result4, expected4)

	// Test case 5
	fmt.Println("\nTest 5: s = \"aaaaa\"")
	result5 := longestPrefix("aaaaa")
	expected5 := "aaaa"
	fmt.Printf("Result: %q (expected: %q)\n", result5, expected5)
}
```
