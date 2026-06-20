# 0316 — Remove Duplicate Letters

## Deskripsi

**Soal:** [0316. Remove Duplicate Letters](https://leetcode.com/problems/remove-duplicate-letters/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func removeDuplicateLetters(s string) string`

## Solusi Go

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
  // Iterasi seluruh elemen
	for i := range s {
		lastOccur[s[i]-'a'] = i
	}

  // Membuat slice untuk menyimpan hasil
	stack := make([]byte, 0, len(s))
	seen := [26]bool{}

  // Loop standar: indeks 0 sampai n-1
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
