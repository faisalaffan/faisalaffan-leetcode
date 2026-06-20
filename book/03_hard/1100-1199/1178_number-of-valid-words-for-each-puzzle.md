# 1178 — Number Of Valid Words For Each Puzzle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func wordMask(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1178: Number of Valid Words for Each Puzzle
// https://leetcode.com/problems/number-of-valid-words-for-each-puzzle/
// Difficulty: Hard

import "fmt"

// wordMask converts a word to a bitmask of its letters (26 bits).
func wordMask(s string) int {
	mask := 0
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		mask |= 1 << (s[i] - 'a')
	}
	return mask
}

func findNumOfValidWords(words []string, puzzles []string) []int {
	// Count words by mask
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	for _, w := range words {
		m := wordMask(w)
		// Optimization: skip words with > 2^7 bits (not subset of any puzzle with 7 letters)
		// Actually just count normally
		freq[m]++
	}

  // Alokasi slice
	result := make([]int, len(puzzles))
	for i, p := range puzzles {
		pmask := wordMask(p)
		firstBit := 1 << (p[0] - 'a')
		count := 0

		// Enumerate all subsets of pmask that include the first letter
		sub := pmask
		for sub > 0 {
			if (sub & firstBit) != 0 {
				count += freq[sub]
			}
			sub = (sub - 1) & pmask
		}
		// Also check empty subset? No, word must have at least one letter (first letter)
		// Actually the subset loop above only goes through non-zero subsets of pmask
		// that include firstBit. But we already covered that via the loop.
		result[i] = count
	}
	return result
}

func main() {
	// Test case
	words := []string{"aaaa", "asas", "able", "ability", "actt", "access"}
	puzzles := []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}
	fmt.Println(findNumOfValidWords(words, puzzles))
	// Expected: [1, 1, 3, 2, 4, 0]

	// Test case 2: empty
	fmt.Println(findNumOfValidWords([]string{}, []string{"a"})) // [0]
}
```
