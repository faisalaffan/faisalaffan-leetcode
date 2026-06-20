# 1307 — Verbal Arithmetic Puzzle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func isSolvable(words []string, result string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Backtracking

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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
  // HashMap: O(1) lookup
	letterSet := make(map[byte]bool)
	addLetters := func(s string) {
  // Linear scan O(n)
		for i := 0; i < len(s); i++ {
			letterSet[s[i]] = true
		}
	}
	for _, w := range words {
		addLetters(w)
	}
	addLetters(result)

	letters := make([]byte, 0, len(letterSet))
	for c := range letterSet {
		letters = append(letters, c)
	}
	if len(letters) > 10 {
		return false
	}

  // HashMap: O(1) lookup
	nonZero := make(map[byte]bool)
	for _, w := range words {
		if len(w) > 1 {
			nonZero[w[0]] = true
		}
	}
	if len(result) > 1 {
		nonZero[result[0]] = true
	}

  // HashMap: O(1) lookup
	mapping := make(map[byte]int)
	used := make([]bool, 10)

	var dfs func(int) bool
	dfs = func(idx int) bool {
		if idx == len(letters) {
			sum := 0
			for _, w := range words {
				val := 0
  // Linear scan O(n)
				for i := 0; i < len(w); i++ {
					val = val*10 + mapping[w[i]]
				}
				sum += val
			}
			res := 0
  // Linear scan O(n)
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
