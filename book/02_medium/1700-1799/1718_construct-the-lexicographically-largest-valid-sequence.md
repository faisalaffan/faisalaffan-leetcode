# 1718 — Construct The Lexicographically Largest Valid Sequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func constructDistancedSequence(n int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Backtracking

**Waktu:** O(n!), Space: O(n) for backtracking  |  **Ruang:** O(n) for backtracking

> 🎓 **Fresh Grad Tips:** Kuasai **Backtracking** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1718: Construct the Lexicographically Largest Valid Sequence
// https://leetcode.com/problems/construct-the-lexicographically-largest-valid-sequence/
// Difficulty: Medium
// Time: O(n!), Space: O(n) for backtracking

import "fmt"

func constructDistancedSequence(n int) []int {
	length := 2*n - 1
  // Alokasi slice
	result := make([]int, length)
	used := make([]bool, n+1)

	var backtrack func(pos int) bool
	backtrack = func(pos int) bool {
		if pos == length {
			return true
		}
		if result[pos] != 0 {
			return backtrack(pos + 1)
		}

		// Try largest number first for lexicographically largest
		for num := n; num >= 1; num-- {
			if used[num] {
				continue
			}
			if num == 1 {
				result[pos] = 1
				used[1] = true
				if backtrack(pos + 1) {
					return true
				}
				result[pos] = 0
				used[1] = false
			} else {
				nextPos := pos + num
				if nextPos < length && result[nextPos] == 0 {
					result[pos] = num
					result[nextPos] = num
					used[num] = true
					if backtrack(pos + 1) {
						return true
					}
					result[pos] = 0
					result[nextPos] = 0
					used[num] = false
				}
			}
		}
		return false
	}

	backtrack(0)
	return result
}

func main() {
	fmt.Println(constructDistancedSequence(3)) // Expected: [3, 1, 2, 3, 2]
	fmt.Println(constructDistancedSequence(5)) // Expected: [5, 3, 1, 4, 3, 5, 2, 4, 2]
	fmt.Println(constructDistancedSequence(1)) // Expected: [1]
}
```
