# 0420 — Strong Password Checker

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func strongPasswordChecker(password string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #420: Strong Password Checker
// https://leetcode.com/problems/strong-password-checker/
// Difficulty: Hard
//
// A password is strong if:
//   1. Length 6..20
//   2. Contains at least one lowercase, one uppercase, one digit
//   3. No three consecutive repeating characters
// Returns the minimum number of changes (insert, delete, replace).
// Handles three cases: too short, too long, or in-range with missing types.

import (
	"fmt"
)

func main() {
	// Example 1: "a" -> 5 (insert 5 chars: need length+missing types)
	fmt.Println(strongPasswordChecker("a"))
	// Example 2: "aA1" -> 3 (need 3 more chars + missing types)
	fmt.Println(strongPasswordChecker("aA1"))
	// Example 3: "1337C0d3" -> 0 (already strong)
	fmt.Println(strongPasswordChecker("1337C0d3"))
	// Example 4: "aaa123" -> 1 (replace one 'a')
	fmt.Println(strongPasswordChecker("aaa123"))
	// Example 5: "aaa" -> 3 (insert 3: need length + fix repeat + missing types)
	fmt.Println(strongPasswordChecker("aaa"))
	// Example 6: "abababababababababaaa" (20+ chars with repeating) -> 3
	fmt.Println(strongPasswordChecker("abababababababababaaa"))
}

func strongPasswordChecker(password string) int {
	n := len(password)

	// Count missing character types
	hasLower, hasUpper, hasDigit := false, false, false
	for i := 0; i < n; i++ {
		ch := password[i]
		if ch >= 'a' && ch <= 'z' {
			hasLower = true
		} else if ch >= 'A' && ch <= 'Z' {
			hasUpper = true
		} else if ch >= '0' && ch <= '9' {
			hasDigit = true
		}
	}

	missing := 0
	if !hasLower {
		missing++
	}
	if !hasUpper {
		missing++
	}
	if !hasDigit {
		missing++
	}

	if n < 6 {
		// Too short. We need to add (6-n) chars.
		// Each add can fix one missing type and one repeat.
		return max(6-n, missing)
	}

	if n <= 20 {
		// In range. Only need to fix repeats and missing types.
		// One replace fixes one repeat block (every 3rd char) or one missing type.
		replaces := 0
		i := 0
		for i < n {
			j := i
			for j < n && password[j] == password[i] {
				j++
			}
			repeats := j - i
			if repeats >= 3 {
				replaces += repeats / 3
			}
			i = j
		}
		return max(replaces, missing)
	}

	// Too long (n > 20). Need to delete (n-20) chars + fix repeats + fix missing.
	// Strategy: delete chars strategically to break up repeat sequences.
	// For each repeat block of length L:
	//   - Deleting 1 reduces replacements needed by 1 if L%3 == 0
	//   - Deleting 2 reduces replacements needed by 1 if L%3 == 1 (after first delete)
	//   - Deleting 3 reduces replacements needed by 1 otherwise
	// We prioritize deletes that give the biggest reduction in replaces.

	over := n - 20
	replaces := 0

	// Count repeats and group by L%3
	type repeatInfo struct{ length int }
	repeats := make([]repeatInfo, 0)

	i := 0
	for i < n {
		j := i
		for j < n && password[j] == password[i] {
			j++
		}
		l := j - i
		if l >= 3 {
			repeats = append(repeats, repeatInfo{l})
			replaces += l / 3
		}
		i = j
	}

	if over <= 0 {
		return max(replaces, missing)
	}

	// Use deletes to reduce replaces
	// Priority: blocks where L%3==0 -> delete 1 reduces replaces by 1
	// Then: blocks where L%3==1 -> delete 2 reduces replaces by 1
	// Then: any block -> delete 3 reduces replaces by 1
  // Range loop
	for i := range repeats {
		if over <= 0 {
			break
		}
		if repeats[i].length%3 == 0 {
			need := 1
			if over >= need {
				replaces -= repeats[i].length / 3
				repeats[i].length -= need
				replaces += repeats[i].length / 3
				over -= need
			}
		}
	}
  // Range loop
	for i := range repeats {
		if over <= 0 {
			break
		}
		if repeats[i].length%3 == 1 {
			need := 2
			if over >= need {
				replaces -= repeats[i].length / 3
				repeats[i].length -= need
				replaces += repeats[i].length / 3
				over -= need
			}
		}
	}
  // Range loop
	for i := range repeats {
		if over <= 0 {
			break
		}
		// Each 3 removes reduces replaces by 1
		need := repeats[i].length - 2 // we save replaces once length drops below 3
		if need <= 0 {
			// Already below 3 after previous deletes; no replace needed
			continue
		}
		if over >= need {
			replaces -= repeats[i].length / 3
			repeats[i].length -= need
			replaces += repeats[i].length / 3
			over -= need
		} else {
			// Partial: each group of 3 deletes eliminates 1 replace
			// Actually more nuanced: for any block with length >= 3, each 3 deletes
			// that keep length still >= 3 reduces replaces by 1
			replaces -= over / 3
			break
		}
	}

	return over + max(replaces, missing)
}
```
