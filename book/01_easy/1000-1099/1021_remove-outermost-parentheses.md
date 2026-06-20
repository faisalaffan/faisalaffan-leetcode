# 1021 — Remove Outermost Parentheses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func removeOuterParentheses(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1021: Remove Outermost Parentheses
// https://leetcode.com/problems/remove-outermost-parentheses/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))          // "()()()"
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))  // "()()()()(())"
	fmt.Println(removeOuterParentheses("()()"))                // ""
}

// LeetCode submission: removeOuterParentheses
func removeOuterParentheses(s string) string {
	ans := make([]byte, 0, len(s))
	depth := 0
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if depth > 0 {
				ans = append(ans, '(')
			}
			depth++
		} else {
			depth--
			if depth > 0 {
				ans = append(ans, ')')
			}
		}
	}
	return string(ans)
}
```
