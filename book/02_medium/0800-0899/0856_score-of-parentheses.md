# 0856 — Score Of Parentheses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func ScoreOfParentheses(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #856: Score of Parentheses
// https://leetcode.com/problems/score-of-parentheses/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ScoreOfParentheses("()"))
	fmt.Println(ScoreOfParentheses("(())"))
	fmt.Println(ScoreOfParentheses("()()"))
	fmt.Println(ScoreOfParentheses("(()(()))"))
}

// Time: O(n) | Space: O(n)
func ScoreOfParentheses(s string) int {
	stack := []int{0}
	for _, ch := range s {
		if ch == '(' {
			stack = append(stack, 0)
		} else {
			x := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if x != 0 {
				x *= 2
			} else {
				x = 1
			}
			stack[len(stack)-1] += x
		}
	}
	return stack[0]
}
```
