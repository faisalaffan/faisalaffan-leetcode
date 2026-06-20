# 0856 — Score Of Parentheses

## Deskripsi

**Soal:** [0856. Score Of Parentheses](https://leetcode.com/problems/score-of-parentheses/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

## Solusi Go

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
