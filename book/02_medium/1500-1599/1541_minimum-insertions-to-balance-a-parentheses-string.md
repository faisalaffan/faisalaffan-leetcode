# 1541 — Minimum Insertions To Balance A Parentheses String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func MinInsertions(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1541: Minimum Insertions to Balance a Parentheses String
// https://leetcode.com/problems/minimum-insertions-to-balance-a-parentheses-string/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinInsertions("(()))"))
	fmt.Println(MinInsertions("())"))
	fmt.Println(MinInsertions("))())("))
}

func MinInsertions(s string) int {
	// Time: O(N), Space: O(1)
	// Each '(' needs two ')' to balance.
	insertions := 0
	open := 0 // number of '(' that need closing

  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			open++
		} else { // ')'
			if open > 0 {
				// Check if next char is also ')'
				if i+1 < len(s) && s[i+1] == ')' {
					// Both ')' found, consume both
					i++ // skip next ')'
				} else {
					// Need one more ')'
					insertions++
				}
				open--
			} else {
				// Need a '(' before this ')'
				if i+1 < len(s) && s[i+1] == ')' {
					// Insert '(' and consume both ')'
					insertions++ // for '('
					i++          // consume next ')'
				} else {
					// Insert '(' and one ')'
					insertions += 2
				}
			}
		}
	}

	// Each remaining '(' needs two ')'
	insertions += open * 2

	return insertions
}
```
