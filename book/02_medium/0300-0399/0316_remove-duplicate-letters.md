# 0316 — Remove Duplicate Letters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func removeDuplicateLetters(s string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #316: Remove Duplicate Letters
// https://leetcode.com/problems/remove-duplicate-letters/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func removeDuplicateLetters(s string) string {
	// Count last occurrence of each character
	lastOccur := [26]int{}
  // Range loop
	for i := range s {
		lastOccur[s[i]-'a'] = i
	}

	stack := make([]byte, 0, len(s))
	seen := [26]bool{}

  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if seen[ch-'a'] {
			continue
		}

		// Pop while stack top is greater and appears later
		for len(stack) > 0 && ch < stack[len(stack)-1] && i < lastOccur[stack[len(stack)-1]-'a'] {
			seen[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, ch)
		seen[ch-'a'] = true
	}

	return string(stack)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", removeDuplicateLetters("bcabc"))
	// Expected: "abc"

	// Test case 2
	fmt.Println("Test 2:", removeDuplicateLetters("cbacdcbc"))
	// Expected: "acdb"

	// Test case 3: Single character
	fmt.Println("Test 3:", removeDuplicateLetters("aaaa"))
	// Expected: "a"
}
```
