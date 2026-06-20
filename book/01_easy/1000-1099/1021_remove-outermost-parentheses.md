# 1021 — Remove Outermost Parentheses

## Deskripsi

**Soal:** [1021. Remove Outermost Parentheses](https://leetcode.com/problems/remove-outermost-parentheses/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
	ans := make([]byte, 0, len(s))
	depth := 0
  // Loop standar: indeks 0 sampai n-1
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
