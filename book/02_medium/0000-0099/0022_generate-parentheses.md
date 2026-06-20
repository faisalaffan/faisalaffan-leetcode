# 0022 — Generate Parentheses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func generateParenthesis(n int) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Backtracking

**Waktu:** O(4^n / sqrt(n))  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Backtracking** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #22: Generate Parentheses
// https://leetcode.com/problems/generate-parentheses/
// Difficulty: Medium

import "fmt"

func generateParenthesis(n int) []string {
	result := []string{}
	var backtrack func(curr string, open, close int)
	backtrack = func(curr string, open, close int) {
		if len(curr) == 2*n {
			result = append(result, curr)
			return
		}
		if open < n {
			backtrack(curr+"(", open+1, close)
		}
		if close < open {
			backtrack(curr+")", open, close+1)
		}
	}
	backtrack("", 0, 0)
	return result
}

func main() {
	// Test case 1
	fmt.Println(generateParenthesis(3)) // ["((()))","(()())","(())()","()(())","()()()"]

	// Test case 2
	fmt.Println(generateParenthesis(1)) // ["()"]

	// Test case 3
	fmt.Println(generateParenthesis(2)) // ["(())","()()"]
}

// Time: O(4^n / sqrt(n)) | Space: O(n)
```
