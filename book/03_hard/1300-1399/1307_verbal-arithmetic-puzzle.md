# 1307 — Verbal Arithmetic Puzzle

## Deskripsi

**Soal:** [1307. Verbal Arithmetic Puzzle](https://leetcode.com/problems/verbal-arithmetic-puzzle/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Backtracking (pelacakan mundur)

**Fungsi Solusi:** `func isSolvable(words []string, result string) bool`

> **Ide Kunci:** Backtracking with digit assignment.

## Solusi Go

```go
package main

// LeetCode #1307: Verbal Arithmetic Puzzle
// https://leetcode.com/problems/verbal-arithmetic-puzzle/
// Difficulty: Hard
//
// Approach: Backtracking with digit assignment.
// Collect all unique letters (max 10). Try each digit 0-9 for each letter,
// respecting the constraint that leading letters cannot be 0.
// When all letters are assigned, verify the equation: word[0] + ... + word[k] == result.
// Prune: leading-letter-zero check, no digit reuse.

import "fmt"

func isSolvable(words []string, result string) bool {
  // Membuat map untuk pencarian O(1): key → value
	letterSet := make(map[byte]bool)
	addLetters := func(s string) {
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(s); i++ {
			letterSet[s[i]] = true
		}
	}
	for _, w := range words {
		addLetters(w)
	}
	addLetters(result)

  // Membuat slice untuk menyimpan hasil
	letters := make([]byte, 0, len(letterSet))
	for c := range letterSet {
		letters = append(letters, c)
	}
	if len(letters) > 10 {
		return false
	}

  // Membuat map untuk pencarian O(1): key → value
	nonZero := make(map[byte]bool)
	for _, w := range words {
		if len(w) > 1 {
			nonZero[w[0]] = true
		}
	}
	if len(result) > 1 {
		nonZero[result[0]] = true
	}

  // Membuat map untuk pencarian O(1): key → value
	mapping := make(map[byte]int)
  // Membuat slice untuk menyimpan hasil
	used := make([]bool, 10)

	var dfs func(int) bool
	dfs = func(idx int) bool {
		if idx == len(letters) {
			sum := 0
			for _, w := range words {
				val := 0
  // Loop standar: indeks 0 sampai n-1
				for i := 0; i < len(w); i++ {
					val = val*10 + mapping[w[i]]
				}
				sum += val
			}
			res := 0
  // Loop standar: indeks 0 sampai n-1
			for i := 0; i < len(result); i++ {
				res = res*10 + mapping[result[i]]
			}
			return sum == res
		}

		c := letters[idx]
		for d := 0; d <= 9; d++ {
			if used[d] {
				continue
			}
			if d == 0 && nonZero[c] {
				continue
			}
			used[d] = true
			mapping[c] = d
			if dfs(idx + 1) {
				return true
			}
			used[d] = false
		}
		return false
	}

	return dfs(0)
}

func main() {
	fmt.Println(isSolvable([]string{"SEND", "MORE"}, "MONEY"))                  // true
	fmt.Println(isSolvable([]string{"SIX", "SEVEN", "SEVEN"}, "TWENTY"))        // true
	fmt.Println(isSolvable([]string{"LEET", "CODE"}, "POINT"))                  // false
	fmt.Println(isSolvable([]string{"A", "B"}, "A"))                            // true (B=0)
}
```
