# 2019 — The Score Of Students Solving Math Expression

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func scoreOfStudents(s string, answers []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP, Stack

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2019: The Score of Students Solving Math Expression
// https://leetcode.com/problems/the-score-of-students-solving-math-expression/
// Difficulty: Hard
//
// Compute all possible results from different parenthesizations of an expression
// containing digits, '+', and '*'. Score student answers: +5 for correct answer
// (using standard *-before-+ precedence), +2 for any other achievable result, 0 otherwise.

import "fmt"

func main() {
	// Example 1
	fmt.Println(scoreOfStudents("7+3*1*2", []int{20, 13, 42}))

	// Example 2
	fmt.Println(scoreOfStudents("3+5*2", []int{13, 0, 10, 13, 13, 16, 16}))

	// Example 3
	fmt.Println(scoreOfStudents("6+0*1", []int{12, 9, 6, 4, 8, 6}))

	// Single number
	fmt.Println(scoreOfStudents("5", []int{5, 0, 10}))
}

func scoreOfStudents(s string, answers []int) int {
	// Parse expression into numbers and operators
	nums := []int{}
	ops := []byte{}
	num := 0
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		} else {
			nums = append(nums, num)
			ops = append(ops, s[i])
			num = 0
		}
	}
	nums = append(nums, num)

	n := len(nums)

	// dp[l][r] = set of possible results for subexpression nums[l..r]
  // Matriks 2D
	dp := make([][]map[int]bool, n)
  // Range loop
	for i := range dp {
		dp[i] = make([]map[int]bool, n)
		for j := range dp[i] {
			dp[i][j] = make(map[int]bool)
		}
		dp[i][i][nums[i]] = true
	}

	// Fill DP for increasing subexpression lengths
	for length := 2; length <= n; length++ {
		for l := 0; l+length <= n; l++ {
			r := l + length - 1
			for k := l; k < r; k++ {
				for a := range dp[l][k] {
					for b := range dp[k+1][r] {
						var val int
						if ops[k] == '+' {
							val = a + b
						} else {
							val = a * b
						}
						if val <= 1000 {
							dp[l][r][val] = true
						}
					}
				}
			}
		}
	}

	possible := dp[0][n-1]
	correct := evaluateStandard(nums, ops)

	total := 0
	for _, ans := range answers {
		if ans == correct {
			total += 5
		} else if possible[ans] {
			total += 2
		}
	}
	return total
}

// evaluateStandard computes the result using standard *-before-+ precedence.
func evaluateStandard(nums []int, ops []byte) int {
	stack := []int{nums[0]}
	for i, op := range ops {
		if op == '+' {
			stack = append(stack, nums[i+1])
		} else {
			stack[len(stack)-1] *= nums[i+1]
		}
	}
	sum := 0
	for _, v := range stack {
		sum += v
	}
	return sum
}
```
