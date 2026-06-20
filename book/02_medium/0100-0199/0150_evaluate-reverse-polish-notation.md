# 0150 — Evaluate Reverse Polish Notation

## Deskripsi

**Soal:** [0150. Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO), LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func evalRPN(tokens []string) int`

## Solusi Go

```go
package main

// LeetCode #150: Evaluate Reverse Polish Notation
// https://leetcode.com/problems/evaluate-reverse-polish-notation/
// Difficulty: Medium

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
  // Membuat slice untuk menyimpan hasil
	stack := make([]int, 0, len(tokens))

	for _, token := range tokens {
		switch token {
		case "+":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		case "-":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a-b)
		case "*":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		case "/":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a/b)
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func main() {
	// Test case 1
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"})) // 9

	// Test case 2
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"})) // 6

	// Test case 3
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})) // 22
}

// Time: O(n) | Space: O(n)
```
